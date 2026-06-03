package main

import "fmt"

const NMAX int = 9999

type barang [NMAX]data
type baru [NMAX]update
type data struct {
	nama, kategori string
	stok, x        int
}
type update struct {
	nama, kategori    string
	tanggal, bulan, tahun, stok, l, idx int
}

func menu(Menu *string) {
	fmt.Println("----------MENU-----------")
	fmt.Println("1. Show Data Sembako")
	fmt.Println("2. Edit Data Sembako")
	fmt.Println("3. Exit")
	fmt.Println("---------------------------")
	fmt.Print("Pilih Menu: ")
	*Menu = "0"
	for *Menu != "1" && *Menu != "2" && *Menu != "3" {
		fmt.Print("Pilih Menu: ")
		fmt.Scanln(Menu)
	}
	fmt.Println()
}

func showData(show *string) {
	fmt.Println("----------SHOW DATA SEMBAKO-----------")
	fmt.Println("1. Show All Data")
	fmt.Println("2. Show Data by Category")
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

func showAllData(x barang) {
		var k int = 1
		if x[0].s == 0 {
			fmt.Println("Data masih kosong.")
		}else {
			fmt.Println("---------------------------")
			fmt.Println("Data Sembako:")
			fmt.Println("---------------------------")
		for i := 0; i < x[0].s; i++ {
				if k < 10 {
			fmt.Printf("%d. Nama: %s, Kategori: %s, Stok: %d\n", k, x[i].nama, x[i].kategori, x[i].stok)
		
		fmt.Println("---------------------------")
			}else if k >= 10 {
				fmt.Printf("%d. Nama: %s, Kategori: %s, Stok: %d\n", k, x[i].nama, x[i].kategori, x[i].stok)
			
		fmt.Println("-------------------------")
			}
			k++
		}
		fmt.Println()
	}
}


func InputData(x *barang, y *baru) {
		var e string = "0"
		var j int = 0
		var i int = 0
		fmt.Println("----------INPUT DATA SEMBAKO-----------")
		for e != "2" {

		fmt.Println("-----------------------------")
		fmt.Print("Masukkan Nama Barang: ")
		fmt.Scanln(&x[x[0].s].nama)
			for i = 0; i < y[0].s; i++ {
				if x[x[0].s].nama == y[i].nama {
				fmt.Printf("Nama barang '%s' sudah ada. Masukkan nama barang lain.\n", x[x[0].s].nama)
					x[x[0].s].kategori = y[i].kategori
					j = 3
				}
			}
			if j == 0 {
				fmt.Print("Masukkan Kategori Barang: ")
				fmt.Scanln(&x[x[0].s].kategori)
				y[y[0].s].nama = x[x[0].s].nama
				y[y[0].s].kategori = x[x[0].s].kategori
				y[0].s++
			}
			j = 0
			fmt.Print("Masuk dan Keluar Barang: ")
			fmt.Scanln(&x[x[0].s].keterangan)
			for x[x[0].s].keterangan != "Masuk" && x[x[0].s].keterangan != "Keluar" {
				fmt.Println("Keterangan tidak valid. Masukkan 'Masuk' atau 'Keluar'.")
				fmt.Print("Masuk dan Keluar Barang: ")
				fmt.Scanln(&x[x[0].s].keterangan)
			}
			fmt.Print("Tanggal(ex:21)/Bulan(ex:11)/Tahun(ex:2026) : ")
			fmt.Scan(&x[x[0].s].tanggal)
			fmt.Scan(&x[x[0].s].bulan)
			fmt.Scan(&x[x[0].s].tahun)
			for !printdate(x[x[0].s].tanggal, x[x[0].s].bulan, x[x[0].s].tahun) {
				fmt.Println("Tanggal tidak valid. Masukkan tanggal yang benar.")
				fmt.Print("Tanggal(ex:21)/Bulan(ex:11)/Tahun(ex:2026) : ")
				fmt.Scan(&x[x[0].s].tanggal)
				fmt.Scan(&x[x[0].s].bulan)
				fmt.Scan(&x[x[0].s].tahun)
			}
			fmt.Print("Masukkan Stok Barang: ")
			fmt.Scanln(&x[x[0].s].stok)
			if x[x[0].s].keterangan == "Masuk" {
				for p := 0; p < y[0].s; p++ {
					if x[p].nama == x[x[0].s].nama {
						x[p].stok += x[x[0].s].stok
					}
				}
			}else if x[x[0].s].keterangan == "Keluar" {
				for p := 0; p < y[0].s; p++ {
					if x[x[0].s].nama == y[p].nama {
						y[p].stok -= x[x[0].s].stok
					}
				}
			}
			x[x[0].s] = x[x[0].s]
			x[0].c++

			fmt.Println("-----------------------------")
			fmt.Print("Apakah Anda ingin memasukkan data lagi? (1. Ya / 2. Tidak): ")
			fmt.Scanln(&e)
			for e != "1" && e != "2" {
			}
		}
		fmt.Println("-------------------------")
		fmt.Println("Data berhasil disimpan.")
		fmt.Println("-------------------------")
		fmt.Println()
}


func checkdate(tanggal, bulan, tahun int)bool {
	if bulan > 12 {
		return false
	}else if bulan > 31 {
		return false
	}else if bulan == 4 || bulan == 6 || bulan == 9 || bulan == 11 {
		if tanggal > 30 {
			return false
		}
	}else if bulan == 2 {
		if checkkabisat(tahun) {
			if tanggal > 29 {
				return false
			}
		}else if tanggal > 28 {
			return false
		}
	}
	return true
}

func checkkabisat(tahun int) bool {
	if tahun%4 == 0 {
			return true
		}else if tahun%400 != 0 && tahun%100 != 0 && tahun%4 == 0 {
			return true
		}else {
			return false
		}
}

func idxedit(x *barang, id *baru){
	var idxn, y, pil string
	var klo string
	var prev string
	for {
		fmt.Print("Masukkan nama barang yang ingin diubah: ")
		fmt.Scanln(&idxn)
		if idxn >= (x[0].c+1) || idxn <= 0 { 
	 	 	 	fmt.Printf("Index yang dimasukkan tidak valid\n\n") 
	 	} else {
	
	fmt.Print("Masukkan nama barang yang ingin diubah: ")
	fmt.Scanln(&nama)

	fmt.Println("1. Nama Barang")
	fmt.Println("2. Katagori") 
	fmt.Println("3. Keterangan") 
	fmt.Println("4. Tanggal")
	fmt.Println("5. Stok")  
	fmt.Println("6. Ganti semua") 

	fmt.Print("Pilih data yang ingin diubah: ")
	fmt.Scanln(&pil)
	switch pil {
	case "1":
		fmt.Print("Masukkan nama baru: ")
		fmt.Scanln(&x[idxn-1].nama)
	case "2":
		fmt.Print("Masukkan katagori baru: ")
		fmt.Scanln(&x[idxn-1].katagori)	}
	case "3":
		fmt.Print("Masukkan keterangan baru: ")
		fmt.Scanln(&x[idxn-1].keterangan)	
	case "4":
		fmt.Print("Masukkan tanggal baru (ex:21/11/2026): ")
		fmt.Scan(&x[idxn-1].tanggal)
		fmt.Scan(&x[idxn-1].bulan)
		fmt.Scan(&x[idxn-1].tahun)
		for !checkdate(x[idxn-1].tanggal, x[idxn-1].bulan, x[idxn-1].tahun) {
			fmt.Println("Tanggal tidak valid. Masukkan tanggal yang benar.")
			fmt.Print("Masukkan tanggal baru (ex:21/11/2026): ")
			fmt.Scan(&x[idxn-1].tanggal)
			fmt.Scan(&x[idxn-1].bulan)
			fmt.Scan(&x[idxn-1].tahun)
		}
	case "5":
		fmt.Print("Masukkan stok baru: ") 
		fmt.Scanln(&x[idxn-1].stok)
	}
}

func hapusData(x *barang,  *baru) {
	var nama string
	fmt.Print("Masukkan nama barang yang ingin dihapus: ")
	fmt.Scanln(&nama)
	idx := chIdx(x, &baru)
	if idx != -1 {
		for i := idx; i < NMAX-1; i++ {
			x[i] = x[i+1]
		}

		x[0].s--
		fmt.Println("Data berhasil dihapus.")
	} else {
		fmt.Println("Data tidak ditemukan.")
	}
}

func updateData(x *barang, y *baru) {
	var nama string
	fmt.Print("Masukkan nama barang yang ingin diubah: ")
	fmt.Scanln(&nama)
	idx := chIdx(x, &baru)

	if idx != -1 {
		fmt.Printf("Masukkan nama baru untuk barang '%s': ", x[idx].nama)
		fmt.Scanln(&x[idx].nama)
		fmt.Printf("Masukkan kategori baru untuk barang '%s': ", x[idx].nama)	
		fmt.Scanln(&x[idx].kategori)
		fmt.Printf("Masukkan stok baru untuk barang '%s': ", x[idx].nama)
		fmt.Scanln(&x[idx].stok)
		fmt.Println("Data berhasil diubah.")
	}
	else {
		fmt.Println("Data tidak ditemukan.")
	}
}

func showkategori(x barang) {
	var kategori string
	fmt.Print("Masukkan kategori yang ingin ditampilkan: ")
	fmt.Scanln(&kategori)
	fmt.Println("Data barang dengan kategori", kategori, ":")
	for i := 0; i < jumlahBarang; i++ {
		if x[i].kategori == kategori {
			fmt.Printf("Nama: %s, Stok: %d\n", x[i].nama, x[i].stok)
		}
	}
}

func showstok(x barang) {