package main

import "fmt"

const NMAX int = 9999

type Barang struct {
	nama, kategori, keterangan string
	tanggal, bulan, tahun      int
	stok                       int
}

var dataBarang [NMAX]Barang
var jumlahBarang int = 0

func menu(Menu *string) {
    fmt.Println("----------MENU-----------")
    fmt.Println("1. Show Data Sembako")
    fmt.Println("2. Edit Data Sembako")
    fmt.Println("3. Exit")
    fmt.Println("4. Search Data")
    fmt.Println("5. Laporan Masuk/Keluar per Kategori")
    fmt.Println("---------------------------")
    fmt.Print("Pilih Menu: ")
    *Menu = "0"
    for *Menu != "1" && *Menu != "2" && *Menu != "3" && *Menu != "4" && *Menu != "5" {
        fmt.Print("Pilih Menu: ")
        fmt.Scanln(Menu)
    }
    fmt.Println()
}


func showData(show *string) {
    fmt.Println("----------SHOW DATA SEMBAKO-----------")
    fmt.Println("1. Show All Data")
    fmt.Println("2. Show Data by Keterangan (Masuk/Keluar)")
    fmt.Println("3. Show Data by Stock")
    fmt.Println("4. Back to Main Menu")
    fmt.Println("-------------------------------------")
    fmt.Print("Pilih Menu: ")
    *show = "0"
    for *show != "1" && *show != "2" && *show != "3" && *show != "4" {
        fmt.Print("Pilih Menu: ")
        fmt.Scanln(show)
    }
    fmt.Println()
}


func editData(edit *string) {
	fmt.Println("----------EDIT DATA SEMBAKO-----------")
	fmt.Println("1. Edit Data by Name")
	fmt.Println("2. Edit Data by Category")
	fmt.Println("3. Edit Data by Stock")
	fmt.Println("4. Back to Main Menu")
	fmt.Println("-------------------------------------")
	fmt.Print("Pilih Menu: ")
	*edit = "0"
	for *edit != "1" && *edit != "2" && *edit != "3" && *edit != "4" {
		fmt.Print("Pilih Menu: ")
		fmt.Scanln(edit)
	}
	fmt.Println()
}

func showAllData() {
	if jumlahBarang == 0 {
		fmt.Println("Data tidak ditemukan, silahkan isi data terlebih dahulu")
		return
	}

	fmt.Println("====================================")
	fmt.Println("Menampilkan Semua Data:")
	fmt.Println("====================================")

	k := 1
	for i := 0; i < jumlahBarang; i++ {
		if k < 10 {
			fmt.Printf("%d | Nama: %s\n  | Kategori: %s\n  | Keterangan: %s\n  | Tanggal: %d/%d/%d\n  | Stok: %d\n",
				k, dataBarang[i].nama, dataBarang[i].kategori, dataBarang[i].keterangan,
				dataBarang[i].tanggal, dataBarang[i].bulan, dataBarang[i].tahun, dataBarang[i].stok)
		} else {
			fmt.Printf("%d| Nama: %s\n  | Kategori: %s\n  | Keterangan: %s\n  | Tanggal: %d/%d/%d\n  | Stok: %d\n",
				k, dataBarang[i].nama, dataBarang[i].kategori, dataBarang[i].keterangan,
				dataBarang[i].tanggal, dataBarang[i].bulan, dataBarang[i].tahun, dataBarang[i].stok)
		}
		fmt.Println("====================================")
		k++
	}
	fmt.Println()
}

