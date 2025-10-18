package autocorrection

func FixSpace(text string) string {
	space := false
	normalizedSpace := []rune{}

	for _, char := range text {
		if char == ' ' {
			if !space {
				space = true
				normalizedSpace = append(normalizedSpace, char)
				continue
			}
			continue
		}

		normalizedSpace = append(normalizedSpace, char)
		space = false
	}

	return string(normalizedSpace)
}
