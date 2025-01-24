package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Parkir struct {
	kapasitas       int
	kendaraanParkir []string
}

func NewParkir(kapasitas int) *Parkir {
	return &Parkir{
		kapasitas:       kapasitas,
		kendaraanParkir: []string{},
	}
}

func (p *Parkir) MasukParkir() {
	if len(p.kendaraanParkir) < p.kapasitas {
		fmt.Print("Masukkan nama kendaraan yang ingin parkir: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		kendaraan := strings.TrimSpace(scanner.Text())
		p.kendaraanParkir = append(p.kendaraanParkir, kendaraan)
		fmt.Printf("%s berhasil masuk ke area parkir.\n", kendaraan)
	} else {
		fmt.Println("Parkir sudah penuh!")
	}
}

func (p *Parkir) KeluarParkir() {
	fmt.Print("Masukkan nama kendaraan yang ingin keluar: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	kendaraan := strings.TrimSpace(scanner.Text())

	for i, v := range p.kendaraanParkir {
		if v == kendaraan {
			p.kendaraanParkir = append(p.kendaraanParkir[:i], p.kendaraanParkir[i+1:]...)
			fmt.Printf("%s telah keluar dari area parkir.\n", kendaraan)
			return
		}
	}
	fmt.Printf("%s tidak ditemukan di parkir.\n", kendaraan)
}

func (p *Parkir) StatusParkir() {
	if len(p.kendaraanParkir) == 0 {
		fmt.Println("Area parkir kosong.")
	} else {
		fmt.Println("Kendaraan yang terparkir:")
		for _, kendaraan := range p.kendaraanParkir {
			fmt.Println(kendaraan)
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukkan kapasitas parkir: ")
	var kapasitasParkir int
	fmt.Scanln(&kapasitasParkir)

	parkir := NewParkir(kapasitasParkir)

	for {
		fmt.Println("\nPilih menu:")
		fmt.Println("1. Masukkan kendaraan")
		fmt.Println("2. Keluarkan kendaraan")
		fmt.Println("3. Lihat status parkir")
		fmt.Println("4. Keluar")

		fmt.Print("Pilihan Anda: ")
		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			parkir.MasukParkir()
		case 2:
			parkir.KeluarParkir()
		case 3:
			parkir.StatusParkir()
		case 4:
			fmt.Println("Terima kasih telah menggunakan sistem parkir!")
			return
		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}