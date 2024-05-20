package main

import (
	"strings"
)

var Operate func(f func(xs ...interface{}) interface{}, i ...interface{}) interface{} // реализуй меня

var Concat func(xs ...interface{}) interface{} // реализуй меня для string

var Sum func(xs ...interface{}) interface{} // реализуй меня для int и float64

func init() {

	Operate = func(f func(xs ...interface{}) interface{}, i ...interface{}) interface{} {
		switch i[0].(type) {

		case string:
			Concat = func(xs ...interface{}) interface{} {
				var loc_result strings.Builder
				for _, x := range xs {
					if str, ok := x.(string); ok {
						loc_result.WriteString(str)
					}
				}
				return loc_result.String()
			}
			return Concat(i...)

		case int:
			Sum = func(xs ...interface{}) interface{} {
				var loc_result int = 0
				for _, x := range xs {
					loc_result += x.(int)
				}
				return loc_result
			}
			return Sum(i...)

		case float64:
			Sum = func(xs ...interface{}) interface{} {
				var loc_result float64
				for _, x := range xs {
					loc_result += x.(float64)
				}
				return loc_result
			}
			return Sum(i...)
		default:
			return nil
		}
	}
}

//func main() {
//	fmt.Println(Operate(Concat, "Hello, ", "World!"))  // Вывод: "Hello, World!"
//	fmt.Println(Operate(Sum, 1, 2, 3, 4, 5))           // Вывод: 15
//	fmt.Println(Operate(Sum, 1.1, 2.2, 3.3, 4.4, 5.5)) // Вывод: 16.5
//}
