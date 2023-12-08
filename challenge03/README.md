En el taller de Santa, un elfo travieso ha estado jugando en la cadena de fabricación de regalos, añadiendo o eliminando un paso no planificado.

Tienes la secuencia original de pasos en la fabricación original y la secuencia modificada modified que puede incluir un paso extra o faltar un paso.

Tu tarea es escribir una función que identifique y devuelva el primer paso extra que se ha añadido o eliminado en la cadena de fabricación. Si no hay ninguna diferencia entre las secuencias, devuelve una cadena vacía.

```go
	original := "abcd"
	modified := "abcde"
	result := findNaughtyStep(original, modified) // "e"
	fmt.Println(result)

	original = "stepfor"
	modified = "stepor"
	result = findNaughtyStep(original, modified) // "f"
	fmt.Println(result)

	original = "abcde"
	modified = "abcde"
	result = findNaughtyStep(original, modified) // ""
	fmt.Println(result)
```
A tener en cuenta:

Siempre habrá un paso de diferencia o ninguno.
La modificación puede ocurrir en cualquier lugar de la cadena.
La secuencia original puede estar vacía