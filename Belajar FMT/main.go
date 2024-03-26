package main

import (
	"cobaMod/CobaFolder"
	"fmt"

	test "cobaMod/katas"
)


func main() {
	Number()
	HelloWorld()
	cobafolder.CobaFile2()
	cobafolder.Mantap()
	test.Kata() //memanggil katanya itu sebagai packages bukan sebagai folder
	Boolean()
	StringAja()
	Variables1()
}

func Number() {
	satu := 1
	lima := 5
	tigakomalima := 3.5
	fmt.Println("Angka = ", satu)
	fmt.Printf("Angka = %d \n", lima)
	fmt.Println("Angka = ", tigakomalima)

}

func HelloWorld() {
	tatang := "komplit"
	marsha := "Mie ayam"
	Jiddan := 2
	Della := Jiddan * 4
	fmt.Println("mamang " + tatang)
	fmt.Println(Della)
	fmt.Printf("hallo %s %s \n", marsha, tatang) // ini di inject ke string "%s buat type string"
	fmt.Printf("anak ku ada %d", Jiddan)         // ini di inject ke string "%d buat type int"

}

func Boolean()  {
	fmt.Println("Ini", true) // ini boolean true
	fmt.Println("ini", false) // ini boolean false
}

func StringAja()  {
fmt.Println(len("mamang"))	
fmt.Println("mamang"[2])
}

func Variables1()  {
	arya := "arya"
	zia := 2
	Ryan := false
	const CobaConstant = "cobain aja ini mah" // ini Belajar Constanta
	const(
		CobaConstant2 = "Ini kalo pake langsung banyak" // ini Constanta yang sekaligus banyak
		CobaConstant3 = "ini juga"
	)

	namaDepan := "Ahmad"
	namaTengah := "Muharam Jiddan"
	namaBelakang := "Wahono"

	

	
	fmt.Println(arya)
	fmt.Println(zia)
	fmt.Println(Ryan)
	fmt.Println(CobaConstant)
	fmt.Println(CobaConstant2)
	fmt.Println(CobaConstant3)
	fmt.Println("nama gua", namaDepan, namaTengah, namaBelakang)
	
}
