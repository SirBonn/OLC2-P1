# Instrucciones

1. Cargar los archivos directamente desde la interfaz de usuario.
2. Mencionar las modificaciones que se harán en el código. (si es necesario)
3. Mostrar la salida de la ejecución del código.

# Consideraciones:

- Los archivos _en principio_ no tienen errores.
- Algunos archivos cuentan con errores comentados. Descomentarlos uno a uno para ver el resultado.

```go
// El error siempre es acompañado del siguiente tipo de comentario:
// ! ERROR: Comentar esta línea para que el programa compile
```

Durante la calificacion, se verificaran algunas funcionalidades haciendo cambios que hagan un forzado de errores semanticos. Por ejemplo: Validacion de tipos, manejo de entornos o llamada de funciones que no fueron declaradas. 

- Cada archivo cuenta con un resumen de la puntuación obtenida en la evaluación de la salida del programa.

```go
    fmt.Println("\n=== Tabla de Resultados ===")
	fmt.Println("+--------------------------+--------+-------+")
	fmt.Println("| Característica           | Puntos | Total |")
	fmt.Println("+--------------------------+--------+-------+")
	fmt.Println("| Declaración de variables | ", puntosDeclaracion, "    | 2     |")
	fmt.Println("| Asignación de variables  | ", puntosAsignacion, "    | 2     |")
	fmt.Println("| Operaciones Aritméticas  | ", puntosOperacionesAritmeticas, "    | 2     |")
	fmt.Println("| Operaciones Relacionales | ", puntosOperacionesRelacionales, "    | 2     |")
	fmt.Println("| Operaciones Lógicas      | ", puntosOperacionesLogicas, "    | 2     |")
	fmt.Println("| fmt.Println              | ", puntosPrintln, "    | 2     |")
	fmt.Println("| Manejo de valor nulo     | ", puntosValorNulo, "    | 2     |")
	fmt.Println("| Opcionalidad del ;       | ", puntosPuntoYComa, "    | 2     |")
	fmt.Println("+--------------------------+--------+-------+")
	fmt.Println("| TOTAL                    | ", puntos, "   | 16    |")
	fmt.Println("+--------------------------+--------+-------+")
```

