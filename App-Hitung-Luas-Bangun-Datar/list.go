package main

import "fmt"

type ListBangunDatar struct {
	BangunDatars []BangunDatar
}

func (list *ListBangunDatar) PrintList() {
	fmt.Println("Selamat Datang Di Aplikasi Hitung Luas Bangun Datar")
	fmt.Println("===================================================")
	for i, list := range list.BangunDatars {
		fmt.Println(i+1, ".", list.name )
	}
	fmt.Println("===================================================")
}