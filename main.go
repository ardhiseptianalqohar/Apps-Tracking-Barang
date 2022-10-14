package main

import (
	"Aplikasi/config"
	"Aplikasi/controller/CariBarang"
	"Aplikasi/controller/DataTrackingan"
	"Aplikasi/controller/Login"
	"Aplikasi/controller/Register"
	"Aplikasi/controller/read-account"
	"fmt"
)

func main() {
	db := config.ConnectToDB()

	defer db.Close()
	run := true
	for run {
		fmt.Print("MENU:\n1. Registrasi\n2. Login\n3. Tracking Barang\n0. Keluar\n")
		fmt.Println("Masukkan piihan anda:")
		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			{
				register.PendaftaranAkun(db)
			}
		case 2:
			{
				login.Logins(db)

				run2 := true
				for run2 {
					fmt.Print("MENU:\n1. Read Account\n2. Update Trackingan Barang\n3. Masukan Data Resi\n0. Keluar\n")
					fmt.Println("Masukkan piihan anda:")
					var pilihan2 int
					fmt.Scanln(&pilihan2)

					switch pilihan2 {
					case 1:
						{
							readaccount.LihatAkun(db)
						}

					case 2:
						{

						}

					case 3:
						{
							datatrackingan.TambahData(db)
						}

					case 0:
						{
							run2 = false
						}
					}
				}
			}
		case 3:
			{
				caribarang.SearchBarang(db)
			}
		case 0:
			{
				run = false
				fmt.Println("=============================")
				fmt.Println("\tTERIMA KASIH")
				fmt.Println("=============================")
			}
		}
	}
}
