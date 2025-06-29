.data
heap: .space 4096

false_string: .asciz "false"
true_string: .asciz "true"
.text
.global _start
_start:
    adr x10, heap
// Declaración de variable: puntos
// Constant: 0
MOV x0, #0
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
LDR x0, [SP], #16
MOV X0, x0
BL print_string
// Declaración de variable: entero
// Constant: 42
MOV x0, #42
STR x0, [SP, #-16]!
// Declaración de variable: decimal
// Constant: 3.14
MOVZ X0, #34079, LSL #0
MOVK X0, #20971, LSL #16
MOVK X0, #7864, LSL #32
MOVK X0, #16393, LSL #48
STR x0, [SP, #-16]!
// Declaración de variable: booleano
// Boolean constant: true
MOV x0, #1
STR x0, [SP, #-16]!
// Relational operation
// Visiting left expression: 
MOV x0, #32
ADD x0, sp, x0
LDR x0, [x0, #0]
STR x0, [SP, #-16]!
// Visiting right expression: 
// Constant: 42
MOV x0, #42
STR x0, [SP, #-16]!
// Popping operands
LDR x1, [SP], #16
LDR x0, [SP], #16
CMP x1, x0
// Setting condition flags
BEQ L2
MOV x0, #0
STR x0, [SP, #-16]!
B L3
L2:
MOV x0, #1
STR x0, [SP, #-16]!
L3:
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
LDR x0, [SP], #16
MOV X0, x0
BL print_string
B L0
L1:
L0:
// Relational operation
// Visiting left expression: 
MOV x0, #16
ADD x0, sp, x0
LDR x0, [x0, #0]
STR x0, [SP, #-16]!
// Visiting right expression: 
// Constant: 3.0
MOVZ X0, #0, LSL #0
MOVK X0, #0, LSL #16
MOVK X0, #0, LSL #32
MOVK X0, #16392, LSL #48
STR x0, [SP, #-16]!
// Popping operands
LDR d1, [SP], #16
LDR d0, [SP], #16
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
LDR x0, [SP], #16
MOV X0, x0
BL print_string
B L4
L5:
L4:
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
