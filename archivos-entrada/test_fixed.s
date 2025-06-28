.data
print_fmt: .asciz "%s"
.Lstr_0:
	.asciz "\n==== Switch con break explícito ===="
.Lstr_6:
	.asciz "No se debería imprimir"
.Lstr_7:
	.asciz "Caso 2 - Se ejecuta este y debe detenerse"
.Lstr_8:
	.asciz "No debería ejecutarse si el break funciona"

.text
.global main

.global main
main:
	stp x29, x30, [sp, #-16]!
	mov x29, sp
	// Print statement
	adr x0, .Lstr_0
	// Print integer value
	mov x1, x0
	adr x0, print_fmt
	bl printf
	// Print newline
	mov x0, #10
	bl putchar
	// Variable declaration: numeroBreak
	mov x0, #2
	// Store variable numeroBreak at offset -8
	str x0, [x29, #-8]
	// Switch statement
	// Load variable numeroBreak from offset -8
	ldr x0, [x29, #-8]
	mov x9, x0 // Guardar valor del switch en x9
	mov x0, #1
	cmp x9, x0 // Comparar con valor del case
	beq .Lcase_0_3 // Saltar si igual
	mov x0, #2
	cmp x9, x0 // Comparar con valor del case
	beq .Lcase_1_4 // Saltar si igual
	mov x0, #3
	cmp x9, x0 // Comparar con valor del case
	beq .Lcase_2_5 // Saltar si igual
	b .Lswitch_end_1
.Lcase_0_3:
	// Print statement
	adr x0, .Lstr_6
	// Print integer value
	mov x1, x0
	adr x0, print_fmt
	bl printf
	// Print newline
	mov x0, #10
	bl putchar
	b .Lswitch_end_1
.Lcase_1_4:
	// Print statement
	adr x0, .Lstr_7
	// Print integer value
	mov x1, x0
	adr x0, print_fmt
	bl printf
	// Print newline
	mov x0, #10
	bl putchar
	// Print statement
	adr x0, .Lstr_8
	// Print integer value
	mov x1, x0
	adr x0, print_fmt
	bl printf
	// Print newline
	mov x0, #10
	bl putchar
	b .Lswitch_end_1
.Lcase_2_5:
	// Print statement
	adr x0, .Lstr_6
	// Print integer value
	mov x1, x0
	adr x0, print_fmt
	bl printf
	// Print newline
	mov x0, #10
	bl putchar
	b .Lswitch_end_1
.Lswitch_end_1:
	ldp x29, x30, [sp], #16
	ret