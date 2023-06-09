package tipe_data

import (
	"fmt"
	tools "lab/common"
)

// Map adalah tipe data asosiatif yang ada di Go, berbentuk key-value pair. Untuk setiap data (atau value) yang disimpan, disiapkan juga key-nya.
// Key harus unik, karena digunakan sebagai penanda (atau identifier) untuk pengaksesan value yang bersangkutan.
func MapSample() {
	mapNilaiPerBulan := map[string]int{}

	mapNilaiPerBulan["Januari"] = 50
	mapNilaiPerBulan["Februari"] = 60
	mapNilaiPerBulan["Maret"] = 75
	mapNilaiPerBulan["April"] = 75
	mapNilaiPerBulan["Mei"] = 90
	mapNilaiPerBulan["Mei"] = 95 // Nilai Mei sebelumnya akan hilang dan digantikan dengan nilai 95

	for i, v := range mapNilaiPerBulan {
		fmt.Printf("%s value %v\n", i, v)
	}
	tools.WriteLine()
	desember := "Desember"
	var value, isExist = mapNilaiPerBulan[desember]
	if isExist {
		fmt.Printf("Exist with value: %v\n", value)
	} else {
		fmt.Println("not exist")
	}
	tools.WriteLine()

	mapNilaiPerBulan[desember] = 93
	value, isExist = mapNilaiPerBulan[desember]
	if isExist {
		fmt.Printf("Exist with value: %v\n", value)
	} else {
		fmt.Println("not exist")
	}
	tools.WriteLine()

	// Map Kombinasi dengan Array/Slice
	// Masing-masing Map bisa memiliki key yang berbeda
	peoples := []map[string]string{
		{"name": "Aa", "gender": "Male", "dob": "17 Agustus 1945"},
		{"name": "Aci", "gender": "Female"},
		{"name": "Acil", "gender": "Female"},
	}

	for _, people := range peoples {
		// fmt.Println(i, v)
		fmt.Println(people["name"], people["gender"], people["dob"])
	}
}
