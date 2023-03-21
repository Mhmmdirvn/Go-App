package main

import "fmt"

type Formula struct {
	BangunDatar BangunDatar
	Sisi   float32
	Lebar  float32
	Tinggi float32
	Alas   float32
}


func (formula *Formula) InputPilihan(BangunDatars []BangunDatar) {
	var selectedBangundatar int
	fmt.Print("Masukkan pilihan \t: ")
	fmt.Scanf("%d\n", &selectedBangundatar)
	formula.BangunDatar = BangunDatars[selectedBangundatar-1]

	switch selectedBangundatar {
	case 1: 
			formula.LuasPersegi()
	case 2: 
			formula.LuasPersegiPanjang()
	case 3: 
			formula.LuasSegitiga()
	default:
		fmt.Println("Mohon Input Pilihan Yang Benar ")
	}
}

func (formula *Formula) LuasPersegi() {
	// Input Sisi Persegi
	fmt.Print("Masukkan Sisi Persegi \t: ")
	fmt.Scanf("%f\n", &formula.Sisi)

	// Hitung Luas Persegi
	luasPersegi := formula.Sisi * formula.Sisi
	fmt.Println("Luas Persegi Adalah \t:",luasPersegi)
	fmt.Println()
}

func (formula *Formula) LuasPersegiPanjang() {
	// Input Lebar Persegi Panjang
	fmt.Print("Masukkan lebar \t\t: ")
	fmt.Scanf("%f\n", &formula.Lebar)

	// Input Tinggi Persegi Panjang
	fmt.Print("Masukkan Tinggi \t: ")
	fmt.Scanf("%f\n", &formula.Tinggi)

	// Hitung Luas Persegi Panjang
	luasPersegiPanjang := formula.Lebar * formula.Tinggi
	fmt.Println("Luas Persegi Panjang \t:",luasPersegiPanjang)
	fmt.Println()
}

func (formula *Formula) LuasSegitiga() {
	// Input Alas Segitiga
	fmt.Print("Masukkan Alas \t\t: ")
	fmt.Scanf("%f\n", &formula.Alas)

	// Input Tinggi Segitiga
	fmt.Print("Masukkan Tinggi \t: ")
	fmt.Scanf("%f\n", &formula.Tinggi)

	// Hitung Luas Segitiga
	luasSegitiga := 0.5 * (formula.Alas * formula.Tinggi)
	fmt.Println("Luas Segitiga Adalah \t:",luasSegitiga)
	fmt.Println()
}


