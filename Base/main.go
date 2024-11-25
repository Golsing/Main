package main

/*
Калькулятор индекса массы тела
16 и менее	Выраженный дефицит массы тела
16-18,5	Недостаточная (дефицит) масса тела
18,5-25	Норма
25-30	Избыточная масса тела (предожирение)
30-35	Ожирение первой степени
35-40	Ожирение второй степени
40 и более	Ожирение третьей степени (морбидное)
*/
import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2
	var userHeight float64
	var userKg float64
	fmt.Print("Введите свой рост: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userKg)
	IMT := userKg / math.Pow(userHeight, IMTPower)
	fmt.Printf("Ваш индекс массы тела : %0.f", IMT)

}
