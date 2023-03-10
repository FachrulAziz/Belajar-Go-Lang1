package main

import (
	"fmt"
)

/*
func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("menghitung dari = ", i)
	}
}
*/

func main() {
	a := 1
	i := 0
	for i <= 10 {
		fmt.Println("hasil dari = ", a, "+", i, " = ", a+i)
		i = i + 1
	}

	total := 0
	for num := 0; num <= 100; num++ {
		total = total + num
	}

	fmt.Println("hasil dari penambahan hasil n adalah =", total)
}
