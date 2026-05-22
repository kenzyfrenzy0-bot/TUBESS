package main

import "fmt"

// kamus global
const NMAX int = 1000

// array untuk menyimpan data barang sembako
var namaBarang [NMAX]string
var kategoriBarang [NMAX]string
var jumlahBarang int

func main() {
	menu()
}

// fungsi untuk menampilkan menu
func menu() {
	fmt.Println("Menu:")
	fmt.Println("1. Masukkan nama barang ")
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
// pencatatan data barang masuk
func input() {
	for {
		if jumlahBarang >= NMAX {
			fmt.Println("Data penuh, tidak bisa menambah barang lagi.")
			break
		}
		fmt.Printf("Masukkan nama barang atau kembali ke menu: ")
		var nama string
		fmt.Scanln(&nama)
		if nama == "kembali ke menu" {
			menu() // langsung kembali ke menu utama
			return
		}

		fmt.Print("Masukkan kategori: ")
		fmt.Scanln(&kategoriBarang[jumlahBarang])
		namaBarang[jumlahBarang] = nama
		jumlahBarang++

		var kembali string
		fmt.Print("kembali ke menu atau Enter untuk lanjut input: ")
		fmt.Scanln(&kembali)
		if kembali == "kembali ke menu" {
			menu()
			return
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
		fmt.Scanln(&namaBarang[i])
	}
}

// hapus data barang
func deleteItem() {
	var index int
	fmt.Print("Masukkan nomor barang yang ingin dihapus: ")
	fmt.Scanln(&index)
	if index >= 1 && index <= jumlahBarang {
		for i := index - 1; i < jumlahBarang-1; i++ {
			namaBarang[i] = namaBarang[i+1]
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
		fmt.Printf("Barang ke-%d: %s\n", i+1, namaBarang[i])
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
		if namaBarang[i] == keyword {
			fmt.Printf("Barang ditemukan di posisi ke-%d: %s\n", i+1, namaBarang[i])
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
			if namaBarang[j] > namaBarang[j+1] {
				namaBarang[j], namaBarang[j+1] = namaBarang[j+1], namaBarang[j]
			}
		}
	}
	fmt.Println("=== Data Barang (Terurut) ===")
	for i := 0; i < jumlahBarang; i++ {
		fmt.Printf("%d. %s\n", i+1, namaBarang[i])
	}
}
