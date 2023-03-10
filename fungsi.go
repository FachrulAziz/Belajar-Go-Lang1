package main

import (
	"fmt"     // fmt untuk print
	"strconv" //  strcronv untuk convert dari integer ke string atau sebaliknya
)

func multiply(angka1 int, angka2 int) int {
	return angka1 * angka2
}

func getBiography(age int, name string, status string) string { // script parameter dan nilai
	ageNow := strconv.Itoa(age) // strconv.Itoa untuk convert integer ke string
	return name + " adalah seorang " + status + " saat ini berusia " + ageNow + " tahun"
}

func main() {
	fmt.Println("Fungsi ini menghasilkan ", multiply(10, 10))
	fmt.Println("Hasil perkalian ", multiply(70, 89))
	fmt.Println("=============================================")
	fmt.Println(getBiography(27, "Fachrul Aziz", "Programmer"))
	fmt.Println(getBiography(27, "Deyana Putri Nurma Intani", "Terapis Wicara"))
	fmt.Println(getBiography(2, "Naura Almeera Fachrul", "Anak Bayi"))
}
