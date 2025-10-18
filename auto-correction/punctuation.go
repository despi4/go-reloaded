package autocorrection

import (
	"regexp"
	"strings"
)

var (
	manyDots                = regexp.MustCompile(`(\s*\.\s*\.\s*\.)`)
	ordinaryQuote           = regexp.MustCompile(`'\s*(.*?)\s*'`)
	doubleQuote             = regexp.MustCompile(`"\s*(.*?)\s*"`)
	questionExclamationMark = regexp.MustCompile(`\s*!\s*\?`)
	punctacia               = regexp.MustCompile(`\s*[,.?!;:]+`)
	contractionRe           = regexp.MustCompile(`\b([A-Za-z]+)\s*'\s*([A-Za-z]+)\b`)
)

func Punctuation(text string) string {
	symbols := map[string]struct{}{
		"m":  {},
		"s":  {},
		"ve": {},
		"t":  {},
		"re": {},
		"ll": {},
		"d":  {},
	}

	text = contractionRe.ReplaceAllStringFunc(text, func(s string) string {
		index := -1
		runes := []rune{}

		for i, char := range s {
			if char == []rune("'")[0] {
				index = i
				break
			}
		}

		if index != -1 {
			_, ok := symbols[strings.TrimSpace(s[index+1:])]
			if ok {
				for _, char := range s {
					if char != ' ' {
						if char == []rune(string("'"))[0] {
							runes = append(runes, '`')
							continue
						}

						runes = append(runes, char)
					}
				}
			}
		}

		if len(runes) != 0 {
			return string(runes)
		}

		return s
	})

	text = manyDots.ReplaceAllStringFunc(text, func(s string) string {
		return "... "
	})

	text = punctacia.ReplaceAllStringFunc(text, func(s string) string {
		s = strings.TrimSpace(s)

		return s + " "
	})

	text = ordinaryQuote.ReplaceAllStringFunc(text, func(s string) string {
		s = strings.TrimSpace(s[1 : len(s)-1])
		s = " '" + s + "' "
		return s
	})

	text = doubleQuote.ReplaceAllStringFunc(text, func(s string) string {
		if len(s) <= 2 {
			return s
		}

		inner := strings.TrimSpace(s[1 : len(s)-1])
		if len(inner) == 0 {
			return s
		}

		inner = string(inner[0]) + inner[1:]
		return ` "` + inner + `" `
	})

	text = questionExclamationMark.ReplaceAllStringFunc(text, func(s string) string {
		return "!?"
	})

	setOfPunctuation := regexp.MustCompile(`([.,;:!?\[\]]\s*){2,}`)

	text = setOfPunctuation.ReplaceAllStringFunc(text, func(s string) string {
		runes := []rune{}

		for _, char := range s {
			if char != ' ' {
				runes = append(runes, char)
			}
		}
		return string(runes) + " "
	})

	// text = toCap(text)

	text = FixSpace(text)

	text = strings.TrimSpace(text)

	return text
}

// func toCap(text string) string {
// 	copyText := []rune{}
// 	cap := false
// 	str := []rune{}

// 	for _, char := range text {
// 		if char == '.' {
// 			str = append(str, char)
// 		} else if char != ' ' {
// 			str = []rune{}
// 		}

// 		if char == '.' || char == '?' || char == '!' {
// 			if len(str) >= 3 {
// 				cap = false
// 				copyText = append(copyText, char)
// 				continue
// 			}
// 			cap = true
// 		}

// 		if cap && char != ' ' && char != '.' && char != '?' && char != '!' {
// 			char = []rune(strings.ToUpper(string(char)))[0]
// 			cap = false
// 		}

// 		copyText = append(copyText, char)
// 	}

// 	return string(copyText)
// }
