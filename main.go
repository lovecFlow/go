package main //Определение пакета. Является стартом любой программы на Go.

import "fmt" //Импорт сторонних пакетов для использования их функционала. Данный пакет нужен для реализации форматирвоания входных и выходных данных.

func main() {
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello" + "World")
}
