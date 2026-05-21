package main

import "fmt"

// kamus global
const Nmax int = 1000

// array untuk menyimpan data barang sembako
var sembako [Nmax]string
var kategoriBarang [Nmax]string
var jumlahBarang int

func main() {
	menu()
}

// fungsi untuk menampilkan menu
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
		menggantiBarang()
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

// pencatatan data barang masuk
func input() {
	for {
		if jumlahBarang >= Nmax {
			fmt.Println("Data penuh, tidak bisa menambah barang lagi.")
			break
		}
		fmt.Printf("Masukkan nama barang ke-%d (atau ketik 'stop' untuk berhenti): ", jumlahBarang+1)
		var nama string
		fmt.Scanln(&nama)
		if nama == "stop" {
			break
		}
		fmt.Printf("Masukkan kategori barang ke-%d: ", jumlahBarang+1)
		fmt.Scanln(&kategoriBarang[jumlahBarang])
		sembako[jumlahBarang] = nama
		jumlahBarang++
		var kembali string
		fmt.Print("Ketik 'menu' untuk kembali ke menu utama, atau Enter untuk lanjut input: ")
		fmt.Scanln(&kembali)
		if kembali == "menu" {
			return // keluar dari fungsi inputBarang, otomatis balik ke menu()
		}
	}
}

// kategori barang
func category() {
	var kategori string
	fmt.Print("Masukkan kategori barang: ")
	fmt.Scanln(&kategori)
	fmt.Println("Kategori", kategori, "belum diimplementasikan detailnya.")
}

// ubah data barang
func menggantiBarang() {
	for i := 0; i < jumlahBarang; i++ {
		fmt.Printf("Masukkan nama barang baru untuk barang ke-%d: ", i+1)
		fmt.Scanln(&sembako[i])
	}
}

// hapus data barang
func deleteItem() {
	var index int
	fmt.Print("Masukkan nomor barang yang ingin dihapus: ")
	fmt.Scanln(&index)
	if index >= 1 && index <= jumlahBarang {
		for i := index - 1; i < jumlahBarang-1; i++ {
			sembako[i] = sembako[i+1]
			kategoriBarang[i] = kategoriBarang[i+1]
		}
		jumlahBarang--
		fmt.Println("Barang berhasil dihapus.")
	} else {
		fmt.Println("Nomor barang tidak valid.")
	}
}

// log transaksi masuk dan keluar barang
func logTransaction() {
	fmt.Println("=== Log Transaksi Barang ===")
	for i := 0; i < jumlahBarang; i++ {
		fmt.Printf("Barang ke-%d: %s\n", i+1, sembako[i])
	}
	if jumlahBarang == 0 {
		fmt.Println("Belum ada transaksi barang.")
	}
}

// search data sesuai kategori (sederhana: cari nama barang)
func searchData() {
	var keyword string
	fmt.Print("Masukkan kata kunci pencarian: ")
	fmt.Scanln(&keyword)
	found := false
	for i := 0; i < jumlahBarang; i++ {
		if sembako[i] == keyword {
			fmt.Printf("Barang ditemukan di posisi ke-%d: %s\n", i+1, sembako[i])
			found = true
		}
	}
	if !found {
		fmt.Println("Barang tidak ditemukan.")
	}
}

// menampilkan data dan mengurutkan data (bubble sort sederhana)
func displayData() {
	if jumlahBarang == 0 {
		fmt.Println("Tidak ada data barang.")
		return
	}
	// bubble sort
	for i := 0; i < jumlahBarang-1; i++ {
		for j := 0; j < jumlahBarang-i-1; j++ {
			if sembako[j] > sembako[j+1] {
				sembako[j], sembako[j+1] = sembako[j+1], sembako[j]
			}
		}
	}
	fmt.Println("=== Data Barang (Terurut) ===")
	for i := 0; i < jumlahBarang; i++ {
		fmt.Printf("%d. %s\n", i+1, sembako[i])
	}
}
