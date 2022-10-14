package register

import (
	"Aplikasi/entities"
	"database/sql"
	"fmt"
)

func PendaftaranAkun(db *sql.DB) {
	fmt.Println("======================")
	fmt.Println("FITUR REGISTRASI AKUN")
	fmt.Println("======================")

	daftar := entities.Profile{}
	fmt.Println("Email:")
	fmt.Scanln(&daftar.Email)
	fmt.Println("Name:")
	fmt.Scanln(&daftar.Nama)
	fmt.Println("Address:")
	fmt.Scanln(&daftar.Alamat)
	fmt.Println("Gender:")
	fmt.Scanln(&daftar.Gender)
	fmt.Println("Nomor:")
	fmt.Scanln(&daftar.Nomor)
	fmt.Println("Password:")
	fmt.Scanln(&daftar.Password)
	fmt.Println("Status:")
	fmt.Scanln(&daftar.Status)

	var query = "Insert into Profile (Email, Nama, Alamat, Gender, Nomor, Password, Status) values (?, ?, ?, ?, ?, ?, ?)"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		fmt.Println("error prepare insert", errPrepare.Error())
	}
	result, errExec := statement.Exec(daftar.Email, daftar.Nama, daftar.Alamat, daftar.Gender, daftar.Nomor, daftar.Password, daftar.Status)
	if errExec != nil {
		fmt.Println("error exec insert", errExec.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("========================================================")
			fmt.Println("\tANDA SUKSES MELAKUKAN PENDAFTARAN")
			fmt.Println("\tSELAMAT MENIKMATI FITUR DAN LAYANAN KAMI")
			fmt.Println("========================================================")

		} else {
			fmt.Println("gagal insert")
		}
	}
}
