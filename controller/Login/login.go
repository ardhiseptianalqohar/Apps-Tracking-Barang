package login

import (
	"Aplikasi/entities"
	"database/sql"
	"fmt"
)

func Logins(db *sql.DB) {
	masuk := entities.Profile{}
	fmt.Println("Silahkan Masukkan Alamat Email Anda :")
	fmt.Scanln(&masuk.Email)
	fmt.Println("Silahkan Masukkan Password Anda :")
	fmt.Scanln(&masuk.Password)

	row, err := LoginUser(db, masuk)
	if err != nil {
		fmt.Println("Gagal Login", err)
	}

	fmt.Println("====================================")
	fmt.Println("\tSelamat Datang", row.Nama)
	fmt.Println("====================================")

	fmt.Println(
		"\nNomor Anda     :", row.Nomor,
		"\nAlamat Anda    :", row.Alamat,
		"\nJenis Kelamin  :", row.Gender,
		"\nStatus Akun    :", row.Status)
	fmt.Println("==================================")
}

func LoginUser(db *sql.DB, profile entities.Profile) (entities.Profile, error) {
	statm, err := db.Prepare("select nama, nomor, alamat, gender, status from profile where email = ? and password = ?")
	if err != nil {
		// fmt.Println("Statment Error", err)
		return entities.Profile{}, err
	}

	var row entities.Profile
	errs := statm.QueryRow(profile.Email, profile.Password).Scan(&row.Nama, &row.Nomor, &row.Alamat, &row.Gender, &row.Status)
	if errs != nil {
		// fmt.Println("Error", errs)
		return entities.Profile{}, errs

	}

	return row, nil
}
