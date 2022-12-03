package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func SplitAny(s string, seps string) []string { // функция разделения строки по нескольким значениям
	splitter := func(r rune) bool {
		return strings.ContainsRune(seps, r)
	}
	return strings.FieldsFunc(s, splitter)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите выражение для вычисления арабскими или римскими цифрами в формате: 1 + 1 или I + VI, " +
		"используйте числа от 1 до 10")
	valueEnter, _ := reader.ReadString('\n')   // ждет ввода данных в формате строки
	valueEnter = strings.ToTitle(valueEnter)   //делает все буквы заглавными
	valueEnter = strings.TrimSpace(valueEnter) // очищает все пустоты (пробелы, табуляцию)
	numbers := SplitAny(valueEnter, "+-*/")    // разделяет строку по оператору

	err := errors.New("недопустимый ввод")
	if len(numbers) < 2 || len(numbers) > 2 {
		fmt.Println(err)
	}

	romanToArab := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9,
		"X": 10, "XI": 11, "XII": 12, "XIII": 13, "XIV": 14, "XV": 15, "XVI": 16, "XVII": 17, "XVIII": 18, "XIX": 19,
		"XX": 20, "XXI": 21, "XXII": 22, "XXIII": 23, "XXIV": 24, "XXV": 25, "XXVI": 26, "XXVII": 27, "XXVIII": 28,
		"XXIX": 29, "XXX": 30, "XXXI": 31, "XXXII": 32, "XXXIII": 33, "XXXIV": 34, "XXXV": 35, "XXXVI": 36, "XXXVII": 37,
		"XXXVIII": 38, "XXXIX": 39, "XL": 40, "XLI": 41, "XLII": 42, "XLIII": 43, "XLIV": 44, "XLV": 45, "XLVI": 46,
		"XLVII": 47, "XLVIII": 48, "XLIX": 49, "L": 50, "LI": 51, "LII": 52, "LIII": 53, "LIV": 54, "LV": 55, "LVI": 56,
		"LVII": 57, "LVIII": 58, "LIX": 59, "LX": 60, "LXI": 61, "LXII": 62, "LXIII": 63, "LXIV": 64, "LXV": 65,
		"LXVI": 66, "LXVII": 67, "LXVIII": 68, "LXIX": 69, "LXX": 70, "LXXI": 71, "LXXII": 72, "LXXIII": 73, "LXXIV": 74,
		"LXXV": 75, "LXXVI": 76, "LXXVII": 77, "LXXVIII": 78, "LXXIX": 79, "LXXX": 80, "LXXXI": 81, "LXXXII": 82,
		"LXXXIII": 83, "LXXXIV": 84, "LXXXV": 85, "LXXXVI": 86, "LXXXVII": 87, "LXXXVIII": 88, "LXXXIX": 89, "XC": 90,
		"XCI": 91, "XCII": 92, "XCIII": 93, "XCIV": 94, "XCV": 95, "XCVI": 96, "XCVII": 97, "XCVIII": 98, "XCIX": 99, "C": 100}

	arabToRoman := map[int]string{1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX",
		10: "X", 11: "XI", 12: "XII", 13: "XIII", 14: "XIV", 15: "XV", 16: "XVI", 17: "XVII", 18: "XVIII", 19: "XIX",
		20: "XX", 21: "XXI", 22: "XXII", 23: "XXIII", 24: "XXIV", 25: "XXV", 26: "XXVI", 27: "XXVII", 28: "XXVIII",
		29: "XXIX", 30: "XXX", 31: "XXXI", 32: "XXXII", 33: "XXXIII", 34: "XXXIV", 35: "XXXV", 36: "XXXVI", 37: "XXXVII",
		38: "XXXVIII", 39: "XXXIX", 40: "XL", 41: "XLI", 42: "XLII", 43: "XLIII", 44: "XLIV", 45: "XLV", 46: "XLVI",
		47: "XLVII", 48: "XLVIII", 49: "XLIX", 50: "L", 51: "LI", 52: "LII", 53: "LIII", 54: "LIV", 55: "LV", 56: "LVI",
		57: "LVII", 58: "LVIII", 59: "LIX", 60: "LX", 61: "LXI", 62: "LXII", 63: "LXIII", 64: "LXIV", 65: "LXV",
		66: "LXVI", 67: "LXVII", 68: "LXVIII", 69: "LXIX", 70: "LXX", 71: "LXXI", 72: "LXXII", 73: "LXXIII", 74: "LXXIV",
		75: "LXXV", 76: "LXXVI", 77: "LXXVII", 78: "LXXVIII", 79: "LXXIX", 80: "LXXX", 81: "LXXXI", 82: "LXXXII",
		83: "LXXXIII", 84: "LXXXIV", 85: "LXXXV", 86: "LXXXVI", 87: "LXXXVII", 88: "LXXXVIII", 89: "LXXXIX", 90: "XC",
		91: "XCI", 92: "XCII", 93: "XCIII", 94: "XCIV", 95: "XCV", 96: "XCVI", 97: "XCVII", 98: "XCVIII", 99: "XCIX", 100: "C"}

	var re = regexp.MustCompile(`0|1|2|3|4|5|6|7|8|9|10`) // проверяет наличие в строке арабских чисел
	var reRoman = regexp.MustCompile(`[IXCLVixclv]`)      // проверяет наличие в строке римских чисел
	var numberOne, numberTwo int
	var result int

	if re.MatchString(numbers[0]) == true && re.MatchString(numbers[1]) == true { // если в строке оба числа арабские
		numberOne, _ = strconv.Atoi(numbers[0]) // конвертирует строку в число
		numberTwo, _ = strconv.Atoi(numbers[1])
		if numberOne > 10 || numberTwo > 10 {
			fmt.Println(err)
		} else if strings.Contains(valueEnter, "+") == true { // проверяет, содержит ли строка символ "+"
			result = numberOne + numberTwo
			fmt.Println(result)
		} else if strings.Contains(valueEnter, "-") == true {
			result = numberOne - numberTwo
			fmt.Println(result)
		} else if strings.Contains(valueEnter, "*") == true {
			result = numberOne * numberTwo
			fmt.Println(result)
		} else if strings.Contains(valueEnter, "/") == true {
			result = numberOne / numberTwo
			fmt.Println(result)
		}
	} else if re.MatchString(numbers[0]) == true && re.MatchString(numbers[1]) == false {
		fmt.Println(err)
	} else if re.MatchString(numbers[0]) == false && re.MatchString(numbers[1]) == true {
		fmt.Println(err)
	} else if re.MatchString(numbers[0]) == false && re.MatchString(numbers[1]) == false {
		if reRoman.MatchString(numbers[0]) == true && reRoman.MatchString(numbers[1]) == true {
			numberOne, _ = romanToArab[numbers[0]]
			numberTwo, _ = romanToArab[numbers[1]]
			if numberOne > 10 || numberTwo > 10 {
				fmt.Println(err)
			} else if strings.Contains(valueEnter, "+") == true { // проверяет, содержит ли строка символ "+"
				result = numberOne + numberTwo
			} else if strings.Contains(valueEnter, "-") == true {
				result = numberOne - numberTwo
				if result <= 0 {
					fmt.Println(err)
				}
			} else if strings.Contains(valueEnter, "*") == true {
				result = numberOne * numberTwo
			} else if strings.Contains(valueEnter, "/") == true {
				result = numberOne / numberTwo
			}
			strResult := arabToRoman[result]
			fmt.Println(strResult)
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}
