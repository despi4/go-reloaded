package autocorrection

import (
	"regexp"
	"strings"
	"unicode"
)

var (
	symbolsBeforeFlagHex = regexp.MustCompile(`(?i)([^)^\s+]*)\s*\(hex?\)`)
	symbolsBeforeFlagBin = regexp.MustCompile(`(?i)([^"^'^)^\s+]*)\s*\(bin?\)`)
	hexFlag              = regexp.MustCompile(`(?i)\s*\(hex\s*,?\s*\d*\)`)
	binFlag              = regexp.MustCompile(`(?i)\s*\(bin\s*,?\s*\d*\)`)
	spaceInFlag          = regexp.MustCompile(`(?i)\(\s*(hex|bin)\s*[^)]*\)`)
)

func NumericNavigator(text string) string {
	text = NumericTrimExtraSpace(text)
	text = NumericDelExtraSymbols(text)
	text = DelFlagBin(text)
	text = DelFlagHex(text)

	return text
}

func NumericTrimExtraSpace(text string) string {
	trimSpace := spaceInFlag.ReplaceAllStringFunc(text, func(s string) string {
		s = strings.TrimSpace(s[1 : len(s)-1])

		return "(" + s + ")"
	})

	return trimSpace
}

func NumericDelExtraSymbols(text string) string {
	text = hexFlag.ReplaceAllString(text, " (hex) ")
	text = binFlag.ReplaceAllString(text, " (bin) ")

	text = strings.TrimSpace(text)

	return text
}

// func DelFlagHex(text string) string {
// 	flagDelete := symbolsBeforeFlagHex.ReplaceAllStringFunc(text, func(s string) string {
// 		stack := []rune{}
// 		unesSymbol := false

// 		for i, char := range s {
// 			r := []rune(strings.ToLower(string(char)))[0]

// 			if i == len(s)-5 {
// 				break
// 			}

// 			if r >= 'a' && r <= 'f' || r >= '0' && r <= '9' {
// 				stack = append(stack, char)
// 			} else if r != ' ' {
// 				unesSymbol = true
// 				stack = append(stack, char)
// 			}
// 		}

// 		fmt.Println(string(stack))

// 		if unesSymbol || len(stack) == 0 {
// 			return string(stack)
// 		}

// 		return string(stack) + " (hex)"
// 	})

// 	return flagDelete
// }

// func DelFlagBin(text string) string {
// 	flagDelete := symbolsBeforeFlagBin.ReplaceAllStringFunc(text, func(s string) string {
// 		stack := []rune{}
// 		unesSymbol := false

// 		for i, char := range s {
// 			if i == len(s)-5 {
// 				break
// 			}

// 			if char >= '0' && char <= '1' || char == '(' || char == ')' || char >= 'a' && char <= 'f' || char >= 'A' && char <= 'F' {
// 				stack = append(stack, char)
// 			} else if char != ' ' {
// 				unesSymbol = true
// 				stack = append(stack, char)
// 			}
// 		}

// 		if unesSymbol || len(stack) == 0 {
// 			return string(stack)
// 		}

// 		return string(stack) + " (bin)"
// 	})

// 	return flagDelete
// }

func DelFlagHex(text string) string {
	return symbolsBeforeFlagHex.ReplaceAllStringFunc(text, func(match string) string {
		submatch := symbolsBeforeFlagHex.FindStringSubmatch(match)
		if len(submatch) < 2 {
			return match
		}
		prefix := submatch[1]
		trimmed := strings.TrimRightFunc(prefix, unicode.IsSpace)

		// Должен быть непустой и содержать ТОЛЬКО hex-символы
		if trimmed == "" {
			return prefix // удаляем флаг
		}

		for _, r := range trimmed {
			if !((r >= '0' && r <= '9') || (r >= 'a' && r <= 'f') || (r >= 'A' && r <= 'F')) {
				// Недопустимый символ — включая '(', ')', буквы g-z и т.д.
				return prefix // удаляем флаг
			}
		}

		// Только если ВСЕ символы — валидные hex → оставляем флаг
		return trimmed + " (hex)"
	})
}

func DelFlagBin(text string) string {
	return symbolsBeforeFlagBin.ReplaceAllStringFunc(text, func(match string) string {
		submatch := symbolsBeforeFlagBin.FindStringSubmatch(match)
		if len(submatch) < 2 {
			return match
		}
		prefix := submatch[1]
		trimmed := strings.TrimRightFunc(prefix, unicode.IsSpace)

		if trimmed == "" {
			return prefix
		}

		for _, r := range trimmed {
			if r != '0' && r != '1' {
				// Любые другие символы — включая '(', 'a', пробелы в середине и т.д.
				return prefix
			}
		}

		return trimmed + " (bin)"
	})
}
