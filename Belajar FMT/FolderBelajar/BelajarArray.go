package katanya

import "fmt"

func BelajarArray1() {
	var mamang [4]string
	mamang[0] = "Ahmad"
	mamang[1] = "Muharam"
	mamang[2] = "Jiddan"
	mamang[3] = "Wahono"

	fmt.Println(mamang[0])
	fmt.Println(mamang[1])
	fmt.Println(mamang[2])
	fmt.Println(mamang[3])

	var values1 = [...]int{
		90,
		80,
		70,
		10,
	}
	fmt.Println(values1)
	fmt.Println(len(values1))
	fmt.Println(values1[2])
	values1[2] = 80 // ini belom terlalu jelas penggunaan nya

}
