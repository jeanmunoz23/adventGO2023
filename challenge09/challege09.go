package main

import (
	"fmt"
)

func adjustLights(lights []string) int {
	changes := 0

	for i := 0; i < len(lights); i++ {
		if i%2 == 0 && lights[i] != "🟢" {
			changes++
		} else if i%2 == 1 && lights[i] != "🔴" {
			changes++
		}
	}

	// Si el número de cambios es mayor a la mitad del número de luces,
	// es más eficiente cambiar las luces que no se han contado.
	if changes > len(lights)/2 {
		changes = len(lights) - changes
	}

	return changes
}

func main() {
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
}
