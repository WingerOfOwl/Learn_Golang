package katanya

import "fmt"

func OperasiBoolean() {
	Jiddan := 90
	Della := 80
	iniNilaiJiddan := Jiddan >= 90
	iniNilaiDella := Della >= 90

	Kelulusan := iniNilaiDella && iniNilaiJiddan
	SalahSatu := iniNilaiDella || iniNilaiJiddan
	fmt.Println(Kelulusan)
	fmt.Println(SalahSatu)

}