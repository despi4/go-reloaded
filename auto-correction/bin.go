package autocorrection

import (
	"regexp"
	"strconv"
)

// var binPattern = `(?i)([0-9a-f]+)\s*\(bin(?:,\s*\d+)?\)`

// func BinToDec(text string) string {
// 	regular := regexp.MustCompile(binPattern)

// 	powInt := func(x, y int) int { return int(math.Pow(float64(x), float64(y))) }

// 	binConverter := regular.ReplaceAllStringFunc(text, func(s string) string {
// 		length := -1
// 		decimal := 0
// 		notBin := []rune{}

// 		for _, char := range s {
// 			if char == '0' || char == '1' {
// 				length++
// 				continue
// 			}
// 			break
// 		}

// 		for _, char := range s {
// 			if char == '0' {
// 				decimal += 0 * powInt(2, length)
// 				length--
// 				continue
// 			} else if char == '1' {
// 				decimal += 1 * powInt(2, length)
// 				length--
// 				continue
// 			} else if char >= '2' && char <= '9' {
// 				notBin = append(notBin, char)
// 			}
// 			break
// 		}

// 		if len(notBin) != 0 {
// 			return string(notBin)
// 		}

// 		return strconv.Itoa(decimal)
// 	})

// 	return binConverter
// }

// Ищем: двоичное число (только 0 и 1), за которым следует "(bin" (с необязательным ,число)
var binPattern = regexp.MustCompile(`(?i)\b([01]+)\s*\(bin(?:,\s*\d+)?\)`)

func BinToDec(text string) string {
	return binPattern.ReplaceAllStringFunc(text, func(match string) string {
		// Извлекаем первую группу — двоичную строку
		// К сожалению, ReplaceAllStringFunc не даёт доступа к группам напрямую,
		// поэтому делаем Submatch вручную
		submatches := binPattern.FindStringSubmatch(match)
		if len(submatches) < 2 {
			return match // на всякий случай
		}
		binStr := submatches[1]

		// Преобразуем двоичную строку в число
		num, err := strconv.ParseInt(binStr, 2, 64)
		if err != nil {
			return match // если ошибка — оставляем как есть
		}

		return strconv.FormatInt(num, 10)
	})
}
