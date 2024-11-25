package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2
	userHeight := 1.8
	userKg := 100.0
	fmt.Print("Введите свой рост: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userKg)
	IMT := userKg / math.Pow(userHeight, IMTPower)
	fmt.Print(IMT)
}
