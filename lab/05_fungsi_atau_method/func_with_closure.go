package fungsiataumethod

import "fmt"

// Definisi closure adalah sebuah fungsi yang bisa disimpan dalam variable
var getMinMax = func(n []int) (int, int) {
	var min, max int
	for i, value := range n {
		switch {
		case i == 0:
			min, max = value, value
		case value > max:
			max = value
		case value < min:
			min = value
		}
	}
	return min, max
}

var numbers = []int{1, 2, 3, 4, 5}

func ClosureMethod1() {
	min, max := getMinMax(numbers)
	fmt.Printf("Min %v, Max %v", min, max)
}

func ClosureMethod2() {
	add := func() func(int) int {
		sum := 0
		return func(x int) int {
			sum += x
			return sum
		}
	}

	// Panggil fungsi closure
	addFunc := add()

	// Gunakan fungsi closure
	fmt.Println(addFunc(2)) // Output: 2
	fmt.Println(addFunc(3)) // Output: 5
	fmt.Println(addFunc(5)) // Output: 10
}
