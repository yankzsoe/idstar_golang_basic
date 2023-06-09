package common

import (
	"fmt"
	"reflect"
	"strconv"
)

func WriteHeader(i interface{}) {
	fmt.Println("==============================================================")
	fmt.Println("Data Type		: ", reflect.TypeOf(i))
	fmt.Println("==============================================================")
}

func WriteLine() {
	fmt.Println("===============================================================")
}

func StrToInt(p string) (r int) {
	if p == "" {
		p = "0"
	}

	val, err := strconv.Atoi(p)
	if err != nil {
		panic(err.Error())
	}
	r = val
	return r
}

func IntToStr(i int) string {
	return strconv.Itoa(i)
}
