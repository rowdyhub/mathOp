package main

import (
	"fmt"
	"os"
	"strconv"
)

// Возведение числа "a" в степень "b".
func exponentiation(a int, b int) int {
	var i int = 1
	var c int = a
	for i <= b-1 {
		c = c * a
		i = i + 1
	}
	return c
}
// Квадратный корень из числа "a"
//func sqr(a int) int {
//	
//}
// Вычисление факториала числа "a"
func factorial(a int) int {
	
		var i int = 1
		var b int = 1
		for i <= a {
			b = b * i
			i = i + 1
		}
		return b
	}

func main() {
	for true {
		var comand string
		fmt.Print("\n---------------------------------------------------\nВведите команду: \n")
		fmt.Fscan(os.Stdin, &comand)

		switch comand{
			case "help":
				fmt.Print("---------------------------------------------------\nСписок доступных команд:\nexit - выход;\nhelp - список команд;\nf - вычесление факториала;\ne - возведение в степень;\n")
			case "f":
				var atr int
				var c int
				fmt.Print("---------------------------------------------------\nВывести факториал числа: ")
				fmt.Fscan(os.Stdin, &atr)
				c = factorial(atr)
				fmt.Print("Ответ: " + strconv.Itoa(c))
			case "e":
				var atr1 int
				var atr2 int
				var c int
				fmt.Print("---------------------------------------------------\nВозвести число: ")
				fmt.Fscan(os.Stdin, &atr1)
				fmt.Print("В степень: ")
				fmt.Fscan(os.Stdin, &atr2)
				c = exponentiation(atr1, atr2)
				fmt.Print("Ответ: " + strconv.Itoa(c))
			case "exit":
				break
			default:
				fmt.Print("---------------------------------------------------\nКоманда не найдена! Введите help для вывода списка команд\n")
		}
	}
}