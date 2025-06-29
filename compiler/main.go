package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"main/arm"
	"main/errors"
	compiler "main/parser"
	repl "main/repl"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/antlr4-go/antlr/v4"
)

var (
	currentReportHTML       string
	reportFilePath          string
	currentErrorsReportHTML string
	symbolsReportPath       string
	errorsReportPath        string
)

func main() {
	// -------------------- INTERFAZ --------------------
	//Inicio de la ventana
	a := app.NewWithID("ocl2.vlang.cherry")

	w := a.NewWindow("V-Lang Cherry")
	os.Mkdir("reportes", 0755)
	symbolsReportPath = "./reportes/reporte_simbolos.html"
	errorsReportPath = "./reportes/reporte_errores.html"

	ta1 := widget.NewMultiLineEntry()
	ta1.SetPlaceHolder("Escribe tu código aquí...")
	ta1.Resize(fyne.NewSize(770, 700))
	ta1.Move(fyne.NewPos(10, 10))

	ta2 := widget.NewMultiLineEntry()
	ta2.Wrapping = fyne.TextWrapWord
	ta2.Resize(fyne.NewSize(770, 700))
	ta2.Move(fyne.NewPos(810, 10))

	positionLabel := widget.NewLabel("Línea: 1, Columna: 1")

	updateCursorPosition := func() {
		pos := ta1.CursorRow // empieza en 0
		col := ta1.CursorColumn

		// Mostrar posición empezando desde 1
		positionLabel.SetText("Línea: " + strconv.Itoa(pos+1) + ", Columna: " + strconv.Itoa(col+1))
	}

	ta1.OnChanged = func(s string) {
		updateCursorPosition()
	}

	// Detectar movimiento del cursor
	ta1.OnCursorChanged = func() {
		updateCursorPosition()
	}

	b1 := widget.NewButton("Ejecutar", func() {
		inputCode := ta1.Text

		//startTime := time.Now()

		// Canal para generar CST en paralelo (opcional)

		// 1. Análisis Léxico
		lexicalErrorListener := errors.NewLexicalErrorListener()
		lexer := compiler.NewVlangLexer(antlr.NewInputStream(inputCode))
		lexer.RemoveErrorListeners()
		lexer.AddErrorListener(lexicalErrorListener)

		// 2. Tokens
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		// 3. Parser
		parser := compiler.NewVlangParser(stream)
		parser.BuildParseTrees = true

		// Usamos la misma tabla de errores para ambos listeners
		errorTable := repl.NewErrorTable()
		syntaxErrorListener := errors.NewSyntaxErrorListener(errorTable)
		parser.RemoveErrorListeners()
		parser.SetErrorHandler(errors.NewCustomErrorStrategy())
		parser.AddErrorListener(syntaxErrorListener)

		// 4. Árbol sintáctico
		arbolito := parser.Programa()

		// 5. Visitor - usa la misma tabla de errores
		//visitor := repl.NewReplVisitor(errorTable)
		//visitor.Visit(arbolito)
		//nuevaVariable := visitor.Console.GetOutput()
		//ta2.SetText(nuevaVariable)

		// Visitor ARM
		visitorArm := arm.NewVisitorArm()
		visitorArm.Visit(arbolito)
		ta2.SetText(visitorArm.Generator.String())

		err2 := os.WriteFile("./assembler/program.s", []byte(visitorArm.Generator.String()), 0644)
		if err2 != nil {
			fmt.Println("Error al guardar program.s:", err2)
		} else {
			fmt.Println("Archivo program.s generado exitosamente.")
		}

		ta2.SetText(visitorArm.Generator.String())

		// Generación de reporte de errores
		currentErrorsReportHTML = errorTable.GenerateHTMLReport()
		err := os.WriteFile(errorsReportPath, []byte(currentErrorsReportHTML), 0644)
		if err != nil {
			fmt.Println("Error al guardar reporte de errores:", err)
		}
		/*
			//Generación de reporte de símbolos
			if visitor.ScopeTrace != nil {
				currentReportHTML = visitor.ScopeTrace.GenerateHTMLReport()
				err := os.WriteFile(reportFilePath, []byte(currentReportHTML), 0644)
				if err != nil {
					fmt.Println("Error al guardar el reporte:", err)
				} else {
					fmt.Println("Reporte de símbolos generado en:", reportFilePath)
				}
			}
		*/

		//generación del CST
		dotContent := generateDotTree(arbolito, parser.RuleNames)
		cstPath := "./reportes/cst.svg"
		err = generateTreeSVG(dotContent, cstPath)
		if err != nil {
			fmt.Println("Error al generar CST:", err)
		}
	})
	b1.Resize(fyne.NewSize(150, 45))
	b1.Move(fyne.NewPos(720, 730))

	menu := fyne.NewMainMenu(
		fyne.NewMenu("Archivo",
			fyne.NewMenuItem("Crear Archivo", func() {
				ta1.SetText("")
				ta2.SetText("")
			}),
			fyne.NewMenuItem("Abrir archivo", func() {
				fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
					if err != nil {
						dialog.ShowError(err, w)
						return
					}
					if reader == nil {
						return // el usuario canceló
					}

					defer reader.Close()
					data, err := ioutil.ReadAll(reader)
					if err != nil {
						dialog.ShowError(err, w)
						return
					}
					ta1.SetText(string(data))
				}, w)

				fd.SetFilter(storage.NewExtensionFileFilter([]string{".vch"}))
				fd.Show()
			}),
			fyne.NewMenuItem("Guardar archivo", func() {
				fd := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
					if err != nil {
						dialog.ShowError(err, w)
						return
					}
					if writer == nil {
						return // usuario canceló
					}

					// Escribir el contenido
					_, err = io.WriteString(writer, ta1.Text)
					if err != nil {
						dialog.ShowError(err, w)
						return
					}
					writer.Close()
				}, w)

				fd.SetFileName("archivo.vch") // nombre por defecto
				fd.SetFilter(storage.NewExtensionFileFilter([]string{".vch"}))
				fd.Show()
			}),
		),
		fyne.NewMenu("Reportes",
			fyne.NewMenuItem("Errores", func() {
				if currentErrorsReportHTML == "" {
					dialog.ShowInformation("Reporte", "No hay reporte de errores disponible. Ejecuta el código primero.", w)
					return
				}

				var cmd *exec.Cmd
				switch runtime.GOOS {
				case "windows":
					cmd = exec.Command("cmd", "/c", "start", errorsReportPath)
				case "darwin":
					cmd = exec.Command("open", errorsReportPath)
				default: // linux y otros
					cmd = exec.Command("xdg-open", errorsReportPath)
				}

				err := cmd.Start()
				if err != nil {
					fmt.Println("Error al abrir el reporte de errores:", err)
				}
			}),
			fyne.NewMenuItem("Tabla de Símbolos", func() {
				// funcion para el reporte de tabla de simbolos
				if currentReportHTML == "" {
					dialog.ShowInformation("Reporte", "No hay reporte disponible", w)
					return
				}

				// Guardar el reporte en un archivo HTML
				reportFilePath = "./reportes/reporte_simbolos.html"
				err := os.WriteFile(reportFilePath, []byte(currentReportHTML), 0644)
				if err != nil {
					fmt.Println("Error al guardar el reporte:", err)
					return
				}

				dialog.ShowInformation("Reporte", "El reporte se ha guardado en "+reportFilePath, w)

				var cmd *exec.Cmd
				switch runtime.GOOS {
				case "windows":
					cmd = exec.Command("cmd", "/c", "start", reportFilePath)
				case "darwin":
					cmd = exec.Command("open", reportFilePath)
				default: // linux y otros
					cmd = exec.Command("xdg-open", reportFilePath)
				}

				err = cmd.Start()
				if err != nil {
					dialog.ShowError(fmt.Errorf("no se pudo abrir el navegador: %v", err), w)
					return
				}
			}),
			fyne.NewMenuItem("CST", func() {
				cstPath := "./reportes/cst.svg"

				// Verificar si el archivo existe
				if _, err := os.Stat(cstPath); os.IsNotExist(err) {
					dialog.ShowInformation("Reporte", "No hay CST disponible. Ejecuta el código primero.", w)
					return
				}

				// Obtener ruta absoluta
				absPath, err := filepath.Abs(cstPath)
				if err != nil {
					dialog.ShowError(fmt.Errorf("no se pudo obtener ruta absoluta: %v", err), w)
					return
				}

				// Convertir a URL file://
				fileUrl := "file://" + absPath

				// Abrir en el navegador predeterminado
				cmd := exec.Command("xdg-open", fileUrl)
				if err := cmd.Start(); err != nil {
					// Si falla xdg-open, probar con navegadores específicos
					browsers := []string{"google-chrome", "chrome", "chromium", "firefox", "microsoft-edge"}
					for _, browser := range browsers {
						cmd = exec.Command(browser, fileUrl)
						if err := cmd.Start(); err == nil {
							return
						}
					}
					dialog.ShowError(fmt.Errorf("no se pudo abrir el navegador: %v", err), w)
				}
			}),
		),
	)

	w.SetMainMenu(menu)
	content := container.NewBorder(nil, positionLabel, nil, nil, container.NewWithoutLayout(ta1, b1, ta2))

	w.SetContent(content)
	w.Resize(fyne.NewSize(1920, 1080))
	w.ShowAndRun()

}

