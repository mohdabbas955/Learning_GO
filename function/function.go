package main

import "fmt"

func AddtwoNumber(a int, b int) int {
	return a + b
}

func AddthreeNumber(a int, b int, c int) int {
	return a + b + c
}

func main() {
	res := AddtwoNumber(3, 4)
	fmt.Println(res)

	res = AddthreeNumber(4, 5, 6)
	fmt.Println(res)

}
