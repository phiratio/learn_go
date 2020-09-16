package main

import "fmt"

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func main() {
	var op = Sum([]int{1, 2, 3, 4, 5})
	fmt.Println(op)
}
