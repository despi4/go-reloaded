package goreloaded

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	withoutDigit = regexp.MustCompile(`(?i)(['!.,:;a-z0-9@#$%\^&*\(\)\]\[\{\}_+~|-]+)\s*\(up\)`)
	withDigit    = regexp.MustCompile(`(?i)(['!.,:;a-z0-9@#$%\^&*\(\)\]\[\{\}_+~|-]+)\s*\(up(, \d+)\)`)
)

func AlphaUp(text string) string {
	upConverter := ""

	if withoutDigit.MatchString(text) {
		upConverter = upWithoutDigit(text)
	}
	if withDigit.MatchString(upConverter) {
		upConverter = upWithDigit(upConverter)
	}

	return upConverter
}

func upWithoutDigit(text string) string {
	upConverter := ""

	upConverter = withoutDigit.ReplaceAllStringFunc(text, func(s string) string {
		str := ""

		for i := 0; i < len(s); i++ {
			char := s[i]

			if i == len(s)-4 {
				break
			}

			str += string(char)
		}
		str = strings.ToUpper(str)

		return str
	})

	return upConverter
}

func upWithDigit(text string) string {
	numeric := regexp.MustCompile(`(?i)\(up(, \d+)\)`)
	miniRegular := regexp.MustCompile(`(?i)\((up,)\s+|(\d+)\)`)

	findDigits := numeric.FindAllString(text, -1)
	digitArr := []int{}

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

	index := miniRegular.FindAllStringIndex(text, -1)

	noFurther := ""

	if index[1][1] < len(text) {
		noFurther = text[index[1][1]:]
		noFurther = strings.TrimSpace(noFurther)
	}

	replaceUp := numeric.ReplaceAllStringFunc(text, func(s string) string {
		return " "
	})

	strs := strings.Split(replaceUp, " ")

	for i := range strs {
		if strs[i] == noFurther && noFurther != "" {
			strs = strs[:len(strs)-1]
		}
	}

	for i := 0; i < len(digitArr); i++ {
		for j := len(strs) - 1; digitArr[i] != 0; j-- {
			if strs[j] != "" {
				strs[j] = strings.ToUpper(strs[j])
				digitArr[i]--
			}
			if j == 0 {
				break
			}
		}
	}

	str := ""

	for _, s := range strs {
		str += s + " "
	}

	str += noFurther

	return str
}
