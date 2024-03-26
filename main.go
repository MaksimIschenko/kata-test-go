package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Разделение строки по пробелу
func splitInput(text *string) []string {
	var splittedString []string
	splittedString = strings.Split(*text, " ")
	return splittedString
}

// Проверка валидности введённой операции
func validateOperation(operation string) bool {
	var validOperation [4]string = [4]string{"+", "-", "*", "/"}
	for i := range validOperation {
		if validOperation[i] == operation {
			return true
		}
	}
	return false
}

// Проверка валидности ввода римского числа
func validRoman(romanOperand string) (int, bool) {
	romanToArabic := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}
	if val, ok := romanToArabic[romanOperand]; ok {
		return val, ok
	}
	return -1, false
}

// Выполнение математической операции
func calculate(numFirst int, numSecond int, operation string) int {
	switch operation {
	case "+":
		return numFirst + numSecond
	case "-":
		return numFirst - numSecond
	case "*":
		return numFirst * numSecond
	default:
		return numFirst / numSecond
	}
}

// Перевод результата вычисления в римские числа
func resArabicToRoman(result int) string {

	var romanString string

	arabic := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	for result != 0 {
		for i, value := range arabic {
			if result >= arabic[i] {
				result -= value
				romanString += symbols[i]
				break
			}
		}
	}
	return romanString
}

func main() {

	// Создание объекта ввода из терминала
	reader := bufio.NewReader((os.Stdin))

	for {
		fmt.Println("Input:")
		text, _ := reader.ReadString('\n')

		// Разделение строки
		sInput := splitInput(&text)

		// Ошибка ввода
		if len(sInput) != 3 {
			panic("Формат математической операции не удовлетворяет заданию — два операнда и один оператор")
		}

		// Определение переменных
		firstOperand := sInput[0]
		secondOperand := strings.TrimSpace(sInput[2])
		operation := sInput[1]

		// Проверка валидности введённой операции
		isValid := validateOperation(operation)
		if isValid == false {
			panic("Неизвестная операция. Доступные операции: + - / *")
		}

		// Преобразование строки в число
		numFirst, errFirst := strconv.Atoi(firstOperand)
		numSecond, errSecond := strconv.Atoi(secondOperand)

		var romanFlag bool = false

		// Обработка ошибки преобразования
		if (errFirst != nil && errSecond != nil) == true {
			romanFlag = true
		} else if (errFirst == nil && errSecond == nil) == true {
			romanFlag = false
		} else {
			panic("Допустимо использовать только арабские или только римские числа")
		}

		// Преобразрвание строки (римского числа) в арабское для дальнейшего вычисления
		if romanFlag == true {
			numFirst, _ = validRoman(firstOperand)
			numSecond, _ = validRoman(secondOperand)
			if (numFirst == -1 || numSecond == -1) == true {
				panic("Неверный операнд римских чисел")
			}
		}

		// Использование недопустимого числа
		if ((numFirst >= 1 && numFirst <= 10) && (numSecond >= 1 && numSecond <= 10)) == false {
			panic("Неверный операнд. Ввод от 1 до 10.")
		}

		// Вычисление мат. операции
		res := calculate(numFirst, numSecond, operation)

		// Вывод результата
		fmt.Println("Output: ")

		if romanFlag == true {
			if res <= 0 {
				panic("В римской системе исчисления нет отрицательных чисел и нуля.")
			}
			romanString := resArabicToRoman(res)
			fmt.Println(romanString)
		} else {
			fmt.Println(res)
		}
		fmt.Printf("\n")
	}
}
