.data
print_fmt: .asciz "%s"
.Lstr_1:
	.asciz "Switch simple"
.Lstr_11:
	.asciz "Lunes"
.Lstr_13:
	.asciz "Miércoles"
.Lstr_16:
	.asciz "Sábado"
.Lstr_0:
	.asciz "\n==== Switch/Case ===="
.Lstr_12:
	.asciz "Martes"
.Lstr_14:
	.asciz "Jueves"
.Lstr_15:
	.asciz "Viernes"
.Lstr_17:
	.asciz "Domingo"
.Lstr_18:
	.asciz "Día inválido"

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
	// Variable declaration: puntosSwitch
	mov x0, #0
	// Store variable puntosSwitch at offset -8
	str x0, [x29, #-8]
	// Print statement
	adr x0, .Lstr_1
	// Print integer value
	mov x1, x0
	adr x0, print_fmt
	bl printf
	// Print newline
	mov x0, #10
	bl putchar
	// Variable declaration: dia
	mov x0, #6
	// Store variable dia at offset -16
	str x0, [x29, #-16]
	// Switch statement
	// Load variable dia from offset -16
	ldr x0, [x29, #-16]
	mov x9, x0 // Guardar valor del switch en x9
	mov x0, #1
	cmp x9, x0 // Comparar con valor del case
	beq .Lcase_0_4 // Saltar si igual
	mov x0, #2
	cmp x9, x0 // Comparar con valor del case
	beq .Lcase_1_5 // Saltar si igual
	mov x0, #3
	cmp x9, x0 // Comparar con valor del case
	beq .Lcase_2_6 // Saltar si igual
	mov x0, #4
	cmp x9, x0 // Comparar con valor del case
	beq .Lcase_3_7 // Saltar si igual
	mov x0, #5
	cmp x9, x0 // Comparar con valor del case
	beq .Lcase_4_8 // Saltar si igual
	mov x0, #6
	cmp x9, x0 // Comparar con valor del case
	beq .Lcase_5_9 // Saltar si igual
	mov x0, #7
	cmp x9, x0 // Comparar con valor del case
	beq .Lcase_6_10 // Saltar si igual
	b .Lswitch_default_3
.Lcase_0_4:
	// Print statement
	adr x0, .Lstr_11
	// Print integer value
	mov x1, x0
	adr x0, print_fmt
	bl printf
	// Print newline
	mov x0, #10
	bl putchar
	// Assignment
	// Load variable puntosSwitch from offset -8
	ldr x0, [x29, #-8]
	mov x9, x0
	mov x0, #1
	mov x10, x0
	add x0, x9, x10
	// Store variable puntosSwitch at offset -8
	str x0, [x29, #-8]
	b .Lswitch_end_2
.Lcase_1_5:
	// Print statement
	adr x0, .Lstr_12
	// Print integer value
	mov x1, x0
	adr x0, print_fmt
	bl printf
	// Print newline
	mov x0, #10
	bl putchar
	b .Lswitch_end_2
.Lcase_2_6:
	// Print statement
	adr x0, .Lstr_13
	// Print integer value
	mov x1, x0
	adr x0, print_fmt
	bl printf
	// Print newline
	mov x0, #10
	bl putchar
	b .Lswitch_end_2
.Lcase_3_7:
	// Print statement
	adr x0, .Lstr_14
	// Print integer value
	mov x1, x0
	adr x0, print_fmt
	bl printf
	// Print newline
	mov x0, #10
	bl putchar
	b .Lswitch_end_2
.Lcase_4_8:
	// Print statement
	adr x0, .Lstr_15
	// Print integer value
	mov x1, x0
	adr x0, print_fmt
	bl printf
	// Print newline
	mov x0, #10
	bl putchar
	b .Lswitch_end_2
.Lcase_5_9:
	// Print statement
	adr x0, .Lstr_16
	// Print integer value
	mov x1, x0
	adr x0, print_fmt
	bl printf
	// Print newline
	mov x0, #10
	bl putchar
	b .Lswitch_end_2
.Lcase_6_10:
	// Print statement
	adr x0, .Lstr_17
	// Print integer value
	mov x1, x0
	adr x0, print_fmt
	bl printf
	// Print newline
	mov x0, #10
	bl putchar
	b .Lswitch_end_2
.Lswitch_default_3:
	// Print statement
	adr x0, .Lstr_18
	// Print integer value
	mov x1, x0
	adr x0, print_fmt
	bl printf
	// Print newline
	mov x0, #10
	bl putchar
.Lswitch_end_2:
	ldp x29, x30, [sp], #16
	ret