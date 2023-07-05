// add.s
// Implementing addition in assembly language
// Adds 'a' and 'b', then returns the result
TEXT ·Add1(SB), $0-16
    MOVQ a+0(FP), AX
    MOVQ b+8(FP), CX
    ADDQ CX, AX
    MOVQ AX, ret+16(FP)
    RET
