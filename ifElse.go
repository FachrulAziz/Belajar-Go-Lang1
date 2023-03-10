package main

import (
	"fmt"
)

func main() {
	// jika integer
	i := 70
	if i < 100 {
		fmt.Println(i, " lebih kecil dari 100")
	} else {
		fmt.Println(i, " lebih besar dari 100")
	}

	// jika string
	balon := "hijau"      // jelaskan dahulu variablenya
	if balon == "hijau" { // inti dari if. kalo hasil dari balon == ini bukan hijau, maka pilihan kedua akan muncul
		fmt.Println("DOOR.. Meletus balon hijau")
	} else {
		fmt.Println("Party berlanjut")
	}

}
