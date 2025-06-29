.data
heap: .space 4096

false_string: .asciz "false"
true_string: .asciz "true"
.text
.global _start
_start:
    adr x10, heap
// Print statement
// Visiting expresions
// Relational operation
// Visiting left expression: 
// Constant: 5
MOV x0, #5
STR x0, [SP, #-8]!
// Visiting right expression: 
// Constant: 5
MOV x0, #5
STR x0, [SP, #-8]!
// Popping operands
LDR x1, [SP], #8
LDR x0, [SP], #8
CMP x1, x0
// Setting condition flags
BNE L0
MOV x0, #0
STR x0, [SP, #-8]!
B L1
L0:
MOV x0, #1
STR x0, [SP, #-8]!
L1:
// Popping value to print
LDR x0, [SP], #8
CMP X0, #0
BNE L2
ADR X0, false_string
BL print_string
B L3
L2:
ADR X0, true_string
BL print_string
L3:
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
