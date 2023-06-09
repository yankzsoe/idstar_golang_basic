package fungsiataumethod

import (
	"math/rand"
	"time"
)

// Fungsi dengan satu nilai balikan
func GetTimestamp() int {
	return int(time.Now().Unix())
}

// Fungsi dengan dua nilai balikan, pada intinya bisa lebih dari satu.
func GetValue() (int, string) {
	val := rand.Intn(100)
	desc := "ganjil"
	if val%2 == 0 {
		desc = "genap"
	}

	return val, desc
}

// Fungsi dengan parameter tidak terbatas
func GetResult(numbers ...int) int {
	total := 0
	for _, v := range numbers {
		total += v
	}
	return total
}