func InputData() {
	var b Barang
	var lanjut string = "1"

	for lanjut == "1" {
		fmt.Println("=====================================")
		fmt.Println("     <> Mulai Memasukkan Data <>     ")
		fmt.Println("=====================================")

		fmt.Print("Nama Barang: ")
		fmt.Scanln(&b.nama)

		fmt.Print("Kategori Barang: ")
		fmt.Scanln(&b.kategori)

		fmt.Print("Keterangan (Masuk/Keluar): ")
		fmt.Scanln(&b.keterangan)
		for b.keterangan != "Masuk" && b.keterangan != "Keluar" {
			fmt.Print("Mohon tulis Keterangan dengan format Masuk/Keluar: ")
			fmt.Scanln(&b.keterangan)
		}

		fmt.Print("Tanggal (contoh 10 5 2023): ")
		fmt.Scan(&b.tanggal, &b.bulan, &b.tahun)

		fmt.Print("Stok: ")
		fmt.Scanln(&b.stok)

		// Simpan ke array utama
		dataBarang[jumlahBarang] = b
		jumlahBarang++

		fmt.Println("Data berhasil disimpan.")
		fmt.Println("=====================================")
		fmt.Print("Lanjut isi Barang?\n1. Ya    2. Tidak\nPilihan: ")
		fmt.Scanln(&lanjut)
		for lanjut != "1" && lanjut != "2" {
			fmt.Print("Pilihan: ")
			fmt.Scanln(&lanjut)
		}
	}
	fmt.Println("=====================================")
	fmt.Println()
}

func checkdate(tanggal, bulan, tahun int) bool {
	if bulan > 12 {
		return false
	} else if tanggal > 31 {
		return false
	} else if bulan == 4 || bulan == 6 || bulan == 9 || bulan == 11 {
		if tanggal > 30 {
			return false
		}
	} else if bulan == 2 {
		if checkkabisat(tahun) {
			if tanggal > 29 {
				return false
			}
		} else if tanggal > 28 {
			return false
		}
	}
	return true
}

func checkkabisat(tahun int) bool {
    if tahun%400 == 0 {
        return true
    } else if tahun%100 == 0 {
        return false
    } else if tahun%4 == 0 {
        return true
    }
    return false
}

