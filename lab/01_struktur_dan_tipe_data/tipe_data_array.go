package tipe_data

import (
	"fmt"
	tools "lab/common"
)

func ArrayStringSample() {
	// string array
	animals := [5]string{
		"kecoa",
		"jangkrik",
		"katak",
		"ular",
	}

	for i, v := range animals {
		fmt.Printf("index %v value %v\n", i, v)
	}
	tools.WriteLine()

	animals[4] = "cicak"
	for i, v := range animals {
		fmt.Printf("index %v value %v\n", i, v)
	}
	tools.WriteLine()

	// Array Multidimensi
	var numbers = [2][3]int{{1, 2}, {3, 4, 5}}

	fmt.Println(numbers)
	fmt.Println(numbers[0][1])
	fmt.Println(numbers[1][2])
	tools.WriteLine()
}

func ArrayObjectSample() {
	type people struct {
		index int
		name  string
	}

	peoples := [5]people{
		{
			index: 1,
			name:  "Aa",
		},
		{
			index: 2,
			name:  "Abe",
		},
		{
			index: 3,
			name:  "Acil",
		},
		{
			index: 4,
			name:  "Agus",
		},
		{
			index: 5,
			name:  "Agay",
		},
	}

	var antrian []people
	antrian = append(antrian, people{
		index: 1,
		name:  "pertama",
	})

	antrian[0] = people{
		index: 2,
		name:  "kedua",
	}

	antrianCopy := make([]people, len(antrian))
	copy(antrianCopy, antrian)

	for _, v := range antrianCopy {
		fmt.Println(v)
	}

	tools.WriteLine()

	for _, v := range peoples {
		fmt.Println(v)
	}
}
