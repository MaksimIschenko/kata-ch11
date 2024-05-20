package main

import (
	"fmt"
)

func getType(i interface{}) string {
	switch i.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case []int:
		return "[]int"
	default:
		return "Пустой интерфейс"
	}
}

func main() {
	var i interface{} = 42
	fmt.Println(getType(i)) // Вывод: "int"

	var j interface{} = "Hello, World!"
	fmt.Println(getType(j)) // Вывод: "string"

	var k interface{} = []int{1, 2, 3}
	fmt.Println(getType(k)) // Вывод: "[]int"

	var l interface{} = interface{}(nil)
	fmt.Println(getType(l)) // Вывод: "Пустой интерфейс"
}