// Funciones para visualizar nuestro arbol
func PrintVerticalTree(node antlr.Tree, ruleNames []string) {
	printVerticalNode(node, ruleNames, "", true)
}

func printVerticalNode(node antlr.Tree, ruleNames []string, prefix string, isLast bool) {
	connector := "+-- "
	if !isLast {
		connector = "|-- "
	}

	var label string
	switch n := node.(type) {
	case antlr.RuleNode:
		label = ruleNames[n.GetRuleContext().GetRuleIndex()]
	case antlr.TerminalNode:
		label = fmt.Sprintf("\"%s\"", n.GetText())
	default:
		label = fmt.Sprintf("%T", n)
	}

	fmt.Printf("%s%s%s\n", prefix, connector, label)

	// Actualizar el prefijo para los hijos
	childCount := node.GetChildCount()
	for i := 0; i < childCount; i++ {
		child := node.GetChild(i)
		newPrefix := prefix
		if isLast {
			newPrefix += "    "
		} else {
			newPrefix += "|   "
		}
		printVerticalNode(child, ruleNames, newPrefix, i == childCount-1)
	}
}

func readStdin() (string, error) {
	input, err := os.ReadFile("/dev/stdin")
	return string(input), err
}

// Genera el contenido DOT para Graphviz
func generateDotTree(node antlr.Tree, ruleNames []string) string {
	var builder strings.Builder
	builder.WriteString("digraph G {\n")
	builder.WriteString("  node [shape=box, fontname=\"Arial\", fontsize=10];\n")
	builder.WriteString("  edge [fontname=\"Arial\", fontsize=8];\n\n")

	generateDotNodes(&builder, node, ruleNames, 0)

	builder.WriteString("}\n")
	return builder.String()
}

