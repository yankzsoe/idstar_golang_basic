package tipe_data

import (
	"fmt"
	"strings"
)

func NonNumericSample() {
	// Boolean
	y := true
	fmt.Println("bool dengan nilai true: ", y)
	y = false
	fmt.Println("bool dengan nilai false: ", y)

	// String
	var str3 string
	str3 = "hi"

	str1 := "Hello"
	str2 := "World World"
	// String Concate
	result := str1 + " " + str2 + str3
	fmt.Println(result)

	result = strings.Replace(result, "World", "Guys", -1)
	fmt.Println("Replace result: ", result)

}
