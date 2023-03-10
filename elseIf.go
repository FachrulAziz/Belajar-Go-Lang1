// If dengan tipe data boolean yang hasilnya hanya ada 2 jawaban (true/benar atau false/salah)
package main

import (
	"fmt"
)

func bayarUtang(uang int) (pesan string) {
	utang := 5000

	if uang > utang {
		pesan = ("uangmu kebanyakan")
	} else if uang == utang {
		pesan = ("uangmu pas")
	} else {
		pesan = ("uangnya kurang kampret")
	}
	return
}
func main() {
	fmt.Println(bayarUtang(50000))
}