// Genera los nodos y conexiones en formato DOT
func generateDotNodes(builder *strings.Builder, node antlr.Tree, ruleNames []string, id int) int {
	currentID := id

	// Obtener la etiqueta del nodo
	var label string
	switch n := node.(type) {
	case antlr.RuleNode:
		label = ruleNames[n.GetRuleContext().GetRuleIndex()]
	case antlr.TerminalNode:
		label = fmt.Sprintf("\"%s\"", n.GetText())
	default:
		label = fmt.Sprintf("%T", n)
	}

	// Escapar caracteres especiales para DOT
	label = strings.ReplaceAll(label, "\"", "\\\"")

	// Escribir el nodo actual
	builder.WriteString(fmt.Sprintf("  node%d [label=\"%s\"];\n", currentID, label))

	// Procesar hijos
	childCount := node.GetChildCount()
	nextID := currentID + 1

	for i := 0; i < childCount; i++ {
		child := node.GetChild(i)
		if child != nil {
			builder.WriteString(fmt.Sprintf("  node%d -> node%d;\n", currentID, nextID))
			nextID = generateDotNodes(builder, child, ruleNames, nextID)
		}
	}

	return nextID
}

// Genera el archivo SVG del árbol usando Graphviz
func generateTreeSVG(dotContent string, outputPath string) error {
	// Crear directorio si no existe
	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		return fmt.Errorf("no se pudo crear directorio: %v", err)
	}

	// Crear archivo DOT temporal
	dotFile := filepath.Join(os.TempDir(), "temp.dot")
	if err := os.WriteFile(dotFile, []byte(dotContent), 0644); err != nil {
		return fmt.Errorf("error al crear archivo DOT: %v", err)
	}
	defer os.Remove(dotFile)

	// Ejecutar Graphviz
	cmd := exec.Command("dot", "-Tsvg", dotFile, "-o", outputPath)
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("error al ejecutar Graphviz: %v\nSalida: %s", err, output)
	}

	return nil
}
