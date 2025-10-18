package autocorrection

import (
	"regexp"
	"strings"
)

var (
	reAnA = regexp.MustCompile(`(?i)\b(a|an)\s+(\w+)`)

	exceptionAn = map[string]struct{}{
		"hour": {}, "honour": {}, "mba": {}, "sos": {}, "fbi": {}, "sms": {}, "honor": {},
	}

	exceptionA = map[string]struct{}{
		"one": {}, "user": {}, "university": {}, "european": {}, "unicorn": {},
	}
)

func FixArticles(text string) string {
	text = reAnA.ReplaceAllStringFunc(text, func(s string) string {
		slice := strings.Split(s, " ")
		caped := false

		if slice[0][0] == 'A' {
			caped = true
		}

		if _, ok := exceptionA[strings.ToLower(slice[1])]; ok {
			slice[0] = "a"

			if caped {
				slice[0] = strings.ToUpper(slice[0])
			}
		} else if _, ok := exceptionAn[strings.ToLower(slice[1])]; ok {
			slice[0] = "an"

			if caped {
				slice[0] = strings.ToUpper(string(slice[0][0])) + "n"
			}
		} else {
			if letter := strings.ToLower(string(slice[1][0]))[0]; letter == 'a' || letter == 'e' || letter == 'i' || letter == 'o' || letter == 'u' {
				slice[0] = "an"

				if caped {
					slice[0] = strings.ToUpper(string(slice[0][0])) + "n"
				}
			} else {
				slice[0] = "a"

				if caped {
					slice[0] = strings.ToUpper(slice[0])
				}
			}
		}

		s = slice[0] + " " + slice[1]

		return s
	})

	return text
}
