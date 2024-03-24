package main

import "fmt"

func HelloWorld() {
	 tatang:= "komplit"
	 marsha:= "Mie ayam"
	 Jiddan:= 2
	 Della:= Jiddan*4
	fmt.Println("mamang " + tatang )
	fmt.Println(Della)
	fmt.Printf("hallo %s %s \n",  marsha, tatang) // ini di inject ke string "%s buat type string"
	fmt.Printf("anak ku ada %d" , Jiddan) // ini di inject ke string "%d buat type int"
	
}


