package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Selamat Datang Di Aplikasi kalkulator Sederhana")
	fmt.Println("===============================================")
	operasi_perhitungan()
	
}

func operasi_perhitungan(){

	var nilai1, nilai2, hasil float64
	var operator string

		fmt.Print("Masukkan Nilai Pertama 	  : ")
		fmt.Scanf("%f\n", &nilai1)
		fmt.Print("Pilih Operator         	  : ")
		fmt.Scanf("%s\n", &operator)
		fmt.Print("Masukkan Nilai Kedua 	  : ")
		fmt.Scanf("%f\n", &nilai2)
		


	if operator == "+" {
		hasil = nilai1 + nilai2
	} else if operator == "-" {
		hasil = nilai1 - nilai2
	} else if operator == "*" {
		hasil = nilai1 * nilai2
	} else if operator == "/" {
		hasil = nilai1 / nilai2
		if nilai2 == 0{
			fmt.Println("Tidak Bisa Di Bagi Dengan Nol")
			ulangi()
		}
	} else {
		fmt.Println("Mohon Memasukkan Operasi yang benar")
	}
		fmt.Println("Hasil Perhitungan Adalah : ", hasil)

		fmt.Println("===============================================")
		ulangi()
	
	
}

	func ulangi(){
		var pilihan string

		fmt.Print("Tekan yes Untuk Mengulang lagi ? ")
		fmt.Scanf("%s\n", &pilihan)

		if strings.ToUpper(pilihan) == "YES"{
			fmt.Println()
			operasi_perhitungan()
		} else {
			fmt.Println("Selesai....")
		}

	}