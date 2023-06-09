package tipe_data

import (
	"fmt"
	"math"

	tools "lab/common"
)

func NumericSample() {
	intSample := 123
	tools.WriteHeader(intSample)
	fmt.Println("Bilangan integer :", intSample)

	decNumb := 4.050000 // sama seperti "var decNumb = 4.050000" atau "var decNumb float64 = 4.050000"
	tools.WriteHeader(decNumb)

	// Float
	fmt.Println("Bilangan desimal tanpa menggunakan tanda floating	:", decNumb)
	fmt.Printf("Bilangan desimal menggunakan tanda %%f		: %f \n", decNumb)
	fmt.Printf("Bilangan desimal menggunakan tanda %%.3f	: %.3f \n", decNumb)
	fmt.Printf("Bilangan desimal menggunakan tanda %%.9f	: %.9f \n", decNumb)

	fmt.Println("Bilangan desimal menggunakan math.Ceil()	:", math.Ceil(decNumb))
	fmt.Println("Bilangan desimal menggunakan math.Floor()	:", math.Floor(decNumb))
	fmt.Println("Bilangan desimal menggunakan math.Exp()	:", math.Exp(decNumb))
	fmt.Println("Bilangan desimal menggunakan math.Exp2()	:", math.Exp2(decNumb))
	fmt.Print("\n")
}
