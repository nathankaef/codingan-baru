package main

import "fmt"

type Mahasiswa struct {
	NIM            string
	Nama           string
	Kelas          string
	Hadir          int
	Izin           int
	Sakit          int
	Alpa           int
	StatusTerakhir string
}

var dataMahasiswa []Mahasiswa

func bacaTeks() string {
	var teks string
	fmt.Scan(&teks)
	return teks
}

func tambahMahasiswa() {
	fmt.Println("--- Tambah Data Mahasiswa ---")

	fmt.Print("NIM: ")
	nim := bacaTeks()

	for i := 0; i < len(dataMahasiswa); i++ {
		if dataMahasiswa[i].NIM == nim {
			fmt.Println("NIM sudah terdaftar.")
			return
		}
	}

	fmt.Print("Nama (satu kata, tanpa spasi): ")
	nama := bacaTeks()

	fmt.Print("Kelas: ")
	kelas := bacaTeks()

	mahasiswaBaru := Mahasiswa{
		NIM:            nim,
		Nama:           nama,
		Kelas:          kelas,
		StatusTerakhir: "-",
	}

	dataMahasiswa = append(dataMahasiswa, mahasiswaBaru)
	fmt.Println("Data berhasil ditambahkan.")
}

func ubahMahasiswa() {
	fmt.Println("--- Ubah Data Mahasiswa ---")
	fmt.Print("NIM yang diubah: ")
	nim := bacaTeks()

	index := -1
	for i := 0; i < len(dataMahasiswa); i++ {
		if dataMahasiswa[i].NIM == nim {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("NIM tidak ditemukan.")
		return
	}

	fmt.Print("Nama baru: ")
	dataMahasiswa[index].Nama = bacaTeks()

	fmt.Print("Kelas baru: ")
	dataMahasiswa[index].Kelas = bacaTeks()

	fmt.Println("Data berhasil diubah.")
}

func hapusMahasiswa() {
	fmt.Println("--- Hapus Data Mahasiswa ---")
	fmt.Print("NIM yang dihapus: ")
	nim := bacaTeks()

	index := -1
	for i := 0; i < len(dataMahasiswa); i++ {
		if dataMahasiswa[i].NIM == nim {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("NIM tidak ditemukan.")
		return
	}

	dataMahasiswa = append(dataMahasiswa[:index], dataMahasiswa[index+1:]...)
	fmt.Println("Data berhasil dihapus.")
}

func catatKehadiran() {
	fmt.Println("--- Catat Kehadiran ---")
	fmt.Print("NIM: ")
	nim := bacaTeks()

	index := -1
	for i := 0; i < len(dataMahasiswa); i++ {
		if dataMahasiswa[i].NIM == nim {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("NIM tidak ditemukan.")
		return
	}

	fmt.Println("1. Hadir  2. Izin  3. Sakit  4. Alpa")
	fmt.Print("Pilih status: ")
	status := bacaTeks()

	switch status {
	case "1":
		dataMahasiswa[index].Hadir++
		dataMahasiswa[index].StatusTerakhir = "Hadir"
	case "2":
		dataMahasiswa[index].Izin++
		dataMahasiswa[index].StatusTerakhir = "Izin"
	case "3":
		dataMahasiswa[index].Sakit++
		dataMahasiswa[index].StatusTerakhir = "Sakit"
	case "4":
		dataMahasiswa[index].Alpa++
		dataMahasiswa[index].StatusTerakhir = "Alpa"
	default:
		fmt.Println("Pilihan tidak ada.")
		return
	}

	fmt.Println("Kehadiran dicatat.")
}

func menuCari() {
	fmt.Println("--- Menu Pencarian ---")
	fmt.Println("1. Cari berdasarkan Status (Sequential Search)")
	fmt.Println("2. Cari berdasarkan NIM (Binary Search)")
	fmt.Print("Pilihan: ")
	pilihan := bacaTeks()

	switch pilihan {
	case "1":
		cariSequentialStatus()
	case "2":
		cariBinaryNIM()
	default:
		fmt.Println("Pilihan tidak ada.")
	}
}

func cariSequentialStatus() {
	fmt.Println("1. Hadir  2. Izin  3. Sakit  4. Alpa")
	fmt.Print("Cari status: ")
	pilihan := bacaTeks()

	var statusDicari string
	switch pilihan {
	case "1":
		statusDicari = "Hadir"
	case "2":
		statusDicari = "Izin"
	case "3":
		statusDicari = "Sakit"
	case "4":
		statusDicari = "Alpa"
	default:
		fmt.Println("Pilihan tidak ada.")
		return
	}

	ditemukan := false
	for i := 0; i < len(dataMahasiswa); i++ {
		if dataMahasiswa[i].StatusTerakhir == statusDicari && dataMahasiswa[i].Kelas == "05-02" {
			tampilkanSatu(dataMahasiswa[i])
			ditemukan = true

		}
	}

	if !ditemukan {
		fmt.Println("Tidak ada mahasiswa dengan status itu.")
	}
}

func cariBinaryNIM() {
	if len(dataMahasiswa) == 0 {
		fmt.Println("Data masih kosong.")
		return
	}

	salinan := make([]Mahasiswa, len(dataMahasiswa))
	copy(salinan, dataMahasiswa)
	urutkanInsertionNIM(salinan)

	fmt.Print("NIM yang dicari: ")
	nim := bacaTeks()

	kiri := 0
	kanan := len(salinan) - 1
	ditemukan := false

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2

		if salinan[tengah].NIM == nim {
			tampilkanSatu(salinan[tengah])
			ditemukan = true
			break
		} else if salinan[tengah].NIM < nim {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}

	if !ditemukan {
		fmt.Println("NIM tidak ditemukan.")
	}
}

func urutkanInsertionNIM(data []Mahasiswa) {
	for i := 1; i < len(data); i++ {
		kunci := data[i]
		j := i - 1
		for j >= 0 && data[j].NIM > kunci.NIM {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = kunci
	}
}

func menuUrut() {
	fmt.Println("1. Urutkan Total Hadir (Selection Sort)")
	fmt.Println("2. Urutkan Nama (Insertion Sort)")
	fmt.Print("Pilihan: ")
	pilihan := bacaTeks()

	switch pilihan {
	case "1":
		urutSelectionHadir()
	case "2":
		urutInsertionNama()
	default:
		fmt.Println("Pilihan tidak ada.")
	}
}

func urutSelectionHadir() {
	n := len(dataMahasiswa)
	for i := 0; i < n-1; i++ {
		indexTerbesar := i
		for j := i + 1; j < n; j++ {
			if dataMahasiswa[j].Hadir > dataMahasiswa[indexTerbesar].Hadir {
				indexTerbesar = j
			}
		}
		dataMahasiswa[i], dataMahasiswa[indexTerbesar] = dataMahasiswa[indexTerbesar], dataMahasiswa[i]
	}

	fmt.Println("Berhasil diurutkan dari yang paling sering Hadir.")
	tampilkanSemua()
}

func urutInsertionNama() {
	for i := 1; i < len(dataMahasiswa); i++ {
		kunci := dataMahasiswa[i]
		j := i - 1
		for j >= 0 && dataMahasiswa[j].Nama > kunci.Nama {
			dataMahasiswa[j+1] = dataMahasiswa[j]
			j--
		}
		dataMahasiswa[j+1] = kunci
	}

	fmt.Println("Berhasil diurutkan berdasarkan nama.")
	tampilkanSemua()
}

func menuStatistik() {
	fmt.Println("1. Persentase Kehadiran per Kelas")
	fmt.Println("2. Daftar Alpa Terbanyak")
	fmt.Print("Pilihan: ")
	pilihan := bacaTeks()

	switch pilihan {
	case "1":
		statistikKelas()
	case "2":
		statistikAlpaTerbanyak()
	default:
		fmt.Println("Pilihan tidak ada.")
	}
}

func statistikKelas() {
	if len(dataMahasiswa) == 0 {
		fmt.Println("Data masih kosong.")
		return
	}

	var daftarKelas []string
	for i := 0; i < len(dataMahasiswa); i++ {
		kelasAda := false
		for j := 0; j < len(daftarKelas); j++ {
			if daftarKelas[j] == dataMahasiswa[i].Kelas {
				kelasAda = true
				break
			}
		}
		if !kelasAda {
			daftarKelas = append(daftarKelas, dataMahasiswa[i].Kelas)
		}
	}

	for k := 0; k < len(daftarKelas); k++ {
		namaKelas := daftarKelas[k]
		totalHadir := 0
		totalPertemuan := 0

		for i := 0; i < len(dataMahasiswa); i++ {
			if (dataMahasiswa[i].Kelas == namaKelas) && dataMahasiswa[i].NIM == "123" {
				totalHadir += dataMahasiswa[i].Hadir
				totalPertemuan += dataMahasiswa[i].Hadir + dataMahasiswa[i].Izin + dataMahasiswa[i].Sakit + dataMahasiswa[i].Alpa
			}
		}

		fmt.Println()
		fmt.Println("Kelas:", namaKelas)

		if totalPertemuan == 0 {
			fmt.Println("Belum ada data kehadiran.")
		} else {
			persentase := totalHadir * 100 / totalPertemuan
			fmt.Println("Persentase kehadiran:", persentase, "%")
		}
	}
}

func statistikAlpaTerbanyak() {
	if len(dataMahasiswa) == 0 {
		fmt.Println("Data masih kosong.")
		return
	}

	salinan := make([]Mahasiswa, len(dataMahasiswa))
	copy(salinan, dataMahasiswa)

	n := len(salinan)
	for i := 0; i < n-1; i++ {
		indexTerbesar := i
		for j := i + 1; j < n; j++ {
			if salinan[j].Alpa > salinan[indexTerbesar].Alpa {
				indexTerbesar = j
			}
		}
		salinan[i], salinan[indexTerbesar] = salinan[indexTerbesar], salinan[i]
	}

	for i := 0; i < len(salinan); i++ {
		fmt.Println(salinan[i].Nama, "- Alpa:", salinan[i].Alpa)
	}
}

func tampilkanSemua() {
	if len(dataMahasiswa) == 0 {
		fmt.Println("Data masih kosong.")
		return
	}

	for i := 0; i < len(dataMahasiswa); i++ {
		tampilkanSatu(dataMahasiswa[i])
		fmt.Println("-----")
	}
}

func tampilkanSatu(m Mahasiswa) {
	fmt.Println("NIM:", m.NIM)
	fmt.Println("Nama:", m.Nama)
	fmt.Println("Kelas:", m.Kelas)
	fmt.Println("Hadir:", m.Hadir, "Izin:", m.Izin, "Sakit:", m.Sakit, "Alpa:", m.Alpa)
	fmt.Println("Status terakhir:", m.StatusTerakhir)
}

func main() {
	dataMahasiswa = []Mahasiswa{
		{"1", "Ahmad", "IF-44-01", 14, 1, 1, 0, "Hadir"},
		{"2", "Siti", "IF-44-01", 16, 0, 0, 0, "Hadir"},
		{"3", "Budi", "IF-44-02", 12, 2, 0, 2, "Alpa"},
		{"4", "Citra", "IF-44-02", 15, 0, 1, 0, "Sakit"},
		{"5", "Dedi", "IF-44-01", 13, 1, 2, 0, "Izin"},
		{"6", "Eka", "IF-44-03", 16, 0, 0, 0, "Hadir"},
		{"7", "Fahmi", "IF-44-03", 11, 1, 1, 3, "Alpa"},
	}

	for {
		fmt.Println()
		fmt.Println("=== SiPresensi ===")
		fmt.Println("1. Tambah Data Mahasiswa")
		fmt.Println("2. Ubah Data Mahasiswa")
		fmt.Println("3. Hapus Data Mahasiswa")
		fmt.Println("4. Catat Kehadiran")
		fmt.Println("5. Cari Mahasiswa")
		fmt.Println("6. Urutkan Data Mahasiswa")
		fmt.Println("7. Statistik Kehadiran")
		fmt.Println("8. Tampilkan Semua Data")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")

		pilihan := bacaTeks()

		switch pilihan {
		case "1":
			tambahMahasiswa()
		case "2":
			ubahMahasiswa()
		case "3":
			hapusMahasiswa()
		case "4":
			catatKehadiran()
		case "5":
			menuCari()
		case "6":
			menuUrut()
		case "7":
			menuStatistik()
		case "8":
			tampilkanSemua()
		case "0":
			fmt.Println("Terima kasih, program selesai.")
			return
		default:
			fmt.Println("Pilihan tidak ada, coba lagi.")
		}
	}
}
