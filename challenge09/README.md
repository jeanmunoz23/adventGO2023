Están encendiendo las luces de Navidad 🎄 en la ciudad y, como cada año, ¡hay que arreglarlas!

Las luces son de dos colores: 🔴 y 🟢 . Para que el efecto sea el adecuado, siempre deben estar alternadas. Es decir, si la primera luz es roja, la segunda debe ser verde, la tercera roja, la cuarta verde, etc.

Nos han pedido que escribamos una función adjustLights que, dado un array de strings con el color de cada luz, devuelva el número mínimo de luces que hay que cambiar para que estén los colores alternos.

```go
  	lights := []string{"🟢", "🔴", "🟢", "🟢", "🟢"}
	result := adjustLights(lights)
	fmt.Println(result)
	// -> 1 (cambias la cuarta luz a 🔴)

	lights = []string{"🔴", "🔴", "🟢", "🟢", "🔴"}
	result = adjustLights(lights)
	fmt.Println(result) // -> 2 (cambias la segunda luz a 🟢 y la tercera a 🔴)

	lights = []string{"🟢", "🔴", "🟢", "🔴", "🟢"}
	result = adjustLights(lights)
	fmt.Println(result) // -> 0 (ya están alternadas)

	lights = []string{"🔴", "🔴", "🔴"}
	result = adjustLights(lights)
	fmt.Println(result) // -> 1 (cambias la segunda luz a 🟢)
```