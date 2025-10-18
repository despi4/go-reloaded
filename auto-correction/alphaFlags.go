package autocorrection

import (
	"regexp"
	"strings"
)

var (
	AlphaSymbolsInFlag = regexp.MustCompile(`(?i)\s*\([^)]*(up|low|cap)[^)]*\)+`)
	AlphaSpaceInFlag   = regexp.MustCompile(`(?i)\(\s*[^()]*\s*(up|low|cap)\s*,?\s*[^()]*\s*\)`)
	symbolsZa          = regexp.MustCompile(`(?i)^\s*\((up|low|cap)\s*(?:,\s*\d*)?\)`)
)

func AlphaNavigator(text string) string {
	text = AlphaTrimExtraSpace(text)
	text = AlphaDelExtraSymbols(text)
	text = delAlphaFlags(text)

	return text
}

func AlphaTrimExtraSpace(text string) string {
	trimSpace := AlphaSpaceInFlag.ReplaceAllStringFunc(text, func(s string) string {
		s = strings.TrimSpace(s[1 : len(s)-1])

		return "(" + s + ")"
	})

	return trimSpace
}

// func AlphaDelExtraSymbols(text string) string {
// 	var (
// 		digits = regexp.MustCompile(`(?i)\((up|low|cap)\s*,\s*-?\d+\s*\)`)
// 		// upFlag  = regexp.MustCompile(`(?i)\s*\(up\s*[\w\W]*\)`)
// 		// lowFlag = regexp.MustCompile(`(?i)\s*\(low\s*[\w\W]*\)`)
// 		// capFlag = regexp.MustCompile(`(?i)\s*\(cap\s*[\w\W]*\)`)
// 		alphaFlags = regexp.MustCompile(`(?i)\s*\((cap|low|up)\s*,?\s*\d*\)`)
// 		number     = ""
// 	)

// 	// fmt.Printf("withDigits : %v\nflags : %v\n", digits.FindAllString(text, -1), alphaFlags.FindAllString(text, -1))

// 	delSymbols := AlphaSymbolsInFlag.ReplaceAllStringFunc(text, func(s string) string {
// 		minus := false

// 		if digits.MatchString(s) {
// 			str := ""

// 			for _, char := range s {
// 				if char == '-' {
// 					minus = true
// 				}
// 				if char >= '0' && char <= '9' {
// 					str += string(char)
// 				}
// 			}

// 			number += str
// 		}

// 		if alphaFlags.MatchString(s) {
// 			if len(number) != 0 {
// 				s = " (up, " + number + ")"
// 				number = ""
// 			} else {
// 				s = " (up) "
// 			}
// 		} else if alphaFlags.MatchString(s) {
// 			if len(number) != 0 {
// 				s = " (low, " + number + ")"
// 				number = ""
// 			} else {
// 				s = " (low) "
// 			}
// 		} else if alphaFlags.MatchString(s) {
// 			if len(number) != 0 {
// 				s = " (cap, " + number + ")"
// 				number = ""
// 			} else {
// 				s = " (cap) "
// 			}
// 		}

// 		if minus {
// 			s = ""
// 		}

// 		return s
// 	})

// 	return delSymbols
// }

func AlphaDelExtraSymbols(text string) string {
	// Убираем всё лишнее, оставляем только (cmd) или (cmd, N)
	cleanFlag := regexp.MustCompile(`(?i)\((up|low|cap)(?:\s*,\s*(-?\d+))?\)`)

	result := cleanFlag.ReplaceAllStringFunc(text, func(s string) string {
		// Найдём команду и, возможно, число
		matches := cleanFlag.FindStringSubmatch(s)
		if len(matches) < 2 {
			return s // не должно случиться
		}

		cmd := strings.ToLower(matches[1])
		num := matches[2] // может быть пустой

		if strings.Contains(num, "-") {
			return "" // удаляем, если есть минус
		}

		if num != "" {
			return " (" + cmd + ", " + num + ")"
		}
		return " (" + cmd + ") "
	})

	return result
}

func delAlphaFlags(text string) string {
	return symbolsZa.ReplaceAllString(text, "")
}

// func delAlphaFlags(text string) string {
// 	text = symbolsZa.ReplaceAllStringFunc(text, func(s string) string {
// 		index := symbolsZa.FindStringIndex(text)

// 		fmt.Println(s)

// 		if len(strings.TrimSpace(text[:index[0]])) == 0 {
// 			s = ""
// 		}

// 		return s
// 	})

// 	return text
// }
