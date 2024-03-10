package main

import (
	"fmt"
	"strconv"
)

func input() { //Ввод чисел
	var numberOneInput, operator, numberTwoInput string
	fmt.Scan(&numberOneInput, &operator, &numberTwoInput)         //сканируем
	if checkIfInt(numberOneInput) && checkIfInt(numberTwoInput) { //если арабские числа
		var numberOne, numberTwo int
		numberOne, _ = strconv.Atoi(numberOneInput)
		numberTwo, _ = strconv.Atoi(numberTwoInput)
		if numberOne > 10 || numberOne < 0 || numberTwo > 10 || numberTwo < 0 { // проверка на ввод чисел от 0 до 10
			panic("Введите числа от 0 до 10 включительно")
		}
		fmt.Print(calc(numberOne, operator, numberTwo))
	}
	if !checkIfInt(numberOneInput) && checkIfInt(numberTwoInput) || checkIfInt(numberOneInput) && !checkIfInt(numberTwoInput) { //если юзер хочет арабские с римскими смешать
		panic("Нельзя совмещать арабские с римскими числами")
	}
	if !checkIfInt(numberOneInput) && !checkIfInt(numberTwoInput) { //если римские числа
		fmt.Print(calcRomans(numberOneInput, operator, numberTwoInput))
	}

}

func checkIfInt(s string) bool { //Проверка на арабские числа
	_, err := strconv.Atoi(s)
	return err == nil
}
