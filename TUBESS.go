
package main

import "fmt"

const NMAX int = 12345

type data [NMAX]Barang
type logupdate [NMAX]update

type update struct {
	nama, kategori string
	stok, count    int
}
type Barang struct {
	nama, kategori, ket                     string
	tanggal, bulan, tahun, stok, count, idb int
}

// ===== Menu Utama =====
func menu(pmenu *string) {
	selesai := false
	for !selesai {
		fmt.Println("======= Menu Utama =======")
		fmt.Println("1. Tampilkan Data Sembako")
		fmt.Println("2. Edit Data Sembako")
		fmt.Println("3. Keluar")
		fmt.Println("==========================")

		*pmenu = "0"
		for *pmenu != "1" && *pmenu != "2" && *pmenu != "3" {
			fmt.Print("Pilihan: ")
			fmt.Scanln(pmenu)
		}

		if *pmenu == "3" {
			selesai = true
		} else {
			// biar menu muncul lagi setelah aksi
		}
		fmt.Println()
	}
}

// ===== Menu Tampil =====
func tampil(tamp *string) {
	selesai := false
	for !selesai {
		fmt.Println("======== Inventori Sembako ========")
		fmt.Println("1. Tampilkan Semua Data Sembako")
		fmt.Println("2. Tampilkan berdasarkan stok Sembako")
		fmt.Println("3. Tampilkan Sesuai Keterangan")
		fmt.Println("4. Tampilkan Sesuai Kata Kunci")
		fmt.Println("5. Tampilkan berdasarkan tanggal")
		fmt.Println("6. Kembali")
		fmt.Println("===========================================")

		*tamp = "0"
		for *tamp != "1" && *tamp != "2" && *tamp != "3" && *tamp != "4" && *tamp != "5" && *tamp != "6" {
			fmt.Print("Pilihan: ")
			fmt.Scanln(tamp)
		}

		if *tamp == "6" {
			selesai = true
		}
		fmt.Println()
	}
}

// ===== Menu Edit =====
func edit(edt *string) {
	selesai := false
	for !selesai {
		fmt.Println("========= Aplikasi Inventori Sembako ==========")
		fmt.Println("1. Masukan Data Sembako")
		fmt.Println("2. Mengganti Data Sembako")
		fmt.Println("3. Menghapus Data Sembako")
		fmt.Println("4. Kembali")
		fmt.Println("==============================================")

		*edt = "0"
		for *edt != "1" && *edt != "2" && *edt != "3" && *edt != "4" {
			fmt.Print("Pilihan: ")
			fmt.Scanln(edt)
		}

		if *edt == "4" {
			selesai = true
		}
		fmt.Println()
	}
}

// ===== Tampilkan Semua =====
func tampilsemua(x Data) {
	var k int = 1
	if x[0].count == 0 {
		fmt.Printf("Data tidak ditemukan, silahkan isi data terlebih dahulu\n\n")
	} else {
		fmt.Println("====================================")
		fmt.Println("Menampilkan Semua Data: ")
		fmt.Println("====================================")
		for i := 0; i < x[0].count; i++ {
			if k < 10 {
				fmt.Printf("%d | Nama: %s\n  | Kategori: %s\n  | Keterangan: %s\n  | Tanggal: %d/%d/%d\n  | Stok: %d\n",
					k, x[i].nama, x[i].kategori, x[i].ket, x[i].tanggal, x[i].bulan, x[i].tahun, x[i].stok)
				fmt.Println("====================================")
			} else {
				fmt.Printf("%d| Nama: %s\n  | Kategori: %s\n  | Keterangan: %s\n  | Tanggal: %d/%d/%d\n  | Stok: %d\n",
					k, x[i].nama, x[i].kategori, x[i].ket, x[i].tanggal, x[i].bulan, x[i].tahun, x[i].stok)
				fmt.Println("====================================")
			}
			k++
		}
		fmt.Println()
	}
}

// ===== Validasi Tanggal =====
func validasiTanggal(tanggal, bulan, tahun int) bool {
	if tahun < 1 || tahun > 9999 {
		return false
	}
	if bulan < 1 || bulan > 12 {
		return false
	}
	if tanggal < 1 || tanggal > 31 {
		return false
	}
	// Cek jumlah hari dalam bulan
	if bulan == 2 {
		if (tahun%4 == 0 && tahun%100 != 0) || (tahun%400 == 0) {
			return tanggal <= 29
		} else {
			return tanggal <= 28
		}
	}
	if bulan == 4 || bulan == 6 || bulan == 9 || bulan == 11 {
		return tanggal <= 30
	}
	return true
}

