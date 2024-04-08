package katanya

import "fmt"

func TypeDataSlice() {
	names := [...]string{
		"mamang", "dadang", "Jiddan", "mase",
	}
	slice := names[2:4] //ini ditentukan batas awal dan batas akhirnya
	fmt.Println(slice)

	slice1 := names[:2] // ini hanya ditentukan batas akhirnya
	fmt.Println(slice1)

	slice2 := names[1:] // ini hanya ditentukan batas awalnya saja
	fmt.Println(slice2)

	sliceAll := names[:]
	fmt.Println(sliceAll)

	// Kode Program Appand Slice

	days := [...] string {
		"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu",
	}
	daySlice1 := days[5:]
	daySlice1[0] = "Ini Sabtu Yang Baru" // ini buat ngubah index 0 dari pointer yaitu hari sabtu
	daySlice1[1] = "Ini Minggu Yang Baru" // ini buat ngubah index pertama dari pointer yaitu hari minggu
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "INI LIBUR BARU") // ini pointernya diambil dari daySlice1
	// ini karena ukuran Array nya udh penuh sampe minggu jadinya dibuat array baru. Dan ini titik mulainya dari pointer dan karena ditambah sama INI LIBUR BARU jadi dia masuk ke array baru
	daySlice2[0] = "mantap"
	fmt.Println(daySlice2)


	newSlice := make([]string, 2,5)
	newSlice[0] = "ini 1"
	newSlice[1] = "ini 2"
	// newSlice[2] = "ini 2"  kalau menggunakan cara ini pasti akan error makanya harus menggunakan Append
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "ini yang ke 3", "ini yang ke 4")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))


	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice)) // ini dia bisa otomatis ngikutin panjang dan capasity dari yang mau di copy

	copy(toSlice, fromSlice) // ini cara bacanya copy kemana dan dari mana
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	iniArray:= [...]int{1,2,3,4,5} //ini cara menggunakan Array dia ada titik 3 nya
	iniSlice:= []int{1,2,3,4,5} // ini cara menggunakan slice tidak ada titik 3 nya

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}