package arm

import (
	"fmt"
	"math"
	"strings"
)

type StackObjectType int

const (
	IntType StackObjectType = iota
	FloatType
	StringType
	BoolType
)

type StackObject struct {
	Type   StackObjectType
	Length int
	Depth  int
	Id     *string
}

type ArmGenerator struct {
	instructions []string
	stdLib       *StandardLibrary
	stack        []StackObject
	depth        int
	labelCounter int
}

func NewArmGenerator() *ArmGenerator {
	return &ArmGenerator{
		instructions: []string{},
		stdLib:       NewStandardLibrary(),
		stack:        []StackObject{},
		depth:        0,
		labelCounter: 0,
	}
}

// Codigo agregado 28/06
// En el ArmGenerator
func (a *ArmGenerator) InitializeHeapPointer() {
	a.instructions = append(a.instructions, "ADR HP, heap")
	a.instructions = append(a.instructions, "MOV X10, HP") // Mantener consistencia
}

func (a *ArmGenerator) Push(rs string) {
	a.instructions = append(a.instructions, "STR "+rs+", [SP, #-16]!") // Alinear a 16 bytes
}

func (a *ArmGenerator) Pop(rd string) {
	a.instructions = append(a.instructions, "LDR "+rd+", [SP], #16") // Alinear a 16 bytes
}

func (ag *ArmGenerator) DebugStack() {
	ag.Comment("=== DEBUG STACK ===")
	for i, obj := range ag.stack {
		id := "nil"
		if obj.Id != nil {
			id = *obj.Id
		}
		ag.Comment(fmt.Sprintf("Pos %d: Tipo=%v Profundidad=%d ID=%s",
			i, obj.Type, obj.Depth, id))
	}
}

func (ag *ArmGenerator) FindObject(id string) (int, StackObject, bool) {
	for i := len(ag.stack) - 1; i >= 0; i-- {
		if ag.stack[i].Id != nil && *ag.stack[i].Id == id {
			// Calcular offset
			offset := 0
			for j := i + 1; j < len(ag.stack); j++ {
				offset += ag.stack[j].Length
			}
			return offset, ag.stack[i], true
		}
	}
	return 0, StackObject{}, false
}

//------------------------------------------

//var int labelCounter = 0;

func (ag *ArmGenerator) PushObject(obj StackObject) {
	ag.stack = append(ag.stack, obj)
}

func (ag *ArmGenerator) PopObject(rd string) StackObject {
	obj := ag.stack[len(ag.stack)-1]
	ag.stack = ag.stack[:len(ag.stack)-1]
	ag.Pop(rd)
	return obj
}

func (ag *ArmGenerator) TopObject() StackObject {
	if len(ag.stack) == 0 {
		panic("TopObject() called on empty stack")
	}
	return ag.stack[len(ag.stack)-1]
}

func (ag *ArmGenerator) commentStack() {
	stackStr := "Stack: "
	for _, obj := range ag.stack {
		id := "nil"
		if obj.Id != nil {
			id = *obj.Id
		}
		stackStr += fmt.Sprintf("[%v, %d, %s] ", obj.Type, obj.Depth, id)
	}
	ag.Comment(stackStr)
}

func (ag *ArmGenerator) PushConstant(obj StackObject, value interface{}) {
	switch obj.Type {
	case IntType:
		ag.Mov(X0, value.(int))
		ag.Push(X0)
	case FloatType:
		floatBits := math.Float64bits(value.(float64))
		for i := 0; i < 4; i++ {
			part := (floatBits >> (i * 16)) & 0xFFFF
			if i == 0 {
				ag.instructions = append(ag.instructions, fmt.Sprintf("MOVZ X0, #%d, LSL #0", part))
			} else {
				ag.instructions = append(ag.instructions, fmt.Sprintf("MOVK X0, #%d, LSL #%d", part, i*16))
			}
		}
		ag.Push(X0)
	case StringType:
		ag.Push(HP)
		for _, b := range []byte(value.(string)) {
			ag.Comment(fmt.Sprintf("Pushing char %d to heap - (%c)", b, b))
			ag.Mov("w0", int(b))
			ag.Strb("w0", HP)
			ag.Mov(X0, 1)
			ag.Add(HP, HP, X0)
		}
		ag.Comment("Pushing null terminator to heap")
		ag.Mov("w0", 0)
		ag.Strb("w0", HP)
	case BoolType:
		val := 0
		if value.(bool) {
			val = 1
		}
		ag.Mov(X0, val)
		ag.Push(X0)
	}

	ag.PushObject(obj)
}

func (ag *ArmGenerator) IntObject() StackObject {
	return StackObject{IntType, 8, ag.depth, nil}
}

func (ag *ArmGenerator) FloatObject() StackObject {
	return StackObject{FloatType, 8, ag.depth, nil}
}

func (ag *ArmGenerator) StringObject() StackObject {
	return StackObject{StringType, 8, ag.depth, nil}
}

func (ag *ArmGenerator) BoolObject() StackObject {
	return StackObject{BoolType, 8, ag.depth, nil}
}

