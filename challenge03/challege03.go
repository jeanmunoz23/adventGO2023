package challenge03

func FindNaughtyStep(original string, modified string) string {
	var larger, shorter string
	if len(original) >= len(modified) {
		larger, shorter = original, modified
	} else {
		larger, shorter = modified, original
	}
	for i, char := range larger {
		if i < len(shorter) && char != rune(shorter[i]) {
			return string(char)
		}
	}

	if len(larger) > len(shorter) {
		return string(larger[len(shorter)])
	}

	return ""
}
