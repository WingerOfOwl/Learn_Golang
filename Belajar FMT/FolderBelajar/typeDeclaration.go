package katanya

import "fmt"

func TypeDeclaration1() {
	type noKTP int // ini itu buat alias, jadi setiap deklarasiin string itu bakalan selalu pake noKTP

	var ktpGua noKTP = 1962972430
	var marlo int = 22322332332

	fmt.Println(ktpGua)
	fmt.Println(noKTP(marlo)) // ini case nya kalo misalkan udh terlanjur ada type yang belum di aliasin, bisa diubah jadi pake alias

}