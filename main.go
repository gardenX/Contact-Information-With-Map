package main

import "fmt"

func viewAct() {
	fmt.Println("Selamat Datang Di Fitur Informasi Kontak")
	fmt.Println("Silahkan pilih hal yang ingin anda lakukan")
	fmt.Println("1. Tambah Kontak")
	fmt.Println("2. Cari Kontak")
}

func addContact(person map[string]string) map[string]string {
	var name string
	fmt.Print("Masukan Nama : ")
	fmt.Scan(&name)

	var contact string
	fmt.Print("Masukan Kontak : ")
	fmt.Scan(&contact)

	person[name] = contact

	return person

}

func viewContact(person map[string]string) {
	for i, m := range person {
		fmt.Println("Nama :", i, " Nomor :", m)
	}

}

func main() {
	person := make(map[string]string)

	for {
		viewAct()

		fmt.Print("Pilih 1/2 : ")
		var action int
		fmt.Scan(&action)

		switch action {
		case 1:
			person = addContact(person)
			viewContact(person)
		case 2:
			fmt.Println("Cari Berdasarkan : ")
			fmt.Println("1. Semua Kotak")
			fmt.Println("2. Nama Kontak")

			fmt.Print("Pilih 1/2 : ")
			var grupView int

			fmt.Scan(&grupView)

			switch grupView {
			case 1:
				viewContact(person)
			case 2:
				fmt.Print("Masukan Nama : ")
				var namaContact string
				fmt.Scan(&namaContact)

				fmt.Println(person[namaContact])
			}

		}
	}
}
