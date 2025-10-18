package autocorrection

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	noDigit   = regexp.MustCompile(`(?i)(['!.,:;a-z0-9@#$%\^&*\(\)\]\[\{\}_+~|-]+)\s*\(low\)`)
	haveDigit = regexp.MustCompile(`(?i)(['!.,:;a-z0-9@#$%\^&*\(\)\]\[\{\}_+~|-]+)\s*\(low(,\s*\d+)\)`)
)

func ToLower(text string) string {
	lowConverter := ""

	lowConverter = lowWithoutDigit(text)
	if haveDigit.MatchString(lowConverter) {
		lowConverter = lowWithDigit(lowConverter)
	}

	return lowConverter
}

func lowWithoutDigit(text string) string {
	lowConverter := ""

	lowConverter = noDigit.ReplaceAllStringFunc(text, func(s string) string {
		str := ""

		for i := 0; i < len(s); i++ {
			char := s[i]

			if i == len(s)-5 || char == ' ' {
				break
			}

			str += string(char)
		}
		str = strings.ToLower(str)

		return str
	})

	return lowConverter
}

func lowWithDigit(text string) string {
	numeric := regexp.MustCompile(`(?i)\(low(,\s*\d+)\)`)
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
				strs[j] = strings.ToLower(strs[j])

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

// var pattern = regexp.MustCompile(`(?i)(.*?)\s*\(low(?:,\s*(\d+))?\)`)

// func ToLower(text string) string {
// 	for {
// 		matches := pattern.FindAllStringSubmatch(text, -1)
// 		if len(matches) == 0 {
// 			break
// 		}

// 		// Берём последнее совпадение
// 		lastMatch := matches[len(matches)-1]
// 		fullMatch := lastMatch[0]
// 		prefix := lastMatch[1]    // (.*?) — всегда есть (может быть "")
// 		numberStr := lastMatch[2] // может быть "", если группа не участвовала

// 		n := 1
// 		if numberStr != "" {
// 			if num, err := strconv.Atoi(strings.TrimSpace(numberStr)); err == nil && num > 0 {
// 				n = num
// 			}
// 		}

// 		words := strings.Fields(prefix)
// 		toLower := n
// 		if toLower > len(words) {
// 			toLower = len(words)
// 		}

// 		for i := len(words) - toLower; i < len(words); i++ {
// 			words[i] = strings.ToLower(words[i])
// 		}

// 		newPrefix := strings.Join(words, " ")
// 		// Заменяем последнее совпадение
// 		pos := strings.LastIndex(text, fullMatch)
// 		if pos == -1 {
// 			break // на всякий случай
// 		}
// 		text = text[:pos] + newPrefix + text[pos+len(fullMatch):]
// 	}

// 	return text
// }
