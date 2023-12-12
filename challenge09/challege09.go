package challenge09

func AdjustLights(lights []string) int {
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
