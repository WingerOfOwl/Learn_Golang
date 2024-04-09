package katanya

import (
	"fmt"

	
)

func TipeDataMap() {
	person := map[string]string{
		"nama":   "jiddan",
		"alamat": "jakarta",
	}

	fmt.Println(person["nama"])
	fmt.Println(person["alamat"])
	fmt.Println(person)

	number1 := map[int]string{
		1 : "ini 1",
		2: "ini 2",
	}

	fmt.Println(number1[1])
	fmt.Println(number1[2])
	fmt.Println(number1)


	book := make(map[string] string)
	book["title"] = "cinta bersemi"
	book["volume"] = "300"
	book["Penerbit"] = "Jordan"

	delete(book, "Penerbit")

	fmt.Println(book)


}