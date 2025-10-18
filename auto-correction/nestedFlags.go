package autocorrection

func Devide(text string) string {
	flags := []string{}
	textCopy := text

	for {
		start := -2
		end := -2

		for i, char := range textCopy {
			if char == '(' {
				start = i
			} else if char == ')' && start != -2 {
				end = i
				break
			}
		}

		if start == -2 || end == -2 {
			break
		}

		flag := textCopy[start+1 : end]
		flags = append(flags, "~"+flag+"`")

		textCopy = textCopy[:start] + textCopy[end+1:]
	}

	str := ""

	for i := len(flags) - 1; i >= 0; i-- {
		str += flags[i]
	}

	return str
}

func Insert(text string) string {
	for {
		stack := []rune{}
		start := -1
		end := -1

		for i, char := range text {
			if char == '(' {
				stack = append(stack, char)
				if start != -1 {
					continue
				}
				start = i
			} else if char == ')' && start != -1 {
				stack = stack[:len(stack)-1]
				end = i
				if len(stack) == 0 {
					break
				}
			}
		}

		if start == -1 || end == -1 {
			break
		}

		flags := Devide(text[start : end+1])

		text = text[:start] + flags + text[end+1:]
	}

	text = TildeToBrace(text)

	return text
}

func TildeToBrace(text string) string {
	newText := []rune(text)

	for i := range newText {
		switch newText[i] {
		case '~':
			newText[i] = '('
		case '`':
			newText[i] = ')'
		}
	}

	return string(newText)
}
