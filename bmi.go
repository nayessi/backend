package main

import "fmt"

func main() {
	var weight float64
	var height float64

	fmt.Print("Введите вес (кг): ")
	fmt.Scan(&weight)

	fmt.Print("Введите рост (см): ")
	fmt.Scan(&height)

	heightInMeters := height / 100
	bmi := weight / (heightInMeters * heightInMeters)

	fmt.Printf("Ваш ИМТ: %.2f\n", bmi)

	if bmi < 18.5 {
		fmt.Println("Недостаточный вес")
	} else if bmi < 25 {
		fmt.Println("Нормальный вес")
	} else if bmi < 30 {
		fmt.Println("Избыточный вес")
	} else {
		fmt.Println("Ожирение")
	}
}