package main

import (
	"fmt"
	"os"
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
func restart(){
	var comand string
	fmt.Print("\n---------------------------------------------------\n")
	fmt.Print("Введите команду: \n")
	fmt.Fscan(os.Stdin, &comand)

	switch comand{
		case "help":
			fmt.Print("---------------------------------------------------\n")
			fmt.Print("Список доступных команд:\n")
			fmt.Print("exit - выход;\n")
			fmt.Print("help - список команд;\n")
			fmt.Print("f - вычесление факториала;\n")
			fmt.Print("e - возведение в степень;\n")
			restart()
		case "f":
			var atr int
			var c int
			fmt.Print("---------------------------------------------------\n")
			fmt.Print("Вывести факториал числа: ")
			fmt.Fscan(os.Stdin, &atr)
			c = factorial(atr)
			fmt.Print("Ответ: ")
			fmt.Print(c)

			restart()
		case "e":
			var atr1 int
			var atr2 int
			var c int
			fmt.Print("---------------------------------------------------\n")
			fmt.Print("Возвести число: ")
			fmt.Fscan(os.Stdin, &atr1)
			fmt.Print("В степень: ")
			fmt.Fscan(os.Stdin, &atr2)
			c = exponentiation(atr1, atr2)
			fmt.Print("Ответ: ")
			fmt.Print(c)

			restart()			
		case "exit":
			os.Exit(0)
		default:
			fmt.Print("---------------------------------------------------\n")
			fmt.Print("Каманда не найдена! Введите help для вывода списка команд\n")
			restart()
}
}

func main() {
	restart()
}