func (ag *ArmGenerator) CloneObject(obj StackObject) StackObject {
	return StackObject{obj.Type, obj.Length, obj.Depth, obj.Id}
}

func (ag *ArmGenerator) NewScope() {
	ag.depth++
}

func (ag *ArmGenerator) EndScope() int {
	byteOffset := 0
	for i := len(ag.stack) - 1; i >= 0; i-- {
		if ag.stack[i].Depth == ag.depth {
			byteOffset += ag.stack[i].Length
			ag.stack = append(ag.stack[:i], ag.stack[i+1:]...)
		} else {
			break
		}
	}
	ag.depth--
	return byteOffset
}

func (ag *ArmGenerator) TagObject(id string) {
	last := ag.stack[len(ag.stack)-1]
	last.Id = &id
	ag.stack[len(ag.stack)-1] = last
}

func (ag *ArmGenerator) GetObject(id string) (int, StackObject) {
	byteOffset := 0
	for i := len(ag.stack) - 1; i >= 0; i-- {
		if ag.stack[i].Id != nil && *ag.stack[i].Id == id {
			return byteOffset, ag.stack[i]
		}
		byteOffset += ag.stack[i].Length
	}
	panic("Object " + id + " not found in stack")
}

func (a *ArmGenerator) Add(rd string, rs1 string, rs2 string) {
	a.instructions = append(a.instructions, fmt.Sprintf("ADD %s, %s, %s", rd, rs1, rs2))
}

func (a *ArmGenerator) Sub(rd, rs1, rs2 string) {
	a.instructions = append(a.instructions, fmt.Sprintf("SUB %s, %s, %s", rd, rs1, rs2))
}

func (a *ArmGenerator) Mul(rd, rs1, rs2 string) {
	a.instructions = append(a.instructions, fmt.Sprintf("MUL %s, %s, %s", rd, rs1, rs2))
}

func (a *ArmGenerator) Div(rd, rs1, rs2 string) {
	a.instructions = append(a.instructions, fmt.Sprintf("DIV %s, %s, %s", rd, rs1, rs2))
}

func (a *ArmGenerator) Addi(rd, rs1 string, imm int) {
	a.instructions = append(a.instructions, fmt.Sprintf("ADDI %s, %s, #%d", rd, rs1, imm))
}

func (a *ArmGenerator) Str(rs1, rs2 string, offset int) {
	a.instructions = append(a.instructions, fmt.Sprintf("STR %s, [%s, #%d]", rs1, rs2, offset))
}

func (a *ArmGenerator) Strb(rs1, rs2 string) {
	a.instructions = append(a.instructions, fmt.Sprintf("STRB %s, [%s]", rs1, rs2))
}

func (a *ArmGenerator) Ldr(rd, rs1 string, offset int) {
	a.instructions = append(a.instructions, fmt.Sprintf("LDR %s, [%s, #%d]", rd, rs1, offset))
}

func (a *ArmGenerator) Mov(rd string, imm int) {
	a.instructions = append(a.instructions, fmt.Sprintf("MOV %s, #%d", rd, imm))
}

// cod comentado 28/06
/*func (a *ArmGenerator) Push(rs string) {
	a.instructions = append(a.instructions, "STR "+rs+", [SP, #-8]!")
}

func (a *ArmGenerator) Pop(rd string) {
	a.instructions = append(a.instructions, "LDR "+rd+", [SP], #8")
}*/

func (a *ArmGenerator) Svc() {
	a.instructions = append(a.instructions, "SVC #0")
}

func (c *ArmGenerator) Scvtf(rd, rs string) {
	c.instructions = append(c.instructions, fmt.Sprintf("SCVTF %s, %s", rd, rs))
}

func (c *ArmGenerator) Fmov(rd, rs string) {
	c.instructions = append(c.instructions, fmt.Sprintf("FMOV %s, %s", rd, rs))
}

func (c *ArmGenerator) Fadd(rd, rs1, rs2 string) {
	c.instructions = append(c.instructions, fmt.Sprintf("FADD %s, %s, %s", rd, rs1, rs2))
}

func (c *ArmGenerator) Fsub(rd, rs1, rs2 string) {
	c.instructions = append(c.instructions, fmt.Sprintf("FSUB %s, %s, %s", rd, rs1, rs2))
}

func (c *ArmGenerator) Fmul(rd, rs1, rs2 string) {
	c.instructions = append(c.instructions, fmt.Sprintf("FMUL %s, %s, %s", rd, rs1, rs2))
}

func (c *ArmGenerator) Fdiv(rd, rs1, rs2 string) {
	c.instructions = append(c.instructions, fmt.Sprintf("FDIV %s, %s, %s", rd, rs1, rs2))
}

func (gen *ArmGenerator) B(label string) {
	gen.instructions = append(gen.instructions, fmt.Sprintf("B %s", label))
}

func (gen *ArmGenerator) Bl(label string) {
	gen.instructions = append(gen.instructions, fmt.Sprintf("BL %s", label))
}

