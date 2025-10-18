package autocorrection

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	falseDigit = regexp.MustCompile(`(?i)\b(\w+\W+)\s*\(cap\)`)
	trueDigit  = regexp.MustCompile(`(?i)(['!.,:;a-z0-9@#$%\^&*\(\)\]\[\{\}_+~|-]+)\s*\(cap(,\s*\d+)\)`)
)

func FirstBig(text string) string {
	capConverter := ""

	capConverter = capWithoutDigit(text)
	if trueDigit.MatchString(capConverter) {
		capConverter = capWithDigit(capConverter)
	}

	return capConverter
}

func capWithoutDigit(text string) string {
	capConverter := ""

	capConverter = falseDigit.ReplaceAllStringFunc(text, func(s string) string {
		str := ""
		caped := false

		s = s[:len(s)-5]

		for i := 0; i < len(s); i++ {
			char := s[i]
			char = strings.ToLower(string(char))[0]

			if char >= 'a' && char <= 'z' && !caped {
				char = []byte(strings.ToUpper(string(char)))[0]
				caped = true
			} else {
				char = []byte(strings.ToLower(string(char)))[0]
			}

			// if i == len(s)-5 || char == ' ' {
			// 	break
			// }

			str += string(char)
		}

		return str
	})

	return capConverter
}

func capWithDigit(text string) string {
	numeric := regexp.MustCompile(`(?i)\s*\(cap(,\s*\d+)\)`)
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
				runes := []rune(strs[j])
				caped := false

				if string(runes) == "(" || string(runes) == ")" {
					continue
				}

				for k := 0; k < len(runes); k++ {
					char := runes[k]

					char = []rune(strings.ToLower(string(char)))[0]

					if char >= 'a' && char <= 'z' && !caped {
						char = []rune(strings.ToUpper(string(runes[k])))[0]
						caped = true
					} else {
						char = []rune(strings.ToLower(string(runes[k])))[0]
					}

					runes[k] = char
				}
				caped = false

				strs[j] = string(runes)

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
