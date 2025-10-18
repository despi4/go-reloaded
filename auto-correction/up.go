package autocorrection

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	withoutDigit = regexp.MustCompile(`(?i)(['!.,:;a-z0-9@#$%\^&*\(\)\]\[\{\}_+~|-]+)\s*\(up\)`)
	withDigit    = regexp.MustCompile(`(?i)(['!.,:;a-z0-9@#$%\^&*\(\)\]\[\{\}_+~|-]+)\s*\(up(,\s*\d+)\)`)
)

func ToUpper(text string) string {
	upConverter := ""

	upConverter = upWithoutDigit(text)
	if withDigit.MatchString(upConverter) {
		upConverter = upWithDigit(upConverter)
	}

	return upConverter
}

func upWithoutDigit(text string) string {
	upConverter := ""

	upConverter = withoutDigit.ReplaceAllStringFunc(text, func(s string) string {
		str := ""

		s = s[:len(s)-4]

		for i := 0; i < len(s); i++ {
			char := s[i]

			// if i == len(s)-4 || char == ' ' {
			// 	break
			// }

			str += string(char)
		}
		str = strings.ToUpper(str)

		return str
	})

	return upConverter
}

func upWithDigit(text string) string {
	numeric := regexp.MustCompile(`(?i)\(up(,\s*\d+)\)`)
	findDigits := numeric.FindAllString(text, -1)
	resText := text
	digitArr := []int{}
	specialSymbols := map[rune]struct{}{
		'.': {},
		',': {},
		':': {},
		';': {},
		'-': {},
		'_': {},
		'+': {},
		'=': {},
		'|': {},
		'/': {},
		'!': {},
		'?': {},
		' ': {},
	}

	// take number of flag realization
	for _, str := range findDigits {
		number := ""
		for _, char := range str {
			if char >= '0' && char <= '9' {
				number += string(char)
			}
		}
		d, _ := strconv.Atoi(number)

		digitArr = append(digitArr, d)
	}

	for i := 0; i < len(digitArr); i++ {
		index := numeric.FindStringIndex(resText)
		space := false

		if index[1] != len(resText) {
			_, ok := specialSymbols[rune(resText[index[1]])]
			if !ok {
				space = true
			}
		}

		strs := strings.Split(resText[:index[0]], " ")

		for j := len(strs) - 1; digitArr[i] != 0; j-- {
			if strs[j] != "" {
				strs[j] = strings.ToUpper(strs[j])

				if j == 0 {
					break
				}

				digitArr[i]--
			}
		}

		instStr := ""

		for i, str := range strs {
			instStr += str

			if i != len(strs)-1 {
				instStr += " "
			}
		}

		if space {
			resText = instStr + " " + resText[index[1]:]
			continue
		}

		resText = instStr + resText[index[1]:]
	}

	return resText
}
