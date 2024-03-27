package katanya

import "fmt"

func OperasiMatematika() {
	a := 10
	b := 17

	pertambahan := a + b
	pengurangan := a - b
	perkalian := a * b
	pembagian := b / a
	modulo := b % a // modulo ini adalah sisa dari pembagian contoh 5 dibagi 3 maka modulo nya adalah 2

	fmt.Println(pertambahan)
	fmt.Println(pengurangan)
	fmt.Println(perkalian)
	fmt.Println(pembagian)
	fmt.Println(modulo)

}