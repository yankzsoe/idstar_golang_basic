package seleksikondisi

import "fmt"

func IfElseSample1(point int) {
	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else {
		fmt.Printf("tidak lulus. nilai anda %d\n", point)
	}
}

func IfElseSample2(point int) {
	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else {
		if point == 4 {
			fmt.Println("kurang bagus, tapi lulus :p")
		} else {
			fmt.Printf("tidak lulus. nilai anda %d\n", point)
		}
	}
}