func (gen *ArmGenerator) Cbz(rs, label string) {
	gen.instructions = append(gen.instructions, fmt.Sprintf("CBZ %s, %s", rs, label))
}

func (gen *ArmGenerator) Cbnz(rs, label string) {
	gen.instructions = append(gen.instructions, fmt.Sprintf("CBNZ %s, %s", rs, label))
}

func (a *ArmGenerator) Cmp(rs1, rs2 string) {
	a.instructions = append(a.instructions, fmt.Sprintf("CMP %s, %s", rs1, rs2))
}

func (gen *ArmGenerator) Beq(label string) {
	gen.instructions = append(gen.instructions, fmt.Sprintf("BEQ %s", label))
}

func (gen *ArmGenerator) Bne(label string) {
	gen.instructions = append(gen.instructions, fmt.Sprintf("BNE %s", label))
}

func (gen *ArmGenerator) Bgt(label string) {
	gen.instructions = append(gen.instructions, fmt.Sprintf("BGT %s", label))
}

func (gen *ArmGenerator) Blt(label string) {
	gen.instructions = append(gen.instructions, fmt.Sprintf("BLT %s", label))
}

func (gen *ArmGenerator) Bge(label string) {
	gen.instructions = append(gen.instructions, fmt.Sprintf("BGE %s", label))
}

func (gen *ArmGenerator) Ble(label string) {
	gen.instructions = append(gen.instructions, fmt.Sprintf("BLE %s", label))
}

func (gen *ArmGenerator) Ret() {
	gen.instructions = append(gen.instructions, "RET")
}

func (gen *ArmGenerator) RetReg(rs string) {
	gen.instructions = append(gen.instructions, fmt.Sprintf("RET %s", rs))
}

func (gen *ArmGenerator) SetLabel(label string) {
	gen.instructions = append(gen.instructions, fmt.Sprintf("%s:", label))
}

func (gen *ArmGenerator) GetLabel() string {
	label := fmt.Sprintf("L%d", gen.labelCounter)
	gen.labelCounter++
	return label
}

func (a *ArmGenerator) EndProgram() {
	a.Mov(X0, 0)
	a.Mov(X8, 93)
	a.Svc()
}

// -------> Cod para imprimir cada uno de los tipos de datos <--------
func (a *ArmGenerator) PrintInteger(rs string) {
	a.stdLib.Use("print_integer")
	a.instructions = append(a.instructions, fmt.Sprintf("MOV X0, %s", rs))
	a.instructions = append(a.instructions, "BL print_integer")
}

func (a *ArmGenerator) PrintString(rs string) {
	a.stdLib.Use("print_string")
	a.instructions = append(a.instructions, fmt.Sprintf("MOV X0, %s", rs))
	a.instructions = append(a.instructions, "BL print_string")
}

func (a *ArmGenerator) PrintFloat() {
	a.stdLib.Use("print_integer")
	a.stdLib.Use("print_double")
	a.instructions = append(a.instructions, "BL print_double")
}

func (c *ArmGenerator) PrintBool() {
	// Assume value is in X0
	labelTrue := c.GetLabel()
	labelEnd := c.GetLabel()

	c.instructions = append(c.instructions, "CMP X0, #0")
	c.instructions = append(c.instructions, fmt.Sprintf("BNE %s", labelTrue))

	// false
	c.stdLib.Use("print_string")
	c.instructions = append(c.instructions, "ADR X0, false_string")
	c.instructions = append(c.instructions, "BL print_string")
	c.instructions = append(c.instructions, fmt.Sprintf("B %s", labelEnd))

	// true
	c.instructions = append(c.instructions, fmt.Sprintf("%s:", labelTrue))
	c.stdLib.Use("print_string")
	c.instructions = append(c.instructions, "ADR X0, true_string")
	c.instructions = append(c.instructions, "BL print_string")

	// end
	c.instructions = append(c.instructions, fmt.Sprintf("%s:", labelEnd))
}

func (a *ArmGenerator) LoadLabel(rd string, label string) {
	a.instructions = append(a.instructions, fmt.Sprintf("ADR %s, %s", rd, label))
}

func (a *ArmGenerator) Comment(comment string) {
	a.instructions = append(a.instructions, "// "+comment)
}

func (a *ArmGenerator) String() string {
	var sb strings.Builder
	sb.WriteString(".data\n")
	sb.WriteString("heap: .space 4096\n\n")
	sb.WriteString("false_string: .asciz \"false\"\n")
	sb.WriteString("true_string: .asciz \"true\"\n")
	sb.WriteString(".text\n")
	sb.WriteString(".global _start\n")
	sb.WriteString("_start:\n")
	sb.WriteString("    adr x10, heap\n")

	a.EndProgram()

	for _, instr := range a.instructions {
		sb.WriteString(instr + "\n")
	}

	sb.WriteString("\n\n\n// Standard Library\n")
	sb.WriteString(a.stdLib.GetFunctionDefinitions())

	return sb.String()
}
