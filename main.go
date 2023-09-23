package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	Arabic = 0
	Rome   = 1
)

var ArabicAlp = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
var RomeAlp = []string{"I", "V", "X"}

func contains(str string, alp []string) bool {
	for _, el := range str {
		for _, el2 := range alp {
			if string(el) == el2 {
				return true
			}
		}
	}
	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var NumType int

	fmt.Println("DnKrsk GoLang Calculator for Kata.")
	input, _ := reader.ReadString('\n')
	//symbs := strings.Split(input, " ")
	if contains(input, ArabicAlp) {
		NumType = Arabic
		if contains(input, RomeAlp) {
			fmt.Println("братик ты куда")
			os.Exit(0)
		}
	} else {
		NumType = Rome
	}
	fmt.Println(NumType)
}
