.data
heap: .space 4096

false_string: .asciz "false"
true_string: .asciz "true"
.text
.global _start
_start:
    adr x10, heap
// Inicialización de stack y heap
MOV SP, #0x40000
ADR X10, heap
// Constant: 0
MOV x0, #0
STR x0, [SP, #-16]!
LDR x0, [SP], #16
STR x0, [SP, #-16]!
// Print statement
// String constant: === Prueba Básica Simplificada ===
STR x10, [SP, #-16]!
// Pushing char 61 to heap - (=)
MOV w0, #61
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 61 to heap - (=)
MOV w0, #61
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 61 to heap - (=)
MOV w0, #61
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 32 to heap - ( )
MOV w0, #32
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 80 to heap - (P)
MOV w0, #80
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 114 to heap - (r)
MOV w0, #114
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 117 to heap - (u)
MOV w0, #117
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 101 to heap - (e)
MOV w0, #101
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 98 to heap - (b)
MOV w0, #98
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 97 to heap - (a)
MOV w0, #97
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 32 to heap - ( )
MOV w0, #32
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 66 to heap - (B)
MOV w0, #66
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 195 to heap - (Ã)
MOV w0, #195
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 161 to heap - (¡)
MOV w0, #161
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 115 to heap - (s)
MOV w0, #115
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 105 to heap - (i)
MOV w0, #105
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 99 to heap - (c)
MOV w0, #99
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 97 to heap - (a)
MOV w0, #97
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 32 to heap - ( )
MOV w0, #32
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 83 to heap - (S)
MOV w0, #83
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 105 to heap - (i)
MOV w0, #105
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 109 to heap - (m)
MOV w0, #109
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 112 to heap - (p)
MOV w0, #112
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 108 to heap - (l)
MOV w0, #108
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 105 to heap - (i)
MOV w0, #105
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 102 to heap - (f)
MOV w0, #102
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 105 to heap - (i)
MOV w0, #105
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 99 to heap - (c)
MOV w0, #99
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 97 to heap - (a)
MOV w0, #97
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 100 to heap - (d)
MOV w0, #100
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 97 to heap - (a)
MOV w0, #97
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 32 to heap - ( )
MOV w0, #32
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 61 to heap - (=)
MOV w0, #61
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 61 to heap - (=)
MOV w0, #61
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 61 to heap - (=)
MOV w0, #61
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing null terminator to heap
MOV w0, #0
STRB w0, [x10]
LDR x0, [SP], #16
MOV X0, x0
BL print_string
// Constant: 42
MOV x0, #42
STR x0, [SP, #-16]!
LDR x0, [SP], #16
STR x0, [SP, #-16]!
// Constant: 3.14
MOVZ X0, #34079, LSL #0
MOVK X0, #20971, LSL #16
MOVK X0, #7864, LSL #32
MOVK X0, #16393, LSL #48
STR x0, [SP, #-16]!
LDR x0, [SP], #16
STR x0, [SP, #-16]!
MOV x0, #8
ADD x0, sp, x0
LDR x0, [x0, #0]
STR x0, [SP, #-16]!
// Constant: 42
MOV x0, #42
STR x0, [SP, #-16]!
LDR x1, [SP], #16
LDR x0, [SP], #16
SCVTF d0, x0
SCVTF d1, x1
FCMP D0, D1
MOV x0, #0
BEQ L2
B L3
L2:
MOV x0, #1
L3:
STR x0, [SP, #-16]!
LDR x0, [SP], #16
CMP x0, 0
BEQ L1
// Print statement
// String constant: OK entero
STR x10, [SP, #-16]!
// Pushing char 79 to heap - (O)
MOV w0, #79
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 75 to heap - (K)
MOV w0, #75
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 32 to heap - ( )
MOV w0, #32
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 101 to heap - (e)
MOV w0, #101
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 110 to heap - (n)
MOV w0, #110
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 116 to heap - (t)
MOV w0, #116
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 101 to heap - (e)
MOV w0, #101
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 114 to heap - (r)
MOV w0, #114
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 111 to heap - (o)
MOV w0, #111
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing null terminator to heap
MOV w0, #0
STRB w0, [x10]
LDR x0, [SP], #16
MOV X0, x0
BL print_string
B L0
L1:
L0:
MOV x0, #0
ADD x0, sp, x0
LDR x0, [x0, #0]
STR x0, [SP, #-16]!
// Constant: 3.0
MOVZ X0, #0, LSL #0
MOVK X0, #0, LSL #16
MOVK X0, #0, LSL #32
MOVK X0, #16392, LSL #48
STR x0, [SP, #-16]!
LDR x1, [SP], #16
LDR x0, [SP], #16
FMOV d0, x0
FMOV d1, x1
FCMP D0, D1
MOV x0, #0
BGT L6
B L7
L6:
MOV x0, #1
L7:
STR x0, [SP, #-16]!
LDR x0, [SP], #16
CMP x0, 0
BEQ L5
// Print statement
// String constant: OK decimal
STR x10, [SP, #-16]!
// Pushing char 79 to heap - (O)
MOV w0, #79
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 75 to heap - (K)
MOV w0, #75
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 32 to heap - ( )
MOV w0, #32
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 100 to heap - (d)
MOV w0, #100
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 101 to heap - (e)
MOV w0, #101
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 99 to heap - (c)
MOV w0, #99
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 105 to heap - (i)
MOV w0, #105
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 109 to heap - (m)
MOV w0, #109
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 97 to heap - (a)
MOV w0, #97
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 108 to heap - (l)
MOV w0, #108
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing null terminator to heap
MOV w0, #0
STRB w0, [x10]
LDR x0, [SP], #16
MOV X0, x0
BL print_string
B L4
L5:
L4:
// Boolean constant: true
MOV x0, #1
STR x0, [SP, #-16]!
LDR x0, [SP], #16
STR x0, [SP, #-16]!
MOV x0, #0
ADD x0, sp, x0
LDR x0, [x0, #0]
STR x0, [SP, #-16]!
LDR x0, [SP], #16
CMP x0, 0
BEQ L9
// Print statement
// String constant: OK booleano
STR x10, [SP, #-16]!
// Pushing char 79 to heap - (O)
MOV w0, #79
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 75 to heap - (K)
MOV w0, #75
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 32 to heap - ( )
MOV w0, #32
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 98 to heap - (b)
MOV w0, #98
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 111 to heap - (o)
MOV w0, #111
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 111 to heap - (o)
MOV w0, #111
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 108 to heap - (l)
MOV w0, #108
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 101 to heap - (e)
MOV w0, #101
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 97 to heap - (a)
MOV w0, #97
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 110 to heap - (n)
MOV w0, #110
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 111 to heap - (o)
MOV w0, #111
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing null terminator to heap
MOV w0, #0
STRB w0, [x10]
LDR x0, [SP], #16
MOV X0, x0
BL print_string
B L8
L9:
L8:
// Constant: 99
MOV x0, #99
STR x0, [SP, #-16]!
LDR x0, [SP], #16
MOV x1, #24
ADD x1, sp, x1
STR x0, [x1, #0]
MOV x0, #16
ADD x0, sp, x0
LDR x0, [x0, #0]
STR x0, [SP, #-16]!
// Constant: 99
MOV x0, #99
STR x0, [SP, #-16]!
LDR x1, [SP], #16
LDR x0, [SP], #16
SCVTF d0, x0
SCVTF d1, x1
FCMP D0, D1
MOV x0, #0
BEQ L12
B L13
L12:
MOV x0, #1
L13:
STR x0, [SP, #-16]!
LDR x0, [SP], #16
CMP x0, 0
BEQ L11
// Print statement
// String constant: OK asignación entero
STR x10, [SP, #-16]!
// Pushing char 79 to heap - (O)
MOV w0, #79
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 75 to heap - (K)
MOV w0, #75
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 32 to heap - ( )
MOV w0, #32
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 97 to heap - (a)
MOV w0, #97
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 115 to heap - (s)
MOV w0, #115
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 105 to heap - (i)
MOV w0, #105
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 103 to heap - (g)
MOV w0, #103
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 110 to heap - (n)
MOV w0, #110
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 97 to heap - (a)
MOV w0, #97
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 99 to heap - (c)
MOV w0, #99
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 105 to heap - (i)
MOV w0, #105
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 195 to heap - (Ã)
MOV w0, #195
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 179 to heap - (³)
MOV w0, #179
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 110 to heap - (n)
MOV w0, #110
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 32 to heap - ( )
MOV w0, #32
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 101 to heap - (e)
MOV w0, #101
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 110 to heap - (n)
MOV w0, #110
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 116 to heap - (t)
MOV w0, #116
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 101 to heap - (e)
MOV w0, #101
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 114 to heap - (r)
MOV w0, #114
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 111 to heap - (o)
MOV w0, #111
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing null terminator to heap
MOV w0, #0
STRB w0, [x10]
LDR x0, [SP], #16
MOV X0, x0
BL print_string
B L10
L11:
L10:
// Constant: 10
MOV x0, #10
STR x0, [SP, #-16]!
// Constant: 8
MOV x0, #8
STR x0, [SP, #-16]!
LDR x1, [SP], #16
LDR x0, [SP], #16
ADD x0, x0, x1
STR x0, [SP, #-16]!
LDR x0, [SP], #16
STR x0, [SP, #-16]!
MOV x0, #0
ADD x0, sp, x0
LDR x0, [x0, #0]
STR x0, [SP, #-16]!
// Constant: 15
MOV x0, #15
STR x0, [SP, #-16]!
LDR x1, [SP], #16
LDR x0, [SP], #16
SCVTF d0, x0
SCVTF d1, x1
FCMP D0, D1
MOV x0, #0
BEQ L16
B L17
L16:
MOV x0, #1
L17:
STR x0, [SP, #-16]!
LDR x0, [SP], #16
CMP x0, 0
BEQ L15
// Print statement
// String constant: OK suma
STR x10, [SP, #-16]!
// Pushing char 79 to heap - (O)
MOV w0, #79
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 75 to heap - (K)
MOV w0, #75
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 32 to heap - ( )
MOV w0, #32
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 115 to heap - (s)
MOV w0, #115
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 117 to heap - (u)
MOV w0, #117
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 109 to heap - (m)
MOV w0, #109
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 97 to heap - (a)
MOV w0, #97
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing null terminator to heap
MOV w0, #0
STRB w0, [x10]
LDR x0, [SP], #16
MOV X0, x0
BL print_string
B L14
L15:
L14:
// Constant: 5
MOV x0, #5
STR x0, [SP, #-16]!
// Constant: 10
MOV x0, #10
STR x0, [SP, #-16]!
LDR x1, [SP], #16
LDR x0, [SP], #16
SUB x0, x0, x1
STR x0, [SP, #-16]!
LDR x0, [SP], #16
STR x0, [SP, #-16]!
MOV x0, #0
ADD x0, sp, x0
LDR x0, [x0, #0]
STR x0, [SP, #-16]!
// Constant: 5
MOV x0, #5
STR x0, [SP, #-16]!
LDR x1, [SP], #16
LDR x0, [SP], #16
SCVTF d0, x0
SCVTF d1, x1
FCMP D0, D1
MOV x0, #0
BEQ L20
B L21
L20:
MOV x0, #1
L21:
STR x0, [SP, #-16]!
LDR x0, [SP], #16
CMP x0, 0
BEQ L19
// Print statement
// String constant: OK resta
STR x10, [SP, #-16]!
// Pushing char 79 to heap - (O)
MOV w0, #79
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 75 to heap - (K)
MOV w0, #75
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 32 to heap - ( )
MOV w0, #32
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 114 to heap - (r)
MOV w0, #114
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 101 to heap - (e)
MOV w0, #101
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 115 to heap - (s)
MOV w0, #115
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 116 to heap - (t)
MOV w0, #116
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 97 to heap - (a)
MOV w0, #97
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing null terminator to heap
MOV w0, #0
STRB w0, [x10]
LDR x0, [SP], #16
MOV X0, x0
BL print_string
B L18
L19:
L18:
// Constant: 10
MOV x0, #10
STR x0, [SP, #-16]!
// Constant: 10
MOV x0, #10
STR x0, [SP, #-16]!
LDR x1, [SP], #16
LDR x0, [SP], #16
SCVTF d0, x0
SCVTF d1, x1
FCMP D0, D1
MOV x0, #0
BEQ L24
B L25
L24:
MOV x0, #1
L25:
STR x0, [SP, #-16]!
LDR x0, [SP], #16
CMP x0, 0
BEQ L23
// Print statement
// String constant: OK igualdad
STR x10, [SP, #-16]!
// Pushing char 79 to heap - (O)
MOV w0, #79
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 75 to heap - (K)
MOV w0, #75
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 32 to heap - ( )
MOV w0, #32
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 105 to heap - (i)
MOV w0, #105
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 103 to heap - (g)
MOV w0, #103
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 117 to heap - (u)
MOV w0, #117
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 97 to heap - (a)
MOV w0, #97
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 108 to heap - (l)
MOV w0, #108
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 100 to heap - (d)
MOV w0, #100
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 97 to heap - (a)
MOV w0, #97
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 100 to heap - (d)
MOV w0, #100
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing null terminator to heap
MOV w0, #0
STRB w0, [x10]
LDR x0, [SP], #16
MOV X0, x0
BL print_string
B L22
L23:
L22:
// Boolean constant: true
MOV x0, #1
STR x0, [SP, #-16]!
// Boolean constant: true
MOV x0, #1
STR x0, [SP, #-16]!
LDR x1, [SP], #16
LDR x0, [SP], #16
AND X0, X0, X1
STR x0, [SP, #-16]!
LDR x0, [SP], #16
CMP x0, 0
BEQ L27
// Print statement
// String constant: OK lógica AND
STR x10, [SP, #-16]!
// Pushing char 79 to heap - (O)
MOV w0, #79
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 75 to heap - (K)
MOV w0, #75
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 32 to heap - ( )
MOV w0, #32
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 108 to heap - (l)
MOV w0, #108
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 195 to heap - (Ã)
MOV w0, #195
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 179 to heap - (³)
MOV w0, #179
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 103 to heap - (g)
MOV w0, #103
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 105 to heap - (i)
MOV w0, #105
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 99 to heap - (c)
MOV w0, #99
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 97 to heap - (a)
MOV w0, #97
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 32 to heap - ( )
MOV w0, #32
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 65 to heap - (A)
MOV w0, #65
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 78 to heap - (N)
MOV w0, #78
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 68 to heap - (D)
MOV w0, #68
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing null terminator to heap
MOV w0, #0
STRB w0, [x10]
LDR x0, [SP], #16
MOV X0, x0
BL print_string
B L26
L27:
L26:
LDR x0, [SP], #16
CMP x0, 0
BEQ L29
// Print statement
// String constant: OK lógica OR negada
STR x10, [SP, #-16]!
// Pushing char 79 to heap - (O)
MOV w0, #79
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 75 to heap - (K)
MOV w0, #75
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 32 to heap - ( )
MOV w0, #32
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 108 to heap - (l)
MOV w0, #108
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 195 to heap - (Ã)
MOV w0, #195
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 179 to heap - (³)
MOV w0, #179
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 103 to heap - (g)
MOV w0, #103
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 105 to heap - (i)
MOV w0, #105
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 99 to heap - (c)
MOV w0, #99
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 97 to heap - (a)
MOV w0, #97
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 32 to heap - ( )
MOV w0, #32
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 79 to heap - (O)
MOV w0, #79
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 82 to heap - (R)
MOV w0, #82
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 32 to heap - ( )
MOV w0, #32
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 110 to heap - (n)
MOV w0, #110
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 101 to heap - (e)
MOV w0, #101
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 103 to heap - (g)
MOV w0, #103
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 97 to heap - (a)
MOV w0, #97
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 100 to heap - (d)
MOV w0, #100
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 97 to heap - (a)
MOV w0, #97
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing null terminator to heap
MOV w0, #0
STRB w0, [x10]
LDR x0, [SP], #16
MOV X0, x0
BL print_string
B L28
L29:
L28:
// Print statement
// Constant: 42
MOV x0, #42
STR x0, [SP, #-16]!
LDR x0, [SP], #16
MOV X0, x0
BL print_integer
// Print statement
// String constant: Texto de prueba
STR x10, [SP, #-16]!
// Pushing char 84 to heap - (T)
MOV w0, #84
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 101 to heap - (e)
MOV w0, #101
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 120 to heap - (x)
MOV w0, #120
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 116 to heap - (t)
MOV w0, #116
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 111 to heap - (o)
MOV w0, #111
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 32 to heap - ( )
MOV w0, #32
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 100 to heap - (d)
MOV w0, #100
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 101 to heap - (e)
MOV w0, #101
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 32 to heap - ( )
MOV w0, #32
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 112 to heap - (p)
MOV w0, #112
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 114 to heap - (r)
MOV w0, #114
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 117 to heap - (u)
MOV w0, #117
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 101 to heap - (e)
MOV w0, #101
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 98 to heap - (b)
MOV w0, #98
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing char 97 to heap - (a)
MOV w0, #97
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Pushing null terminator to heap
MOV w0, #0
STRB w0, [x10]
LDR x0, [SP], #16
MOV X0, x0
BL print_string
// Print statement
// Boolean constant: true
MOV x0, #1
STR x0, [SP, #-16]!
LDR x0, [SP], #16
CMP X0, #0
BNE L30
ADR X0, false_string
BL print_string
B L31
L30:
ADR X0, true_string
BL print_string
L31:
// Fin del programa - exit syscall
MOV X0, #0
MOV X8, #93
SVC #0
MOV x0, #0
MOV x8, #93
SVC #0
MOV x0, #0
MOV x8, #93
SVC #0



// Standard Library

//--------------------------------------------------------------
// print_string - Prints a null-terminated string to stdout
//
// Input:
//   x0 - The address of the null-terminated string to print
//--------------------------------------------------------------
print_string:
    // Save link register and other registers we'll use
    stp     x29, x30, [sp, #-16]!
    stp     x19, x20, [sp, #-16]!
    
    // x19 will hold the string address
    mov     x19, x0
    
print_loop:
    // Load a byte from the string
    ldrb    w20, [x19]
    
    // Check if it's the null terminator (0)
    cbz     w20, print_done
    
    // Prepare for write syscall
    mov     x0, #1              // File descriptor: 1 for stdout
    mov     x1, x19             // Address of the character to print
    mov     x2, #1              // Length: 1 byte
    mov     x8, #64             // syscall: write (64 on ARM64)
    svc     #0                  // Make the syscall
    
    // Move to the next character
    add     x19, x19, #1
    
    // Continue the loop
    b       print_loop
    
print_done:
    // Restore saved registers
    ldp     x19, x20, [sp], #16
    ldp     x29, x30, [sp], #16
    ret



//--------------------------------------------------------------
// print_integer - Prints a signed integer to stdout
//
// Input:
//   x0 - The integer value to print
//--------------------------------------------------------------
print_integer:
    // Save registers
    stp x29, x30, [sp, #-16]!  // Save frame pointer and link register
    stp x19, x20, [sp, #-16]!  // Save callee-saved registers
    stp x21, x22, [sp, #-16]!
    stp x23, x24, [sp, #-16]!
    stp x25, x26, [sp, #-16]!
    stp x27, x28, [sp, #-16]!
    
    // Check if number is negative
    mov x19, x0                // Save original number
    cmp x19, #0                // Compare with zero
    bge positive_number        // Branch if greater or equal to zero
    
    // Handle negative number
    mov x0, #1                 // fd = 1 (stdout)
    adr x1, minus_sign         // Address of minus sign
    mov x2, #1                 // Length = 1
    mov w8, #64                // Syscall write
    svc #0
    
    neg x19, x19               // Make number positive
    
positive_number:
    // Prepare buffer for converting result to ASCII
    sub sp, sp, #32            // Reserve space on stack
    mov x22, sp                // x22 points to buffer
    
    // Initialize digit counter
    mov x23, #0                // Digit counter
    
    // Handle special case for zero
    cmp x19, #0
    bne convert_loop
    
    // If number is zero, just write '0'
    mov w24, #48               // ASCII '0'
    strb w24, [x22, x23]       // Store in buffer
    add x23, x23, #1           // Increment counter
    b print_result             // Skip conversion loop
    
convert_loop:
    // Divide the number by 10
    mov x24, #10
    udiv x25, x19, x24         // x25 = x19 / 10 (quotient)
    msub x26, x25, x24, x19    // x26 = x19 - (x25 * 10) (remainder)
    
    // Convert remainder to ASCII and store in buffer
    add x26, x26, #48          // Convert to ASCII ('0' = 48)
    strb w26, [x22, x23]       // Store digit in buffer
    add x23, x23, #1           // Increment digit counter
    
    // Prepare for next iteration
    mov x19, x25               // Quotient becomes the new number
    cbnz x19, convert_loop     // If number is not zero, continue
    
    // Reverse the buffer since digits are in reverse order
    mov x27, #0                // Start index
reverse_loop:
    sub x28, x23, x27          // x28 = length - current index
    sub x28, x28, #1           // x28 = length - current index - 1
    
    cmp x27, x28               // Compare indices
    bge print_result           // If crossed, finish reversing
    
    // Swap characters
    ldrb w24, [x22, x27]       // Load character from start
    ldrb w25, [x22, x28]       // Load character from end
    strb w25, [x22, x27]       // Store end character at start
    strb w24, [x22, x28]       // Store start character at end
    
    add x27, x27, #1           // Increment start index
    b reverse_loop             // Continue reversing
    
print_result:
    // Print the result
    mov x0, #1                 // fd = 1 (stdout)
    mov x1, x22                // Buffer address
    mov x2, x23                // Buffer length
    mov w8, #64                // Syscall write
    svc #0
    
    // Clean up and restore registers
    add sp, sp, #32            // Free buffer space
    ldp x27, x28, [sp], #16    // Restore callee-saved registers
    ldp x25, x26, [sp], #16
    ldp x23, x24, [sp], #16
    ldp x21, x22, [sp], #16
    ldp x19, x20, [sp], #16
    ldp x29, x30, [sp], #16    // Restore frame pointer and link register
    ret                        // Return to caller


minus_sign: .ascii "-"