func updateDataByIndex() {
	var idxn, pilihan, pil int
	var lanjut string

	lanjut = "1"
	for lanjut == "1" {
		fmt.Print("Masukkan index data yang ingin diubah: ")
		fmt.Scanln(&idxn)

		if idxn > jumlahBarang || idxn <= 0 {
			fmt.Println("Index yang dimasukkan tidak valid")
		} else {
			fmt.Println("\n==================================")
			fmt.Println("Data apa yang ingin diubah:")
			fmt.Println("1. Nama Barang")
			fmt.Println("2. Kategori")
			fmt.Println("3. Keterangan")
			fmt.Println("4. Tanggal")
			fmt.Println("5. Stok")
			fmt.Println("6. Ganti semua")
			fmt.Println("==================================")
			fmt.Print("Pilihan: ")
			fmt.Scanln(&pilihan)

			switch pilihan {
			case 1:
				fmt.Println("\n==================================")
				fmt.Println("1. Hanya ganti nama pada index ini")
				fmt.Println("2. Ganti semua nama yang sama")
				fmt.Println("==================================")
				fmt.Print("Pilihan: ")
				fmt.Scanln(&pil)

				prev := dataBarang[idxn-1].nama
				fmt.Print("Nama Barang: ")
				fmt.Scanln(&dataBarang[idxn-1].nama)

				if pil == 2 {
					fmt.Printf("Mengganti semua nama barang %s menjadi %s\n\n", prev, dataBarang[idxn-1].nama)
					for ab := 0; ab < jumlahBarang; ab++ {
						if dataBarang[ab].nama == prev {
							dataBarang[ab].nama = dataBarang[idxn-1].nama
						}
					}
				} else {
					fmt.Printf("Nama barang %s diganti jadi %s\n\n", prev, dataBarang[idxn-1].nama)
				}

			case 2:
				prev := dataBarang[idxn-1].kategori
				fmt.Print("Kategori: ")
				fmt.Scanln(&dataBarang[idxn-1].kategori)
				fmt.Printf("Kategori dari barang %s diganti dari %s menjadi %s\n",
					dataBarang[idxn-1].nama, prev, dataBarang[idxn-1].kategori)

			case 3:
				prev := dataBarang[idxn-1].keterangan
				fmt.Print("Keterangan (Masuk/Keluar): ")
				fmt.Scanln(&dataBarang[idxn-1].keterangan)
				for dataBarang[idxn-1].keterangan != "Masuk" && dataBarang[idxn-1].keterangan != "Keluar" {
					fmt.Print("Mohon tulis Keterangan dengan format Masuk/Keluar: ")
					fmt.Scanln(&dataBarang[idxn-1].keterangan)
				}
				fmt.Printf("Keterangan dari barang %s diganti dari %s menjadi %s\n",
					dataBarang[idxn-1].nama, prev, dataBarang[idxn-1].keterangan)

			case 4:
				prev := fmt.Sprintf("%d/%d/%d", dataBarang[idxn-1].tanggal, dataBarang[idxn-1].bulan, dataBarang[idxn-1].tahun)
				fmt.Print("Tanggal (contoh 10 5 2023): ")
				fmt.Scan(&dataBarang[idxn-1].tanggal, &dataBarang[idxn-1].bulan, &dataBarang[idxn-1].tahun)
				for !checkdate(dataBarang[idxn-1].tanggal, dataBarang[idxn-1].bulan, dataBarang[idxn-1].tahun) {
					fmt.Println("Tanggal tidak valid, tolong masukkan tanggal dengan benar!")
					fmt.Print("Tanggal (contoh 10 5 2023): ")
					fmt.Scan(&dataBarang[idxn-1].tanggal, &dataBarang[idxn-1].bulan, &dataBarang[idxn-1].tahun)
				}
				fmt.Printf("Tanggal barang %s diganti dari %s menjadi %d/%d/%d\n",
					dataBarang[idxn-1].nama, prev,
					dataBarang[idxn-1].tanggal, dataBarang[idxn-1].bulan, dataBarang[idxn-1].tahun)

			case 5:
				prev := fmt.Sprintf("%d", dataBarang[idxn-1].stok)
				fmt.Print("Stok: ")
				fmt.Scanln(&dataBarang[idxn-1].stok)
				fmt.Printf("Stok barang %s diganti dari %s menjadi %d\n",
					dataBarang[idxn-1].nama, prev, dataBarang[idxn-1].stok)

			default:
				// Ganti semua field
				prev := fmt.Sprintf("%s/%s/%s/%d/%d/%d/%d",
					dataBarang[idxn-1].nama,
					dataBarang[idxn-1].kategori,
					dataBarang[idxn-1].keterangan,
					dataBarang[idxn-1].tanggal,
					dataBarang[idxn-1].bulan,
					dataBarang[idxn-1].tahun,
					dataBarang[idxn-1].stok)

				fmt.Print("Nama Barang: ")
				fmt.Scanln(&dataBarang[idxn-1].nama)
				fmt.Print("Kategori: ")
				fmt.Scanln(&dataBarang[idxn-1].kategori)
				fmt.Print("Keterangan (Masuk/Keluar): ")
				fmt.Scanln(&dataBarang[idxn-1].keterangan)
				fmt.Print("Tanggal (contoh 10 5 2023): ")
				fmt.Scan(&dataBarang[idxn-1].tanggal, &dataBarang[idxn-1].bulan, &dataBarang[idxn-1].tahun)
				fmt.Print("Stok: ")
				fmt.Scanln(&dataBarang[idxn-1].stok)

				fmt.Printf("Semua data lama (%s) diganti dengan data baru.", prev)
			}

			fmt.Println("Data berhasil diganti!")
			fmt.Println("==================================")
		}

		fmt.Print("Lanjut ganti data?\n1. Ya    2. Tidak\nPilihan: ")
		fmt.Scanln(&lanjut)
		for lanjut != "1" && lanjut != "2" {
			fmt.Print("Pilihan: ")
			fmt.Scanln(&lanjut)
		}
	}
}

func editDataByCategory() {
	var kategori string
	fmt.Print("Masukkan kategori yang ingin diubah: ")
	fmt.Scanln(&kategori)

	ada := false
	for i := 0; i < jumlahBarang; i++ {
		if dataBarang[i].kategori == kategori {
			ada = true
			fmt.Printf("Barang %s dengan kategori %s ditemukan.\n", dataBarang[i].nama, dataBarang[i].kategori)
			fmt.Print("Kategori baru: ")
			fmt.Scanln(&dataBarang[i].kategori)
			fmt.Println("Kategori berhasil diubah.")
		}
	}

	if !ada {
		fmt.Println("Tidak ada data dengan kategori tersebut.")
	}
}

