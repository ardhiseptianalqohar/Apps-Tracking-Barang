package readaccount

import (
	"Aplikasi/entities"
	// "bufio"
	"database/sql"
	"fmt"
	// "os"
)

func LihatAkun(db *sql.DB) {
	fmt.Println("======================================")
	fmt.Println("\tLIHAT PROFILE ANDA")
	fmt.Println("======================================")

	lihat := entities.Profile{}
	fmt.Println("Masukan Email Anda :")
	fmt.Scan(&lihat.Email)
	// in := bufio.NewReader(os.Stdin)
	// line, _ := in.ReadString('\n')
	// lihat.Email = line

	row, err := Cari(db, lihat)
	if err != nil {
		fmt.Println("Nama Anda Tidak Di Temukan", err)
	}

	fmt.Println("Email Anda      : ", row.Email)
	fmt.Println("Nama Anda       : ", row.Nama)
	fmt.Println("Alamat Anda     : ", row.Alamat)
	fmt.Println("Jenis Kelamin   : ", row.Gender)
	fmt.Println("Nomor Telp      : ", row.Nomor)
	fmt.Println("Password        : ", row.Password)
	fmt.Println("Status Akun     : ", row.Status)
}

func Cari(db *sql.DB, cari entities.Profile) (entities.Profile, error) {
	stat, err := db.Prepare("SELECT email, nama, alamat, gender, nomor, password, status FROM profile where email = ?")
	if err != nil {
		return entities.Profile{}, err
	}
	var info entities.Profile
	errs := stat.QueryRow(cari.Email).Scan(&info.Email, &info.Nama, &info.Alamat, &info.Gender, &info.Nomor, &info.Password, &info.Status)
	if errs != nil {
		return entities.Profile{}, err
	}
	return info, nil
}
