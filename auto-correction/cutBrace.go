package autocorrection

import (
	"regexp"
)

var (
	bracePattern = regexp.MustCompile(`\(\s*[^()]*\s*\)`)
	justBraces   = regexp.MustCompile(`\(\s*\)`)
)

func CutSpace(text string) string {
	text = justBraces.ReplaceAllString(text, "")

	text = bracePattern.ReplaceAllStringFunc(text, func(s string) string {
		runes := []rune{}

		for i := 1; i < len(s); i++ {
			char := s[i]

			if char != ' ' {
				runes = append(runes, []rune(s[i:])...)
				break
			}
		}

		for i := len(runes) - 2; i >= 0; i-- {
			char := runes[i]

			if char != ' ' {
				runes = runes[:i+1]
				break
			}
		}

		runes = append([]rune{'('}, runes...)
		runes = append(runes, ')')

		return string(runes)
	})

	return text
}