func editDataByStock() {
	var batas int
	fmt.Print("Masukkan stok yang ingin diubah (cari barang dengan stok tertentu): ")
	fmt.Scanln(&batas)

	ada := false
	for i := 0; i < jumlahBarang; i++ {
		if dataBarang[i].stok == batas {
			ada = true
			fmt.Printf("Barang %s dengan stok %d ditemukan.\n", dataBarang[i].nama, dataBarang[i].stok)
			fmt.Print("Stok baru: ")
			fmt.Scanln(&dataBarang[i].stok)
			fmt.Println("Stok berhasil diubah.")
		}
	}

	if !ada {
		fmt.Println("Tidak ada barang dengan stok tersebut.")
	}
}

func hapusData() {
	var idxn int
	var lanjut string = "1"

	for lanjut == "1" {
		fmt.Print("Masukkan index data yang ingin dihapus: ")
		fmt.Scanln(&idxn)

		if idxn > jumlahBarang || idxn <= 0 {
			fmt.Println("Index yang dimasukkan tidak valid!")
		} else {
			// Geser data ke kiri mulai dari index yang dihapus
			for i := idxn - 1; i < jumlahBarang-1; i++ {
				dataBarang[i] = dataBarang[i+1]
			}
			jumlahBarang--
			fmt.Println("Data berhasil dihapus.")
		}

		fmt.Print("Lanjut hapus data?\n1. Ya    2. Tidak\nPilihan: ")
		fmt.Scanln(&lanjut)
		for lanjut != "1" && lanjut != "2" {
			fmt.Print("Pilihan: ")
			fmt.Scanln(&lanjut)
		}
	}
}

func copyData(src [NMAX]Barang, dest *[NMAX]Barang, jumlah int) {
	for i := 0; i < jumlah; i++ {
		(*dest)[i] = src[i]
	}
}

func sortBesar(src [NMAX]Barang, jumlah int) [NMAX]Barang {
	var dest [NMAX]Barang
	var temp Barang

	// Copy data ke array sementara
	for i := 0; i < jumlah; i++ {
		dest[i] = src[i]
	}

	// Sorting descending berdasarkan stok
	for i := 0; i < jumlah; i++ {
		for k := i + 1; k < jumlah; k++ {
			if dest[i].stok < dest[k].stok {
				temp = dest[i]
				dest[i] = dest[k]
				dest[k] = temp
			}
		}
	}

	return dest
}

func showStok() {
	if jumlahBarang == 0 {
		fmt.Println("Data tidak ditemukan, silahkan isi data terlebih dahulu")
		return
	}

	var pilihan int
	fmt.Println("==============================================")
	fmt.Println("1. Tampilkan dari jumlah stok terbanyak ke paling sedikit")
	fmt.Println("2. Tampilkan dari jumlah stok tersedikit ke paling banyak")
	fmt.Println("3. Default (sesuai urutan input)")
	fmt.Println("==============================================")
	fmt.Print("Pilihan: ")
	fmt.Scanln(&pilihan)
	for pilihan <= 0 || pilihan >= 4 {
		fmt.Print("Pilihan: ")
		fmt.Scanln(&pilihan)
	}
	fmt.Println()

	switch pilihan {
	case 3:
		fmt.Println("====================================")
		fmt.Println("    <> Menampilkan Semua Data <>    ")
		for i := 0; i < jumlahBarang; i++ {
			fmt.Println("====================================")
			fmt.Printf("%d | Nama: %s\n  | Kategori: %s\n  | Stok: %d\n",
				i+1, dataBarang[i].nama, dataBarang[i].kategori, dataBarang[i].stok)
		}
		fmt.Println("====================================")
	case 1:
		hasil := sortBesar(dataBarang, jumlahBarang)
		fmt.Println("====================================")
		fmt.Println("    <> Menampilkan Semua Data <>    ")
		for i := 0; i < jumlahBarang; i++ {
			fmt.Println("====================================")
			fmt.Printf("%d | Nama: %s\n  | Kategori: %s\n  | Stok: %d\n",
				i+1, hasil[i].nama, hasil[i].kategori, hasil[i].stok)
		}
		fmt.Println("====================================")
	case 2:
		hasil := sortBesar(dataBarang, jumlahBarang)
		fmt.Println("====================================")
		fmt.Println("    <> Menampilkan Semua Data <>    ")
		for i := jumlahBarang - 1; i >= 0; i-- {
			fmt.Println("====================================")
			fmt.Printf("%d | Nama: %s\n  | Kategori: %s\n  | Stok: %d\n",
				jumlahBarang-i, hasil[i].nama, hasil[i].kategori, hasil[i].stok)
		}
		fmt.Println("====================================")
	}
	fmt.Println()
}

