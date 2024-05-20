package main

import(
	"testing"
	"strings"
)

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

type testCases struct {
	f func(xs ...interface{}) interface{}
	args []interface{}
	expected interface{}
}

func TestOperate(t *testing.T) {
	testCases := []testCases{
		{f: Concat, args: []interface{}{"Hello, ", "World!"}, expected: "Hello, World!"},
		{f: Sum, args:  []interface{}{1, 2, 3, 4, 5}, expected: 15},
		{f: Sum, args: []interface{}{1.1, 2.2, 3.3, 4.4, 5.5}, expected: 16.5},
	}

	for _, tc := range testCases {
		res := Operate(tc.f, tc.args...)
		if res != tc.expected {
			t.Errorf("Unexpected result.")
		}
	}
}