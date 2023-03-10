package main

import (
	"fmt"     // fmt untuk print
	"strconv" //  strcronv untuk convert dari integer ke string atau sebaliknya
)

// Variable - Tipe Data "String" Tipe data string bisa berupa huruf, nomor atau kalimat yang ditandai dengan tanda kutip ""
var namaDepan, namaBelakang = "Fachrul", "Aziz"
var namaLengkap = namaDepan + " " + namaBelakang

// Variable - Tipe Data Numerik (Int/Integer) hanya berisikan nomor dan hasil tidak ada koma. jika float, ada komanya integer ditandai dengan tidak memakai tanda kutip ""
var umur, satu, dua, tiga = 27, 15, 10, 20
var hasil = satu + dua

// const / konstan adalah variable yang hasilnya tidak bisa dirubah, ditambah atau sebagainya
const a = 50

// strconv untuk convert menggunakan Itoa = Integer ke String dan Atoi = String ke Integer
var variabelSatu, variabelDua, variabelTiga = 30, 20, "100"
var gajian = strconv.Itoa(variabelSatu)
var kiw, _ = strconv.Atoi(variabelTiga)
var totalGaji = kiw + 90

func main() {
	fmt.Println("Hallo semuanya, perkenalkan nama saya " + namaLengkap) // jika String memakai tanda tambah (+)
	fmt.Println("Umur saya ", umur, "Tahun")                            // jika Integer memakai tanda koma (,)
	fmt.Println("saya akan menghitung. jika 15 + 10 maka hasilnya =", hasil)
	fmt.Println("Gaji saya saat ini ", a)
	fmt.Println("Percobaan convert = ", gajian)
}
