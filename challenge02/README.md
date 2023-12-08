En el taller de Santa, los elfos tienen una lista de regalos que desean fabricar y un conjunto limitado de materiales.

Los regalos son cadenas de texto y los materiales son caracteres. Tu tarea es escribir una función que, dada una lista de regalos y los materiales disponibles, devuelva una lista de los regalos que se pueden fabricar.

Un regalo se puede fabricar si contamos con todos los materiales necesarios para fabricarlo.
```go
gifts := []string{"tren", "oso", "pelota"}
	materials := "tronesa"
	result := manufacture(gifts, materials)
	fmt.Println(result) // ["tren", "oso"]
	// "tren" SÍ porque sus letras están en "tronesa"
	// "oso" SÍ porque sus letras están en "tronesa"
	// "pelota" NO porque sus letras NO están en "tronesa"

	gifts = []string{"juego", "puzzle"}
	materials = "jlepuz"
	result = manufacture(gifts, materials) // ["puzzle"]
	fmt.Println(result)                    //

	gifts = []string{"libro", "ps5"}
	materials = "psli"
	result = manufacture(gifts, materials) // []
	fmt.Println(result)         
```