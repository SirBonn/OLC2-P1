.data
.align 4
print_fmt: .asciz "%d"
print_str_fmt: .asciz "%s"
print_space_fmt: .asciz "%c"
.Lstr_0:
	.asciz "For como while simple"
.Lstr_3:
	.asciz "i ="

.text
.align 4
.global main

main:
	stp x29, x30, [sp, #-64]!
	mov x29, sp
	stp x19, x20, [sp, #16]
	str x21, [sp, #32]
	stp x19, x20, [sp, #-16]!
	adr x0, .Lstr_0
	adrp x19, print_str_fmt
	add x19, x19, :lo12:print_str_fmt
	mov x20, x0
	mov x0, x19
	mov x1, x20
	bl printf
	mov x0, #10
	bl putchar
	ldp x19, x20, [sp], #16
	mov x0, #0
	str x0, [x29, #40] // Store i
	mov x0, #0
	str x0, [x29, #48] // Store suma
.Lfor_start_1:
	ldr x0, [x29, #40] // Load i
	mov x19, x0 // Guardar operando izquierdo
	mov x0, #5
	mov x20, x0 // Guardar operando derecho
	cmp x19, x20
	cset x0, lt
	cbz x0, .Lfor_end_2
	stp x19, x20, [sp, #-16]!
	adr x0, .Lstr_3
	adrp x19, print_str_fmt
	add x19, x19, :lo12:print_str_fmt
	mov x20, x0
	mov x0, x19
	mov x1, x20
	mov x2, #' '
	bl printf
	ldr x0, [x29, #40] // Load i
	adrp x19, print_fmt
	add x19, x19, :lo12:print_fmt
	mov x20, x0
	mov x0, x19
	mov x1, x20
	bl printf
	mov x0, #10
	bl putchar
	ldp x19, x20, [sp], #16
	ldr x0, [x29, #48] // Load suma
	mov x19, x0 // Guardar operando izquierdo
	ldr x0, [x29, #40] // Load i
	mov x20, x0 // Guardar operando derecho
	add x0, x19, x20
	str x0, [x29, #48] // Store to suma
	ldr x0, [x29, #40] // Load i
	mov x19, x0 // Guardar operando izquierdo
	mov x0, #1
	mov x20, x0 // Guardar operando derecho
	add x0, x19, x20
	str x0, [x29, #40] // Store to i
	b .Lfor_start_1
.Lfor_end_2:
	mov x0, #0
	ldr x21, [sp, #32]
	ldp x19, x20, [sp, #16]
	ldp x29, x30, [sp], #64
	ret