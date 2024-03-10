package main

import "fmt"

func calc(numberOne int, operator string, numberTwo int) int { //калькулятор арабских чисел

	switch operator {
	case "+":
		return numberOne + numberTwo
	case "-":
		return numberOne - numberTwo
	case "*":
		return numberOne * numberTwo
	case "/":
		return numberOne / numberTwo
	default:
		fmt.Println("Don't corrected operator")
		return 0
	}

}
