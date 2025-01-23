package main

import (
	"fmt"
)

type Barang struct {
	Nama   string
	Harga  float64
	Jumlah int
	Total  float64
}

type Kasir struct {
	DaftarBarang []Barang
	TotalHarga   float64
}

func (k *Kasir) TambahBarang(namaBarang string, harga float64, jumlah int) {
	total := harga * float64(jumlah)
	barang := Barang{
		Nama:   namaBarang,
		Harga:  harga,
		Jumlah: jumlah,
		Total:  total,
	}
	k.DaftarBarang = append(k.DaftarBarang, barang)
	k.TotalHarga += total
	fmt.Printf("%s x%d ditambahkan ke daftar belanja.\n", namaBarang, jumlah)
}

func (k *Kasir) TampilkanRincian() {
	fmt.Println("\n--- Rincian Belanja ---")
	for _, barang := range k.DaftarBarang {
		fmt.Printf("%s - Harga: %.2f - Jumlah: %d - Total: %.2f\n", barang.Nama, barang.Harga, barang.Jumlah, barang.Total)
	}
	fmt.Printf("\nTotal Belanja: %.2f\n", k.TotalHarga)
}

func (k *Kasir) Bayar(uangDibayar float64) {
	if uangDibayar >= k.TotalHarga {
		kembalian := uangDibayar - k.TotalHarga
		fmt.Printf("\nTotal yang harus dibayar: %.2f\n", k.TotalHarga)
		fmt.Printf("Uang yang dibayar: %.2f\n", uangDibayar)
		fmt.Printf("Kembalian: %.2f\n", kembalian)
	} else {
		fmt.Println("\nUang yang dibayar kurang dari total belanja. Transaksi dibatalkan.")
	}
}

func main() {
	var kasir Kasir

	for {
		fmt.Println("\nMenu Kasir")
		fmt.Println("1. Tambah barang")
		fmt.Println("2. Tampilkan rincian belanja")
		fmt.Println("3. Bayar")
		fmt.Println("4. Keluar")

		var pilihan int
		fmt.Print("Pilih menu (1/2/3/4): ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			var namaBarang string
			var harga float64
			var jumlah int

			fmt.Print("Masukkan nama barang: ")
			fmt.Scan(&namaBarang)
			fmt.Print("Masukkan harga barang: ")
			fmt.Scan(&harga)
			fmt.Print("Masukkan jumlah barang: ")
			fmt.Scan(&jumlah)

			kasir.TambahBarang(namaBarang, harga, jumlah)

		case 2:
			kasir.TampilkanRincian()

		case 3:
			kasir.TampilkanRincian()
			var uangDibayar float64
			fmt.Print("Masukkan uang yang dibayar: ")
			fmt.Scan(&uangDibayar)
			kasir.Bayar(uangDibayar)

		case 4:
			fmt.Println("Terima kasih, sampai jumpa!")
			return

		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}