func updateData() [NMAX]Barang {
	var stokBarang [NMAX]Barang
	var jumlahStok int = 0
	var i, j int
	var k int

	// Loop semua transaksi di dataBarang
	for i = 0; i < jumlahBarang; i++ {
		k = 0
		for j = 0; j < jumlahStok; j++ {
			if dataBarang[i].nama == stokBarang[j].nama {
				if dataBarang[i].keterangan == "Masuk" {
					stokBarang[j].stok += dataBarang[i].stok
				} else {
					stokBarang[j].stok -= dataBarang[i].stok
				}
				k = 1
			}
		}
		if k == 0 {
			stokBarang[jumlahStok].nama = dataBarang[i].nama
			stokBarang[jumlahStok].kategori = dataBarang[i].kategori
			stokBarang[jumlahStok].stok = 0
			if dataBarang[i].keterangan == "Masuk" {
				stokBarang[jumlahStok].stok += dataBarang[i].stok
			} else {
				stokBarang[jumlahStok].stok -= dataBarang[i].stok
			}
			jumlahStok++
		}
	}

	return stokBarang
}

func tampket() {
	var ab string
	var a int = 0

	// Input keterangan
	for ab != "Masuk" && ab != "Keluar" {
		fmt.Print("Cari barang dengan keterangan (Masuk/Keluar): ")
		fmt.Scanln(&ab)
	}

	// Cek apakah ada data dengan keterangan tersebut
	for i := 0; i < jumlahBarang; i++ {
		if dataBarang[i].keterangan == ab {
			a++
		}
	}

	if a == 0 {
		fmt.Println("Data tidak ditemukan")
	} else {
		fmt.Println("\n====================================")
		fmt.Printf("Menampilkan Semua Data dengan keterangan %s:\n", ab)
		fmt.Println("====================================")

		a = 0
		for i := 0; i < jumlahBarang; i++ {
			if dataBarang[i].keterangan == ab {
				fmt.Printf("%d | Nama: %s\n  | Kategori: %s\n  | Keterangan: %s\n  | Tanggal: %d/%d/%d\n  | Stok: %d\n",
					a+1, dataBarang[i].nama, dataBarang[i].kategori, dataBarang[i].keterangan,
					dataBarang[i].tanggal, dataBarang[i].bulan, dataBarang[i].tahun, dataBarang[i].stok)
				fmt.Println("====================================")
				a++
			}
		}
		fmt.Println()
	}
}

func tampcus() {
	var ab string
	var k int = 0

	fmt.Print("Masukkan kata kunci (Nama Barang/Kategori): ")
	fmt.Scanln(&ab)
	fmt.Println()

	for i := 0; i < jumlahBarang; i++ {
		if dataBarang[i].nama == ab || dataBarang[i].kategori == ab {
			fmt.Printf("%d | Nama: %s\n  | Kategori: %s\n  | Sisa Stok: %d\n",
				k+1, dataBarang[i].nama, dataBarang[i].kategori, dataBarang[i].stok)
			fmt.Println("====================================")
			k++
		}
	}

	if k == 0 {
		fmt.Println("Tidak dapat menemukan kata kunci yang dimasukkan")
	} else {
		fmt.Println("====================================")
	}
	fmt.Println()
}

