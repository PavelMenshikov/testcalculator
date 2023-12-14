package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanMap = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
}

func romanToInt(roman string) int {
	result := 0
	prevValue := 0

	for i := len(roman) - 1; i >= 0; i-- {
		currentValue := romanMap[rune(roman[i])]

		if currentValue < prevValue {
			result -= currentValue
		} else {
			result += currentValue
		}

		prevValue = currentValue
	}

	return result
}

func intToRoman(number int) string {
	if number <= 0 || number > 3999 {
		return "Недопустимое число"
	}

	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romanDigits := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	romanValue := ""
	for i := 0; i < len(values); i++ {
		for number >= values[i] {
			number -= values[i]
			romanValue += romanDigits[i]
		}
	}

	return romanValue
}

func calculate(a, b string, operator string) string {
	if strings.Contains(a, "I") || strings.Contains(b, "I") {
		arabicA := romanToInt(a)
		arabicB := romanToInt(b)

		result := 0
		switch operator {
		case "+":
			result = arabicA + arabicB
		case "-":
			result = arabicA - arabicB
		case "*":
			result = arabicA * arabicB
		case "/":
			if arabicB != 0 {
				result = arabicA / arabicB
			}
		}

		return intToRoman(result)
	}

	intA, errA := strconv.Atoi(a)
	intB, errB := strconv.Atoi(b)

	if errA != nil || errB != nil {
		return "Ошибка: Некорректные числа"
	}

	result := 0
	switch operator {
	case "+":
		result = intA + intB
	case "-":
		result = intA - intB
	case "*":
		result = intA * intB
	case "/":
		if intB != 0 {
			result = intA / intB
		}
	}

	return strconv.Itoa(result)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Введите выражение для вычисления (например, I + II или 1 + 2):")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		parts := strings.Split(input, " ")
		if len(parts) != 3 {
			fmt.Println("Ошибка: Некорректный формат математической операции.")
			continue
		}

		a := parts[0]
		b := parts[2]
		operator := parts[1]

		if operator != "+" && operator != "-" && operator != "*" && operator != "/" {
			fmt.Println("Ошибка: Недопустимая математическая операция.")
			continue
		}

		result := calculate(a, b, operator)
		fmt.Println("Результат:", result)
	}
}
