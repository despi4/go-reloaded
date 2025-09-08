package goreloaded

import (
	"math"
	"regexp"
	"strconv"
)

var binPattern = `(?i)([0-9a-f]+)\s*\(bin(?:,\d+)?\)`

func BinToDec(text string) string {
	regular := regexp.MustCompile(binPattern)

	powInt := func(x, y int) int { return int(math.Pow(float64(x), float64(y))) }

	binConverted := regular.ReplaceAllStringFunc(text, func(s string) string {
		length := -1
		decimal := 0

		for _, char := range s {
			if char == '0' || char == '1' {
				length++
				continue
			}
			break
		}

		for _, char := range s {
			if char == '0' {
				decimal += 0 * powInt(2, length)
				length--
				continue
			} else if char == '1' {
				decimal += 1 * powInt(2, length)
				length--
				continue
			}
			break
		}

		return strconv.Itoa(decimal)
	})

	return binConverted
}