func sortTanggal(x [NMAX]Barang, jumlah int) [NMAX]Barang {
	var dest [NMAX]Barang
	var temp Barang

	// Copy data ke array sementara
	for i := 0; i < jumlah; i++ {
		dest[i] = x[i]
	}

	// Sorting ascending berdasarkan tahun, bulan, tanggal
	for i := 0; i < jumlah; i++ {
		for j := i + 1; j < jumlah; j++ {
			if dest[i].tahun > dest[j].tahun ||
				(dest[i].tahun == dest[j].tahun && dest[i].bulan > dest[j].bulan) ||
				(dest[i].tahun == dest[j].tahun && dest[i].bulan == dest[j].bulan && dest[i].tanggal > dest[j].tanggal) {
				temp = dest[i]
				dest[i] = dest[j]
				dest[j] = temp
			}
		}
	}

	return dest
}

func showTanggal() {
	if jumlahBarang == 0 {
		fmt.Println("Data tidak ditemukan, silahkan isi data terlebih dahulu")
		return
	}

	var pilihan int
	fmt.Println("==============================================")
	fmt.Println("1. Tampilkan dari tanggal terbaru ke terlama")
	fmt.Println("2. Tampilkan dari tanggal terlama ke terbaru")
	fmt.Println("3. Default (sesuai urutan input)")
	fmt.Println("==============================================")
	fmt.Print("Pilihan: ")
	fmt.Scanln(&pilihan)
	for pilihan <= 0 || pilihan >= 4 {
		fmt.Print("Pilihan: ")
		fmt.Scanln(&pilihan)
	}
	fmt.Println()

	switch pilihan {
	case 3:
		fmt.Println("====================================")
		fmt.Println("    <> Menampilkan Semua Data <>    ")
		for i := 0; i < jumlahBarang; i++ {
			fmt.Println("====================================")
			fmt.Printf("%d | Nama: %s\n  | Kategori: %s\n  | Tanggal: %d/%d/%d\n  | Stok: %d\n",
				i+1, dataBarang[i].nama, dataBarang[i].kategori,
				dataBarang[i].tanggal, dataBarang[i].bulan, dataBarang[i].tahun, dataBarang[i].stok)
		}
		fmt.Println("====================================")
	case 1:
		hasil := sortTanggal(dataBarang, jumlahBarang)
		fmt.Println("====================================")
		fmt.Println("    <> Menampilkan Semua Data <>    ")
		for i := jumlahBarang - 1; i >= 0; i-- {
			fmt.Println("====================================")
			fmt.Printf("%d | Nama: %s\n  | Kategori: %s\n  | Tanggal: %d/%d/%d\n  | Stok: %d\n",
				jumlahBarang-i, hasil[i].nama, hasil[i].kategori,
				hasil[i].tanggal, hasil[i].bulan, hasil[i].tahun, hasil[i].stok)
		}
		fmt.Println("====================================")
	case 2:
		hasil := sortTanggal(dataBarang, jumlahBarang)
		fmt.Println("====================================")
		fmt.Println("    <> Menampilkan Semua Data <>    ")
		for i := 0; i < jumlahBarang; i++ {
			fmt.Println("====================================")
			fmt.Printf("%d | Nama: %s\n  | Kategori: %s\n  | Tanggal: %d/%d/%d\n  | Stok: %d\n",
				i+1, hasil[i].nama, hasil[i].kategori,
				hasil[i].tanggal, hasil[i].bulan, hasil[i].tahun, hasil[i].stok)
		}
		fmt.Println("====================================")
	}
	fmt.Println()
}

