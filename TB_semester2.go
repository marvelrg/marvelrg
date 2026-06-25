package main

import "fmt"

type Setoran struct {
	ID      int
	Berat   float64
	Tanggal string
}

type Warga struct {
	ID      int
	Nama    string
	Setoran []Setoran
}

var daftarWarga = []Warga{
	{
		ID:   1,
		Nama: "Muhammad Rachman",
		Setoran: []Setoran{
			{
				ID:      1,
				Berat:   2.5,
				Tanggal: "2025-06-01",
			},
			{
				ID:      1,
				Berat:   3.2,
				Tanggal: "2025-06-10",
			},
		},
	},
	{
		ID:   2,
		Nama: "Febrando Hasby",
		Setoran: []Setoran{
			{
				ID:      2,
				Berat:   1.8,
				Tanggal: "2025-06-05",
			},
			{
				ID:      2,
				Berat:   2.1,
				Tanggal: "2025-06-12",
			},
			{
				ID:      2,
				Berat:   1.5,
				Tanggal: "2025-06-15",
			},
		},
	},
	{
		ID:   3,
		Nama: "Muhammad ILham Orion",
		Setoran: []Setoran{
			{
				ID:      2,
				Berat:   4.0,
				Tanggal: "2025-06-08",
			},
		},
	},
}

func main() {
	var pilih int
	fmt.Println("tst")
	for {
		fmt.Println("\n=== MENU WASTE-TRACK ===")
		fmt.Println("1. Kelola Data Warga | 2. Kelola Setoran | 3. Pencarian | 4. Pengurutan | 5. Statistik | 0. Keluar")
		fmt.Print("Pilih Menu: ")
		_, err := fmt.Scan(&pilih)
		if err != nil {
			fmt.Println("Input tidak valid!")
			fmt.Scanln()
			continue
		}

		switch pilih {
		case 1: menuKelolaWarga()
		case 2: menuKelolaSetoran()
		case 3: fiturPencarian()
		case 4: fiturPengurutan()
		case 5: fiturStatistik()
		case 0: return
		default: fmt.Println("Menu tidak valid!")
		}
	}
}

func menuKelolaWarga() {
	var pilih int
	fmt.Println("\n-- Kelola Data Warga --")
	fmt.Println("1. Tambah | 2. Ubah | 3. Hapus | 4. Lihat Semua | 0. Kembali")
	fmt.Print("Pilihan: ")
	fmt.Scan(&pilih)

	if pilih == 1 {
		var w Warga
		fmt.Print("ID: "); fmt.Scan(&w.ID)
		fmt.Print("Nama: "); fmt.Scan(&w.Nama)
		daftarWarga = append(daftarWarga, w)
	} else if pilih == 2 {
		var id int
		fmt.Print("Masukkan ID yang akan diubah: ")
		fmt.Scan(&id)
		for i := 0; i < len(daftarWarga); i++ {
			if daftarWarga[i].ID == id {
				fmt.Print("Nama Baru: "); fmt.Scan(&daftarWarga[i].Nama)
			}
		}
	} else if pilih == 3 {
		var id int
		fmt.Print("Masukkan ID yang akan dihapus: ")
		fmt.Scan(&id)
		for i := 0; i < len(daftarWarga); i++ {
			if daftarWarga[i].ID == id {
				daftarWarga = append(daftarWarga[:i], daftarWarga[i+1:]...)
				fmt.Println("Data terhapus.")
				break
			}
		}
	} else if pilih == 4 {
		for i := 0; i < len(daftarWarga); i++ {
			fmt.Printf("%d. ID: %d | Nama: %s\n", i+1, daftarWarga[i].ID, daftarWarga[i].Nama)
		}
	}
}

