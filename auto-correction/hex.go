package autocorrection

import (
	"regexp"
	"strconv"
)

// var hexPattern = `(?i)([0-9a-f]+)\s*\(hex(?:,\s*\d+)?\)`

// func HexToDec(text string) string {
// 	regular := regexp.MustCompile(hexPattern)
// 	powInt := func(x, y int) int { return int(math.Pow(float64(x), float64(y))) }
// 	hexNumbers := map[rune]int{
// 		'A': 10,
// 		'B': 11,
// 		'C': 12,
// 		'D': 13,
// 		'E': 14,
// 		'F': 15,
// 	}

// 	hexConverter := regular.ReplaceAllStringFunc(text, func(s string) string {
// 		decimal := 0
// 		length := -1

// 		for _, char := range s {
// 			if char >= '0' && char <= '9' || char >= 'A' && char <= 'F' || char >= 'a' && char <= 'z' {
// 				length++
// 			} else {
// 				break
// 			}
// 		}

// 		for _, char := range s {
// 			if length < 0 {
// 				break
// 			}

// 			if char >= 'a' && char <= 'z' {
// 				char -= 32
// 			}

// 			value, ok := hexNumbers[char]
// 			if ok {
// 				decimal += value * powInt(16, length)
// 				length--
// 				continue
// 			}

// 			number, err := strconv.Atoi(string(char))
// 			if err != nil {
// 				log.Fatal(err)
// 			}

// 			decimal += number * powInt(16, length)
// 			length--
// 		}

// 		return strconv.Itoa(decimal)
// 	})

// 	return hexConverter
// }

var hexPattern = regexp.MustCompile(`(?i)\b([0-9a-f]+)\s*\(hex(?:,\s*\d+)?\)`)

func HexToDec(text string) string {
	return hexPattern.ReplaceAllStringFunc(text, func(match string) string {
		// Извлекаем hex-часть (первая группа)
		submatches := hexPattern.FindStringSubmatch(match)
		if len(submatches) < 2 {
			return match // fallback
		}
		hexStr := submatches[1]

		// Парсим как шестнадцатеричное число
		num, err := strconv.ParseInt(hexStr, 16, 64)
		if err != nil {
			// Если ошибка — оставляем оригинал (не падаем!)
			return match
		}

		return strconv.FormatInt(num, 10)
	})
}
