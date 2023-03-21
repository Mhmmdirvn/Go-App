package main

import (
	"fmt"
	"strings"
)

func main() {
	list := ListBangunDatar{BangunDatars: []BangunDatar{
		{name: "Persegi"},
		{name: "Persegi Panjang"},
		{name: "Segitiga"},
	}}

	isRunning := true
	formula := Formula{}

	for isRunning {
		// Print List Bangun Datar
		list.PrintList()

		// Input Pilihan
		formula.InputPilihan(list.BangunDatars)

		// Mau Menghitung Lagi
		var InputAgain string

		fmt.Print("Mau Menghitung Luas Apalagi (y/n) : ")
		fmt.Scanf("%s\n", &InputAgain)
		fmt.Println()

		isRunning = strings.ToLower(InputAgain) == "y"

		if !isRunning {
			if strings.ToLower(InputAgain) == "y" {
				isRunning = strings.ToLower(InputAgain) == "y"
			} else {
				fmt.Println("Terima Kasih Telah Menggunakan Aplikasi Kami !!")
			}
		}
	}
}