func menuKelolaSetoran() {
	var id int
	fmt.Print("Masukkan ID Warga: ")
	fmt.Scan(&id)
	for i := 0; i < len(daftarWarga); i++ {
		if daftarWarga[i].ID == id {
			var s Setoran
			s.ID = len(daftarWarga[i].Setoran) + 1
			fmt.Print("Berat (kg): "); fmt.Scan(&s.Berat)
			fmt.Print("Tanggal: "); fmt.Scan(&s.Tanggal)
			daftarWarga[i].Setoran = append(daftarWarga[i].Setoran, s)
			return
		}
	}
	fmt.Println("Warga tidak ditemukan!")
}

func fiturPencarian() {
	var pilih int
	fmt.Println("1. Sequential (ID) | 2. Binary (Nama)")
	fmt.Scan(&pilih)
	if pilih == 1 {
		var id int
		var total float64
		fmt.Print("Cari ID: "); fmt.Scan(&id)
		for i := 0; i < len(daftarWarga); i++ {
			if daftarWarga[i].ID == id {
				for j := 0; j < len(daftarWarga[i].Setoran);j++ {
					total += daftarWarga[i].Setoran[j].Berat
				}	
				fmt.Printf("Ketemu:\n")
				fmt.Printf("ID: %d \n", daftarWarga[i].ID)
				fmt.Printf("Nama: %s \n", daftarWarga[i].Nama)
				fmt.Printf("Total: %v kg \n", total)
				return
			}
		}
		fmt.Println("Tidak ditemukan.")
	} else {

		for i := 0; i < len(daftarWarga)-1; i++ {
			for j := i + 1; j < len(daftarWarga); j++ {
				if daftarWarga[i].Nama > daftarWarga[j].Nama {
					daftarWarga[i], daftarWarga[j] = daftarWarga[j], daftarWarga[i]
				}
			}
		}
		var nama string
		fmt.Print("Cari Nama: "); fmt.Scan(&nama)
		low, high := 0, len(daftarWarga)-1
		for low <= high {
			mid := (low + high) / 2
			if daftarWarga[mid].Nama == nama {
				fmt.Printf("Ketemu:\n")
				fmt.Printf("ID: %d\n", daftarWarga[mid].ID)
				fmt.Printf("Nama: %s\n", daftarWarga[mid].Nama)
				return
			} else if daftarWarga[mid].Nama < nama {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
		fmt.Println("Tidak ditemukan.")
	}
}

func fiturPengurutan() {
	fmt.Println("1. Selection Sort | 2. Insertion Sort")
	var opt int
	fmt.Scan(&opt)
	n := len(daftarWarga)

	if opt == 1 { // Selection Sort
		for i := 0; i < n-1; i++ {
			max := i
			for j := i + 1; j < n; j++ {
				var bJ, bMax float64
				for k := 0; k < len(daftarWarga[j].Setoran); k++ { bJ += daftarWarga[j].Setoran[k].Berat }
				for k := 0; k < len(daftarWarga[max].Setoran); k++ { bMax += daftarWarga[max].Setoran[k].Berat }
				if bJ > bMax { max = j }
			}
			daftarWarga[i], daftarWarga[max] = daftarWarga[max], daftarWarga[i]
		}
	} else { 
		for i := 1; i < n; i++ {
			key := daftarWarga[i]
			var bKey float64
			for k := 0; k < len(key.Setoran); k++ { bKey += key.Setoran[k].Berat }
			j := i - 1
			for j >= 0 {
				var bJ float64
				for k := 0; k < len(daftarWarga[j].Setoran); k++ { bJ += daftarWarga[j].Setoran[k].Berat }
				if bJ < bKey {
					daftarWarga[j+1] = daftarWarga[j]
					j--
				} else { break }
			}
			daftarWarga[j+1] = key
		}
	}
	fmt.Println("Data telah diurutkan.")
}

func fiturStatistik() {
	var total float64
	for i := 0; i < len(daftarWarga); i++ {
		for j := 0; j < len(daftarWarga[i].Setoran); j++ {
			total += daftarWarga[i].Setoran[j].Berat
		}
	}
	fmt.Printf("Total Sampah: %.2f kg\n", total)
}