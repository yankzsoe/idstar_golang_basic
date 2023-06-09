package looping

import "fmt"

func LoopingSample1(loop int) {
	for i := 0; i < loop; i++ {
		fmt.Println("Angka", i)
	}
}

func LoopingSample2(loop int) {
	i := 0
	for i < loop {
		fmt.Println("Angka", i)
		i++
	}
}

func LoopingSample3(loop int) {
	i := 0
	for {
		fmt.Println("Angka", i)
		i++

		if i == 3 {
			continue
		}

		if i == loop {
			break
		}
	}
}

func LoopingSample4(loop int) {
	for i := 0; i < loop; i++ {
		for j := i; j < loop; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}

func LoopingSample5(loop int) {
outerLoop:
	for i := 0; i < loop; i++ {
		for j := 0; j < loop; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("[", i, "][", j, "]", "\n")
		}
	}
}
