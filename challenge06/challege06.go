package challenge06

func MaxDistance(movements string) int {
	derecha := 0
	izquierda := 0
	extras := 0
	result := 0
	for _, move := range movements {
		if string(move) == "<" {
			derecha++
		} else if string(move) == ">" {
			izquierda++
		} else {
			extras++
		}
	}
	if derecha > izquierda {
		result = derecha - izquierda + extras
	} else {
		result = izquierda + extras - derecha
	}
	return result
}
