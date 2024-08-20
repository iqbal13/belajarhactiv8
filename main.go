package main

import (
	"fmt"
	"os"
)

type Participant struct {
	Nama, Alamat, Pekerjaan, Alasan string
}

func main() {

	participantList := []Participant{
		{"Septian", "Jl. Cakung", "Backend Developer", "Mempelajari Hal Baru"},
		{"David", "Jl. Grogol", "Database Administrator", "Mempelajari Golang karna disuruhs"},
		{"Tiansyah", "Jl. Palembang", "Backend Developer", "Mempelajari Bahasa Baru"},
		{"Thomas", "Jl. Sepanjang Jalan Kenangan", "Fullstack Developer", "Memperdalam Backend Developer dengan Golang"},
		{"Iqbal", "Jl. ABC", "Fullstack Developer", "Migrasi dari PHP ke Golang"},
	}

	if len(os.Args) < 2 {
		fmt.Println("Silahkan Masukan Nama sebagai Argumen")
		return
	}

	name := os.Args[1]

	for _, participant := range participantList {
		if participant.Nama == name {
			fmt.Printf("Nama: %s\nAlamat: %s\nPekerjaan: %s\nAlasan: %s\n",
				participant.Nama, participant.Alamat, participant.Pekerjaan, participant.Alasan)
			return
		}
	}

	fmt.Printf("Data untuk Nama %s Tidak Ditemukan.\n", name)
}
