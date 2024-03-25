package main
import "cobaMod/CobaFolder"
import "fmt"
import test "cobaMod/katas"


func main() {
	Number()
	HelloWorld()
	cobafolder.CobaFile2()
	cobafolder.Mantap()
	test.Kata() //memanggil katanya itu sebagai packages bukan sebagai folder
	Boolean()
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
