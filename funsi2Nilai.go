package main

import (
	"fmt"
)

func biodata(nama string, alamat string, umur int) (bio string, umurNanti int) {
	umurNanti = umur + 10
	bio = nama + " Beralamat di " + alamat
	return
}

func main() {
	a, b := biodata("Naura", "Taman Kirana Surya", 2)

	fmt.Println(a)
	fmt.Println("dan 10 tahun kemudian umurnya", b, " tahun")
}