// ===== Edit Data =====
func editdata(x *Data, log *LogUpdate) {
	var nama, kategori, ket string
	var tanggal, bulan, tahun, stok int
	fmt.Print("Masukan Nama Sembako: ")
	fmt.Scanln(&nama)
	fmt.Print("Masukan Kategori Sembako: ")
	fmt.Scanln(&kategori)
	fmt.Print("Masukan Keterangan Sembako: ")
	fmt.Scanln(&ket)
	fmt.Print("Masukan Tanggal (DD MM YYYY): ")
	fmt.Scanln(&tanggal, &bulan, &tahun)	
	fmt.Print("Masukan Stok Sembako: ")
	fmt.Scanln(&stok)
	if !validasiTanggal(tanggal, bulan, tahun) {
		fmt.Printf("Tanggal tidak valid, silahkan coba lagi\n\n")
		return
	}
	// Cek apakah data sudah ada
	for i := 0; i < x[0].count; i++ {
		if x[i].nama == nama && x[i].kategori == kategori {	
			fmt.Printf("Data sudah ada, silahkan coba lagi\n\n")
			return
		}	
	}
	// Masukan data baru
	x[x[0].count] = Barang{nama, kategori, ket, tanggal, bulan, tahun, stok, x[0].count + 1, x[0].count + 1}
	log[log[0].count] = update{nama, kategori, ket, tanggal, bulan, tahun, stok}
	x[0].count++
	log[0].count++
	fmt.Printf("Data berhasil ditambahkan\n\n")
}

// ===== Hapus Data =====
func hapusdata(x *Data, log *LogUpdate) {
	var nama, kategori string
	fmt.Print("Masukan Nama Sembako: ")
	fmt.Scanln(&nama)
	fmt.Print("Masukan Kategori Sembako: ")
	fmt.Scanln(&kategori)

	// Cek apakah data ada
	for i := 0; i < x[0].count; i++ {
		if x[i].nama == nama && x[i].kategori == kategori {
			// Hapus data
			for j := i; j < x[0].count-1; j++ {
				x[j] = x[j+1]
			}
			x[0].count--
			fmt.Printf("Data berhasil dihapus\n\n")
			return
		}
	}
	fmt.Printf("Data tidak ditemukan\n\n")
}
// ===== Sort & Tampilan Stok =====
func tampilstok(x Data) {
	if x[0].count == 0 {
		fmt.Printf("Data tidak ditemukan, silahkan isi data terlebih dahulu\n\n")
	} else {
		// Sort data berdasarkan stok
		for i := 0; i < x[0].count-1; i++ {
			for j := 0; j < x[0].count-i-1; j++ {
				if x[j].stok < x[j+1].stok {
					x[j], x[j+1] = x[j+1], x[j]
				}
			}
		}
		// Tampilkan data
		fmt.Println("====================================")	
		fmt.Println("Menampilkan Data Berdasarkan Stok: ")
		fmt.Println("====================================")
		for i := 0; i < x[0].count; i++ {
			fmt.Printf("%d. Nama: %s\n   Kategori: %s\n   Keterangan: %s\n   Tanggal: %d/%d/%d\n   Stok: %d\n",
				i+1, x[i].nama, x[i].kategori, x[i].ket, x[i].tanggal, x[i].bulan, x[i].tahun, x[i].stok)
			fmt.Println("====================================")
		}
		fmt.Println()
	}
}

func main() {
	var data Data
	var log LogUpdate
	var menuUtama, menuTampil, menuEdit string
	menu(&menuUtama)
	if menuUtama == "1" {
		menu(&menuTampil)
		if menuTampil == "1" {
			tampilsemua(data)
		} else if menuTampil == "2" {
			tampilstok(data)
		}
	} else if menuUtama == "2" {
		menu(&menuEdit)
		if menuEdit == "1" {
			editdata(&data, &log)
		} else if menuEdit == "3" {
			hapusdata(&data, &log)
		}
	}	
}	



