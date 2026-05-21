package main

import "fmt"

//kamus global
//array untuk menyimpan data barang sembako

const Nmax int = 1000

type Sembako [1000]string

//type jumlahBarang int

func main() {

}

//fungsi untuk menampilkan menu
func menu() {
	fmt.Println("Menu:")
	fmt.Println("1. Input data barang masuk")
	fmt.Println("2. Kategori barang")
	fmt.Println("3. Ubah data barang")
	fmt.Println("4. Hapus data barang")
	fmt.Println("5. Log transaksi masuk dan keluar barang")
	fmt.Println("6. Search data sesuai kategori")
	fmt.Println("7. Tampilkan data dan urutkan data")
	fmt.Print("Pilih menu: ")
	var pilihan int
	fmt.Scanln(&pilihan)
	switch pilihan {
	case 1:
		input()
	case 2:
		category()
	case 3:
		change()
	case 4:
		deleteItem()
	case 5:
		logTransaction()
	case 6:
		searchData()
	case 7:
		displayData()
	}
}

//pencatatan data barang masuk
func input() {
	for i := 0; i < Nmax; i++ {
		fmt.Printf("Masukkan nama barang ke-%d: ", i+1)
		fmt.Scanln(&sembako[i])
		jumlahBarang++
	}
}

//kategori barang
func category() {
	var kategori string
	fmt.Print("Masukkan kategori barang: ")
	fmt.Scanln(&kategori)
	//logika untuk mengkategorikan barang berdasarkan input kategori
}

//ubah data barang

func menggantiBarang() {
	for i := 0; i < jumlahBarang; i++ {
		fmt.Printf("Masukkan nama barang baru untuk barang ke-%d: ", i+1)
		fmt.Scanln(&sembako[i])
	}
}

//hapus data barang

func deleteItem() {
	var index int
	fmt.Print("Masukkan nomor barang yang ingin dihapus: ") //nomor barang dimulai dari 1
	fmt.Scanln(&index)
	if index >= 1 && index <= jumlahBarang {
		for i := index - 1; i < jumlahBarang-1; i++ {
			sembako[i] = sembako[i+1]
		}
		jumlahBarang--
		fmt.Println("Barang berhasil dihapus.")
	} else {
		fmt.Println("Nomor barang tidak valid.")
	}
}

//LOG TRANSACTION masuk dan keluar barang

func logTransaction() {

}

//search data seusai kategori
func searchData() {

}

//menapilkan data dan mengurutkan data
func displayData() {

}
