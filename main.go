package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var ArabicAlp = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
var RomeAlp = []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}

type Operator int16

const (
	Plus     Operator = 0
	Minus    Operator = 1
	Multiply Operator = 2
	Divide   Operator = 3
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

	fmt.Println(input)
}

func GetInput() Input {
	reader := bufio.NewReader(os.Stdin)
	input := Input{}

	str, _ := reader.ReadString('\n')
	tmp := strings.Split(str[:len(str)-2], " ")

	if CheckNumType(str) == Rome {
		input.left = slices.Index(RomeAlp, tmp[0]) + 1
		input.right = slices.Index(RomeAlp, tmp[2]) + 1
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
	}

	return input
}

func CheckNumType(str string) NumType {
	if Contains(str, ArabicAlp) {
		if Contains(str, RomeAlp) {
			fmt.Println("братик ты куда")
			os.Exit(0)
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
