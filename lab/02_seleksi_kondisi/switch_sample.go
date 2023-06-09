package seleksikondisi

import "fmt"

func SwitchSample1(point int) {
	switch point {
	case 9, 10:
		fmt.Println("perfect")
	case 8:
		fmt.Println("awesome")
	case 7:
		fmt.Println("good")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("belajar lebih giat lagi")
		}
	}
}

func SwitchSample2(point int) {
	switch {
	case point <= 10 || point > 10:
		fmt.Println("perfect")
	case point == 8:
		fmt.Println("awesome")
	case point == 7:
		fmt.Println("good")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("belajar lebih giat lagi")
		}
	}
}

func SwitchSample3(point int) {
	switch {
	case point <= 10 || point > 10:
		fmt.Println("perfect")
	case point == 8:
		fmt.Println("awesome")
		fallthrough
	case point == 7:
		fmt.Println("good")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("belajar lebih giat lagi")
		}
	}
}
