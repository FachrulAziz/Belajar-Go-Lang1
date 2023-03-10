package main

import (
	"fmt"
)

func lampuMerah(lampu string) (pesan string) {
	switch lampu {
	case "hijau":
		pesan = "silahkan jalan dan hati-hati di jalan"
	case "kuning":
		pesan = "hati-hati"
	case "merah":
		pesan = "stop"
	default: // default ini dipake kalo jawabannya gaada di daftar case atau jawabannya error
		pesan = "Lampu sedang error"
	}
	return
}

func main() {
	fmt.Println(lampuMerah("kuning"))
}
