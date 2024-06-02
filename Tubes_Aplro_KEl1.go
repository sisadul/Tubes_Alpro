package main

import (
	"fmt"
	"sort"
)

const NMAX = 100

type data struct {
	pabrikan string
	produk   mobil
}

type mobil struct {
	nama string
	tk   int
}

type tabMobil [NMAX]data

type pabrikanPenjualan struct {
	pabrikan string
	jumlah   int
}

func main() {
	var A tabMobil
	menu(&A)
}

func menu(A *tabMobil) {
	//{I.S A adalah array tabMobil yang akan digunakan untuk operasi menu
	// F.S Menampilkan menu aplikasi dealer mobil dan melakukan operasi sesuai pilihan pengguna}
	var pilihan int
	var n int = 0
	var y int

	for y == 0 {
		fmt.Println("Aplikasi Dealer Mobil")
		fmt.Println("1. Input Data mobil")
		fmt.Println("2. Kelola Data Pabrikan")
		fmt.Println("3. Tampilkan Daftar Mobil")
		fmt.Println("4. Tampilkan 3 Pabrikan dengan Penjualan Tertinggi")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			inputDataMobil(A, &n)
		case 2:
			kelolaDatapabrikan(A, &n)
		case 3:
			tampilkanDaftarMobil(A, n, "")
		case 4:
			tampilkanTigaPabrikanTeratas(A, n)
		case 5:
			y = 1
			keluar()
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

func inputDataMobil(A *tabMobil, n *int) {
	//{I.S A adalah array tabMobil yang akan diisi dengan data mobil, n adalah jumlah data yang sudah ada sebelum input
	// F.S Data mobil baru berhasil diinput ke dalam array A dan nilai n bertambah sesuai jumlah data yang diinput}
	var y int
	fmt.Print("Masukkan jumlah data mobil yang akan diinput: ")
	fmt.Scan(&y)
	for i := 0; i < y; i++ {
		fmt.Print("Masukkan data mobil\nTahun Keluaran: ")
		fmt.Scan(&A[*n].produk.tk)
		fmt.Print("Jenis Mobil: ")
		fmt.Scan(&A[*n].produk.nama)
		fmt.Print("Asal Pabrikan: ")
		fmt.Scan(&A[*n].pabrikan)
		(*n)++
		fmt.Println("Data berhasil diinput")
	}
}

func kelolaDatapabrikan(A *tabMobil, n *int) {
	/*{I.S Terdapat data mobil yang disimpan dalam sebuah array A dengan ukuran n.

	  	F.S Jika tidak ada data mobil yang tersedia (n == 0), pesan "Tidak ada data mobil yang tersedia." akan ditampilkan, dan proses kelolaDataPabrikan akan selesai.
	  Jika terdapat data mobil yang tersedia (n > 0):
	  Akan ditampilkan daftar pabrikan yang tersedia beserta jumlah mobil dari setiap pabrikan.
	  Pengguna diminta untuk memasukkan nama pabrikan yang ingin dikelola.
	  Setelah memilih pabrikan, pengguna akan diberikan menu pilihan untuk mengelola data mobil dari pabrikan tersebut.
	  Pengguna dapat memilih untuk menambah data mobil, menghapus data mobil, mengedit data mobil, menampilkan daftar mobil, atau kembali ke menu utama
	  Proses akan terus berlanjut hingga pengguna memilih untuk kembali ke menu utama (pilihan 5).
	  Proses akan terus berlanjut hingga pengguna memilih untuk kembali ke menu utama (pilihan 5).}*/
	if *n == 0 {
		fmt.Println("Tidak ada data mobil yang tersedia.")
		return
	}

	// Tampilkan daftar pabrikan yang ada
	pabrikanUnik := make(map[string]int)
	for i := 0; i < *n; i++ {
		if A[i].pabrikan != "" {
			pabrikanUnik[A[i].pabrikan]++
		}
	}
	fmt.Println("Daftar Pabrikan:")
	for pabrikan, jumlah := range pabrikanUnik {
		fmt.Printf("%s: %d mobil\n", pabrikan, jumlah)
	}

	// Pilih pabrikan yang ingin dikelola
	fmt.Print("Masukkan nama pabrikan: ")
	var pabrikanPilihan string
	fmt.Scan(&pabrikanPilihan)

	// Tampilkan menu pilihan untuk mengelola pabrikan
	var y int
	for y == 0 {
		fmt.Println("\nMenu Pabrikan", pabrikanPilihan)
		fmt.Println("1. Tambah Data Mobil")
		fmt.Println("2. Hapus Data Mobil")
		fmt.Println("3. Edit Data Mobil")
		fmt.Println("4. Tampilkan Daftar Mobil")
		fmt.Println("5. Kembali ke Menu Utama")
		fmt.Print("Pilih: ")
		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahDataMobil(A, n, pabrikanPilihan)
		case 2:
			hapusDataMobil(A, n, pabrikanPilihan)
		case 3:
			editDataMobil(A, n, pabrikanPilihan)
		case 4:
			tampilkanDaftarMobil(A, *n, pabrikanPilihan)
		case 5:
			y = 1
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func tambahDataMobil(A *tabMobil, n *int, pabrikanPilihan string) {
	/*{I.S Terdapat sebuah array A yang merupakan kumpulan data mobil.

	  	F.S Jika array A sudah penuh (n >= NMAX), maka pesan "Array penuh! Tidak dapat menambahkan data mobil baru." akan ditampilkan, dan tidak ada penambahan data mobil baru yang dilakukan.
	  Jika array A belum penuh (n < NMAX):
	  Pengguna diminta untuk memasukkan data mobil baru, termasuk tahun keluaran dan jenis mobil.
	  Data mobil baru tersebut akan ditambahkan ke array A dengan informasi pabrikan yang telah dipilih sebelumnya.
	  Jumlah data mobil (n) akan bertambah satu.
	  Pesan "Data mobil baru telah ditambahkan." akan ditampilkan sebagai konfirmasi penambahan data.}*/
	if *n >= NMAX {
		fmt.Println("Array penuh! Tidak dapat menambahkan data mobil baru.")
		return
	}

	// Minta data mobil baru
	fmt.Println("Masukkan data mobil baru:")
	fmt.Print("Tahun Keluaran: ")
	var tahunKeluaran int
	fmt.Scan(&tahunKeluaran)
	fmt.Print("Jenis Mobil: ")
	var jenisMobil string
	fmt.Scan(&jenisMobil)

	// Tambahkan data mobil baru ke array
	A[*n].produk.tk = tahunKeluaran
	A[*n].produk.nama = jenisMobil
	A[*n].pabrikan = pabrikanPilihan
	(*n)++

	fmt.Println("Data mobil baru telah ditambahkan.")
}

func hapusDataMobil(A *tabMobil, n *int, pabrikanPilihan string) {
	/*{I.S Terdapat sebuah array A yang merupakan kumpulan data mobil

	  	F.S Jika tidak ditemukan mobil dari pabrikan yang dipilih (idx == -1), maka akan ditampilkan pesan "Pabrikan tidak ditemukan." dan tidak ada perubahan pada array A.
	  Jika ditemukan mobil dari pabrikan yang dipilih:
	  Array A akan diurutkan berdasarkan tahun keluaran mobil (tk).
	  Seluruh mobil dari pabrikan yang dipilih akan ditampilkan beserta nomor urutnya.
	  Pengguna diminta untuk memasukkan nomor urut mobil yang ingin dihapus.
	  Jika nomor urut tidak valid (di luar rentang yang benar), maka akan ditampilkan pesan "Nomor urut tidak valid." dan tidak ada perubahan pada array A.
	  Jika nomor urut valid, data mobil dengan nomor urut tersebut akan dihapus dari array A.
	  Jumlah data mobil (n) akan berkurang satu.
	  Pesan "Data mobil telah dihapus." akan ditampilkan sebagai konfirmasi penghapusan.}*/

	// Urutkan array berdasarkan tahun keluaran (tk)
	sort.Slice(A[:*n], func(i, j int) bool {
		return A[i].produk.tk < A[j].produk.tk
	})

	// Sequential search untuk menemukan semua mobil dari pabrikan pilihan
	var indeksMobil []int
	for i := 0; i < *n; i++ {
		if A[i].pabrikan == pabrikanPilihan {
			indeksMobil = append(indeksMobil, i)
		}
	}
	if len(indeksMobil) == 0 {
		fmt.Println("Pabrikan tidak ditemukan.")
		return
	}

	// Tampilkan semua mobil yang sesuai
	fmt.Println("\nDaftar Data Mobil:")
	for i, idx := range indeksMobil {
		fmt.Printf("%d. %d %s %s\n", i+1, A[idx].produk.tk, A[idx].produk.nama, A[idx].pabrikan)
	}
	fmt.Print("Masukkan nomor urut mobil yang ingin dihapus: ")
	var nomorUrut int
	fmt.Scan(&nomorUrut)

	// Validasi nomor urut
	if nomorUrut < 1 || nomorUrut > len(indeksMobil) {
		fmt.Println("Nomor urut tidak valid.")
		return
	}

	// Hapus data mobil dari array
	idxHapus := indeksMobil[nomorUrut-1]
	for i := idxHapus; i < *n-1; i++ {
		A[i] = A[i+1]
	}
	A[*n-1] = data{}
	(*n)-- // Decrement n

	fmt.Println("Data mobil telah dihapus.")
}

func editDataMobil(A *tabMobil, n *int, pabrikanPilihan string) {
	/*{I.S Terdapat sebuah array A yang merupakan kumpulan data mobil.
	  F.S Jika tidak ditemukan mobil dari pabrikan yang dipilih (len(indeksMobil) == 0), maka akan ditampilkan pesan "Pabrikan tidak ditemukan." dan tidak ada perubahan pada array A.
	  Jika ditemukan mobil dari pabrikan yang dipilih:
	  Seluruh mobil dari pabrikan yang dipilih akan ditampilkan beserta nomor urutnya.
	  Pengguna diminta untuk memasukkan nomor urut mobil yang ingin diubah.
	  Jika nomor urut tidak valid (di luar rentang yang benar), maka akan ditampilkan pesan "Nomor urut tidak valid." dan tidak ada perubahan pada array A.
	  Jika nomor urut valid, pengguna diminta untuk memasukkan data baru untuk mobil (tahun keluaran baru dan jenis mobil baru).
	  Data mobil dengan nomor urut tersebut akan diubah sesuai dengan data baru yang dimasukkan.
	  Pesan "Data mobil telah diubah." akan ditampilkan sebagai konfirmasi pengeditan.}*/

	// Sequential search untuk menemukan semua mobil dari pabrikan pilihan
	var indeksMobil []int
	for i := 0; i < *n; i++ {
		if A[i].pabrikan == pabrikanPilihan {
			indeksMobil = append(indeksMobil, i)
		}
	}

	if len(indeksMobil) == 0 {
		fmt.Println("Pabrikan tidak ditemukan.")
		return
	}

	fmt.Println("\nDaftar Data Mobil dari Pabrikan:", pabrikanPilihan)
	for i, idx := range indeksMobil {
		fmt.Printf("%d. %d %s\n", i+1, A[idx].produk.tk, A[idx].produk.nama)
	}

	// Minta nomor urut mobil yang mau diubah
	fmt.Print("Masukkan nomor urut mobil yang ingin diubah: ")
	var nomorUrut int
	fmt.Scan(&nomorUrut)

	// Validasi nomor urut
	if nomorUrut < 1 || nomorUrut > len(indeksMobil) {
		fmt.Println("Nomor urut tidak valid.")
		return
	}

	// Meminta data baru untuk mobil
	fmt.Print("Tahun Keluaran Baru: ")
	var tahunKeluaranBaru int
	fmt.Scan(&tahunKeluaranBaru)
	fmt.Print("Jenis Mobil Baru: ")
	var jenisMobilBaru string
	fmt.Scan(&jenisMobilBaru)

	// Ubah data mobil di array
	idxEdit := indeksMobil[nomorUrut-1]
	A[idxEdit].produk.tk = tahunKeluaranBaru
	A[idxEdit].produk.nama = jenisMobilBaru

	fmt.Println("Data mobil telah diubah.")
}

func tampilkanDaftarMobil(A *tabMobil, n int, pabrikanPilihan string) {
	/*{I.S Terdapat sebuah array A yang merupakan kumpulan data mobil.

	  	F.S Akan ditampilkan daftar mobil yang memenuhi kriteria berikut:
	  Jika pabrikanPilihan tidak ditentukan (pabrikanPilihan == ""), maka akan ditampilkan semua mobil dari array A.
	  Jika pabrikanPilihan ditentukan, maka akan ditampilkan hanya mobil-mobil dari pabrikan tersebut.
	  Daftar mobil akan ditampilkan secara urut berdasarkan tahun keluaran mobil (tk), menggunakan algoritma selection sorting.
	  Setiap entri daftar mobil akan menampilkan informasi tahun keluaran, jenis mobil, dan nama pabrikan.
	  Nomor urut akan ditampilkan di depan setiap entri daftar mobil.}*/
	fmt.Println("Daftar Mobil:")
	var mobilList []data
	for i := 0; i < n; i++ {
		if pabrikanPilihan == "" || A[i].pabrikan == pabrikanPilihan {
			mobilList = append(mobilList, A[i])
		}
	}

	for i := 0; i < len(mobilList)-1; i++ {
		minIndex := findMin(mobilList, i)
		mobilList[i], mobilList[minIndex] = mobilList[minIndex], mobilList[i]
	}

	for i, mobil := range mobilList {
		fmt.Printf("%d. %d %s %s\n", i+1, mobil.produk.tk, mobil.produk.nama, mobil.pabrikan)
	}
}

func findMin(arr []data, start int) int {
	// Fungsi akan mengembalikan indeks dari elemen dalam array arr yang memiliki nilai produk.tk minimum, dimulai dari indeks start.
	minIndex := start
	for i := start + 1; i < len(arr); i++ {
		if arr[i].produk.tk < arr[minIndex].produk.tk {
			minIndex = i
		}
	}
	return minIndex
}

func findMax(arr []data, start int) int {
	// findMax adalah fungsi umum yang mencari indeks elemen maksimum dalam array berdasarkan fungsi pembanding
	maxIndex := start
	for i := start + 1; i < len(arr); i++ {
		if arr[i].produk.tk > arr[maxIndex].produk.tk {
			maxIndex = i
		}
	}
	return maxIndex
}

func tampilkanTigaPabrikanTeratas(A *tabMobil, n int) {
	//{I.S  A adalah array tabMobil yang berisi data penjualan mobil, n adalah jumlah data yang ada
	// F.S Menampilkan 3 pabrikan dengan penjualan tertinggi dari data yang diberikan}
	if n == 0 {
		fmt.Println("Tidak ada data mobil yang tersedia.")
	}

	// Menghitung jumlah penjualan per pabrikan secara manual
	var pabrikanPenjualan [NMAX]pabrikanPenjualan
	var jumlahPabrikan int

	for i := 0; i < n; i++ {
		if A[i].pabrikan != "" {
			found := false
			for j := 0; j < jumlahPabrikan; j++ {
				if pabrikanPenjualan[j].pabrikan == A[i].pabrikan {
					pabrikanPenjualan[j].jumlah++
					found = true
				}
			}
			if !found {
				pabrikanPenjualan[jumlahPabrikan].pabrikan = A[i].pabrikan
				pabrikanPenjualan[jumlahPabrikan].jumlah = 1
				jumlahPabrikan++
			}
		}
	}

	// Urutkan array pabrikanPenjualan menggunakan insertion sort
	for i := 1; i < jumlahPabrikan; i++ {
		temp := pabrikanPenjualan[i]
		j := i - 1
		for j >= 0 && pabrikanPenjualan[j].jumlah < temp.jumlah {
			pabrikanPenjualan[j+1] = pabrikanPenjualan[j]
			j--
		}
		pabrikanPenjualan[j+1] = temp
	}

	// Tampilkan 3 pabrikan teratas
	fmt.Println("3 Pabrikan dengan Penjualan Tertinggi:")
	for i := 0; i < 3 && i < jumlahPabrikan; i++ {
		fmt.Printf("%d. %s dengan %d mobil\n", i+1, pabrikanPenjualan[i].pabrikan, pabrikanPenjualan[i].jumlah)
	}
}

func keluar() {
	//{I.S
	// F.S Menampilkan pesan terima kasih setelah menggunakan aplikasi dealer mobil.}
	fmt.Println("Terima kasih telah menggunakan aplikasi dealer mobil.")
}
