package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"main/arm"
	"main/errors"
	compiler "main/parser"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/antlr4-go/antlr/v4"
)

func main() {
	a := app.NewWithID("ocl2.vlang.cherry")
	w := a.NewWindow("V-Lang Cherry")

	codeEntry := widget.NewMultiLineEntry()
	codeEntry.SetPlaceHolder("Escribe tu código aquí...")

	consoleOutput := widget.NewMultiLineEntry()
	consoleOutput.Wrapping = fyne.TextWrapWord
	consoleOutput.Disable()

	positionLabel := widget.NewLabel("Línea: 1, Columna: 1")
	updateCursorPosition := func() {
		pos := codeEntry.CursorRow
		col := codeEntry.CursorColumn
		positionLabel.SetText("Línea: " + strconv.Itoa(pos+1) + ", Columna: " + strconv.Itoa(col+1))
	}
	codeEntry.OnChanged = func(s string) { updateCursorPosition() }
	codeEntry.OnCursorChanged = func() { updateCursorPosition() }

	compileButton := widget.NewButton("Compilar", func() {
		inputCode := codeEntry.Text

		lexicalErrorListener := errors.NewLexicalErrorListener()
		lexer := compiler.NewVlangLexer(antlr.NewInputStream(inputCode))
		lexer.RemoveErrorListeners()
		lexer.AddErrorListener(lexicalErrorListener)

		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		parser := compiler.NewVlangParser(stream)
		parser.BuildParseTrees = true
		syntaxErrorListener := errors.NewSyntaxErrorListener()
		parser.RemoveErrorListeners()
		parser.AddErrorListener(syntaxErrorListener)

		tree := parser.Programa()

		if syntaxErrorListener.HasErrors() {
			consoleOutput.SetText(syntaxErrorListener.GetErrors())
			return
		}

		visitorArm := arm.NewArmVisitor()
		arm.Walk(visitorArm, tree)
		asmCode := visitorArm.Generator.String()
		consoleOutput.SetText(asmCode)

		err := os.WriteFile("./assembler/program.s", []byte(asmCode), 0644)
		if err != nil {
			dialog.ShowError(err, w)
		}
	})

	menu := fyne.NewMainMenu(
		fyne.NewMenu("Archivo",
			fyne.NewMenuItem("Nuevo", func() {
				codeEntry.SetText("")
				consoleOutput.SetText("")
			}),
			fyne.NewMenuItem("Abrir", func() {
				fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
					if err != nil {
						dialog.ShowError(err, w)
						return
					}
					if reader == nil {
						return
					}
					defer reader.Close()
					data, err := ioutil.ReadAll(reader)
					if err != nil {
						dialog.ShowError(err, w)
						return
					}
					codeEntry.SetText(string(data))
				}, w)
				fd.SetFilter(storage.NewExtensionFileFilter([]string{".vch"}))
				fd.Show()
			}),
			fyne.NewMenuItem("Guardar", func() {
				fd := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
					if err != nil {
						dialog.ShowError(err, w)
						return
					}
					if writer == nil {
						return
					}
					defer writer.Close()
					_, err = io.WriteString(writer, codeEntry.Text)
					if err != nil {
						dialog.ShowError(err, w)
					}
				}, w)
				fd.SetFileName("source.vch")
				fd.SetFilter(storage.NewExtensionFileFilter([]string{".vch"}))
				fd.Show()
			}),
		),
	)
	w.SetMainMenu(menu)

	split := container.NewHSplit(
		codeEntry,
		consoleOutput,
	)
	split.SetOffset(0.5)

	mainContainer := container.NewBorder(
		nil,
		container.NewVBox(compileButton, positionLabel),
		nil,
		nil,
		split,
	)

	w.SetContent(mainContainer)
	w.Resize(fyne.NewSize(1200, 700))
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
