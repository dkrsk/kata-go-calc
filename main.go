package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ArabicAlp = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
var RomeAlp = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
var Alphabet = map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100}

type Operator int16

const (
	Plus     Operator = 0
	Minus    Operator = 1
	Multiply Operator = 2
	Divide   Operator = 3
)

const (
	BadRomeResult = "Вывод ошибки, так как в римской системе нет отрицательных чисел."
	TwoTypesUsage = "Вывод ошибки, так как используются одновременно разные системы счисления."
	NotAEquation  = "Вывод ошибки, так как строка не является математической операцией."
	BadFormat     = "Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор"
	BadRange      = "Вывод ошибки, так как числа не в диапазоне от 1 до 10"
)

type NumType int16

const (
	Arabic NumType = 0
	Rome   NumType = 1
)

type Input struct {
	left       int
	right      int
	Operation  Operator
	NumberType NumType
}

func main() {
	fmt.Println("DnKrsk GoLang Calculator for Kata.")

	var input = GetInput()
	fmt.Println(Calc(input))
}

func GetInput() Input {
	reader := bufio.NewReader(os.Stdin)
	input := Input{}

	str, _ := reader.ReadString('\n')
	str = strings.ToUpper(str[:len(str)-2])
	//str = strings.ReplaceAll(str, " ", "")
	tmp := strings.Split(str, " ")
	if len(tmp) != 3 {
		panic(BadFormat)
	}

	if CheckNumType(str) == Rome {
		input.left = RomeToArb(tmp[0])
		input.right = RomeToArb(tmp[2])
		input.NumberType = Rome
	} else {
		input.left, _ = strconv.Atoi(tmp[0])
		input.right, _ = strconv.Atoi(tmp[2])
		input.NumberType = Arabic
	}

	switch tmp[1] {
	case "+":
		input.Operation = Plus
		break
	case "-":
		input.Operation = Minus
		break
	case "*":
		input.Operation = Multiply
		break
	case "/":
		input.Operation = Divide
		break
	default:
		panic(NotAEquation)
	}

	return input
}

func Calc(input Input) string {
	var res int

	if input.left > 10 || input.right > 10 {
		panic(BadRange)
	}

	switch input.Operation {
	case Plus:
		res = input.left + input.right
		break
	case Minus:
		res = input.left - input.right
		break
	case Divide:
		res = input.left / input.right
		break
	case Multiply:
		res = input.left * input.right
		break
	}

	if input.NumberType == Arabic {
		return strconv.Itoa(res)
	} else {
		if res <= 0 {
			panic(BadRomeResult)
		}
		return ArbToRome(res)
	}
}

func CheckNumType(str string) NumType {
	if Contains(str, ArabicAlp) {
		if Contains(str, RomeAlp) {
			panic(TwoTypesUsage)
		}
		return Arabic
	} else {
		return Rome
	}
}

func Contains(str string, alp []string) bool {
	for _, el := range str {
		for _, el2 := range alp {
			if string(el) == el2 {
				return true
			}
		}
	}
	return false
}

func RomeToArb(rome string) int {
	var result int

	for q := 0; q < len(rome)-1; q++ {
		if Alphabet[string(rome[q])] < Alphabet[string(rome[q+1])] {
			result -= Alphabet[string(rome[q])]
		} else {
			result += Alphabet[string(rome[q])]
		}
	}
	result += Alphabet[string(rome[len(rome)-1])]
	if result == 0 {
		panic(NotAEquation)
	}
	return result
}
func ArbToRome(arb int) string {
	var result string
	d := arb / 100
	result += strings.Repeat("C", d)
	arb %= 100
	d = arb / 10
	result += strings.Repeat("X", d)
	arb %= 10
	result += RomeAlp[arb]
	return result
}
