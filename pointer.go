package main

import (
	"fmt"
)

// Gampangnya, pointer itu buat ngambil nilai dari function lain.
//yang mau di ambil di tandain dengan tanda bintang (*) dan ngambilnya pake tanda dan (&)

func alamat(poin *int) { // <== yang mau di ambil nilainya di tandain pake tanda bintan di depan int (*int)
	*poin = 200 // <== nah si nilainya di kasih tanda bintang juga (*poin)
}

func main() {
	var poin = 100
	alamat(&poin) // <== buat ngambilnya pake tanda dan (&poin)

	fmt.Println(poin)
}
