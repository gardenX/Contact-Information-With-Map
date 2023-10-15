package main

import "fmt"

func viewAct() {
	fmt.Println("Selamat Datang Di Fitur Informasi Kontak")
	fmt.Println("Silahkan pilih hal yang ingin anda lakukan")
	fmt.Println("1. Tambah Kontak")
	fmt.Println("2. Cari Kontak")
	fmt.Println("3. Hapus Kontak")
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

func searchContact(person map[string]string) {
	fmt.Print("Masukan Nama : ")
	var namaContact string
	fmt.Scan(&namaContact)

	var value, isExit = person[namaContact]

	if isExit {
		fmt.Println("Nomor Kontak", namaContact, "Adalah", value)
	} else {
		fmt.Println("Nama Kontak ", namaContact, "Tidak Ditemukan ! ")
	}
}

func delContact(person map[string]string) {
	fmt.Print("Masukan Nama : ")
	var namaContact string
	fmt.Scan(&namaContact)

	delete(person, namaContact)

	fmt.Println("Kontak Dengan Nama", namaContact, "Berhasil Dihapus ! ")
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
			fmt.Println("")
			fmt.Println("")
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
				fmt.Println("")
				fmt.Println("")
			case 2:
				searchContact(person)
				fmt.Println("")
				fmt.Println("")
			}
		case 3:
			delContact(person)
			fmt.Println("")
			fmt.Println("")
		}
	}
}