func showAllDataDesc() {
	var k int = 1
	if jumlahBarang == 0 {
		fmt.Printf("Data tidak ditemukan, silahkan isi data terlebih dahulu\n\n")
	} else {
		fmt.Println("====================================")
		fmt.Println("Menampilkan Semua Data: ")
		fmt.Println("====================================")
		for i := jumlahBarang - 1; i >= 0; i-- {
			if k < 10 {
				fmt.Printf("%d | Nama: %s\n  | Kategori: %s\n  | Keterangan: %s\n  | Tanggal: %d/%d/%d\n  | Stok: %d\n",
					k, dataBarang[i].nama, dataBarang[i].kategori, dataBarang[i].keterangan,
					dataBarang[i].tanggal, dataBarang[i].bulan, dataBarang[i].tahun, dataBarang[i].stok)
			} else {
				fmt.Printf("%d| Nama: %s\n  | Kategori: %s\n  | Keterangan: %s\n  | Tanggal: %d/%d/%d\n  | Stok: %d\n",
					k, dataBarang[i].nama, dataBarang[i].kategori, dataBarang[i].keterangan,
					dataBarang[i].tanggal, dataBarang[i].bulan, dataBarang[i].tahun, dataBarang[i].stok)
			}
			fmt.Println("====================================")
			k++
		}
		fmt.Println()
	}
}

func laporanMasukKeluarKategori() {
	if jumlahBarang == 0 {
		fmt.Println("Data tidak ditemukan, silahkan isi data terlebih dahulu")
		return
	}

	// Map untuk menyimpan total Masuk dan Keluar per kategori
	var masukKategori = make(map[string]int)
	var keluarKategori = make(map[string]int)

	// Loop semua data
	for i := 0; i < jumlahBarang; i++ {
		switch dataBarang[i].keterangan {
		case "Masuk":
			masukKategori[dataBarang[i].kategori] += dataBarang[i].stok
		case "Keluar":
			keluarKategori[dataBarang[i].kategori] += dataBarang[i].stok
		}
	}

	fmt.Println("====================================")
	fmt.Println(" <> Laporan Masuk/Keluar per Kategori <> ")
	fmt.Println("====================================")
	k := 1
	for kategori := range masukKategori {
		fmt.Printf("%d | Kategori: %s\n  | Total Masuk: %d\n  | Total Keluar: %d\n",
			k, kategori, masukKategori[kategori], keluarKategori[kategori])
		fmt.Println("====================================")
		k++
	}
	fmt.Println()
}
func stokKritis(batas int) {
	if jumlahBarang == 0 {
		fmt.Println("Data tidak ditemukan, silahkan isi data terlebih dahulu")
		return
	}

	var ada bool = false
	fmt.Println("====================================")
	fmt.Printf(" <> Barang dengan stok di bawah %d <> \n", batas)
	fmt.Println("====================================")

	for i := 0; i < jumlahBarang; i++ {
		if dataBarang[i].stok < batas {
			fmt.Printf("%d | Nama: %s\n  | Kategori: %s\n  | Stok: %d\n",
				i+1, dataBarang[i].nama, dataBarang[i].kategori, dataBarang[i].stok)
			fmt.Println("====================================")
			ada = true
		}
	}

	if !ada {
		fmt.Println("Tidak ada barang dengan stok kritis.")
		fmt.Println("====================================")
	}
	fmt.Println()
}

func main() {
	var pmenu, show, edit string

	fmt.Println("===========================================")
	fmt.Println("======     Algoritma Pemrograman     ======")
	fmt.Println("======          Tugas Besar          ======")
	fmt.Println("======    1301223020 & 1301220413    ======")
	fmt.Println("===========================================")
	fmt.Println("1. Continue\n2. Exit")

	pilihan := "0"
	for pilihan != "1" && pilihan != "2" {
		fmt.Print("Pilihan: ")
		fmt.Scanln(&pilihan)
	}

	if pilihan == "2" {
		fmt.Println("Keluar dari program...")
		return
	}

	fmt.Println()

	for {
    menu(&pmenu)

    switch pmenu {
    case "3":
        fmt.Println("Keluar dari program...")
        return

    case "1":
        showData(&show)
        switch show {
        case "1":
            showAllData()
        case "2":
            tampket()
        case "3":
            showStok()
        case "4":
            // kembali ke menu utama
        }

    case "2":
        editData(&edit)
        switch edit {
        case "1":
            updateDataByIndex()
        case "2":
            editDataByCategory()
        case "3":
            editDataByStock()
        case "4":
            // kembali ke menu utama
        }

    case "4":
        tampcus()

    case "5":
        laporanMasukKeluarKategori()
    }
    }
}
