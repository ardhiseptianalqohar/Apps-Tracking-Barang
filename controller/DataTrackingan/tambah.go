package datatrackingan

import (
	"Aplikasi/entities"
	// "bufio"
	"database/sql"
	"fmt"
	// "os"
)

func TambahData(db *sql.DB) {
	fmt.Println("=====================================")
	fmt.Println("\tFITUR TAMBAH DATA RESI")
	fmt.Println("=====================================")

	daftar := entities.Resi{}
	fmt.Println("NOMOR RESI:")
	fmt.Scan(&daftar.Nomor)
	fmt.Println("PENGIRIM:")
	fmt.Scan(&daftar.Pengirim)
	fmt.Println("PENERIMA:")
	fmt.Scan(&daftar.Penerima)
	fmt.Println("ALAMAT PENERIMA:")
	fmt.Scan(&daftar.Alamat_Penerima)
	fmt.Println("PRODUCT:")
	fmt.Scan(&daftar.Product)
	fmt.Println("PRODUCT TYPE:")
	fmt.Scan(&daftar.Product_Type)
	fmt.Println("STATUS BARANG:")
	fmt.Scan(&daftar.Status_Barang)
	fmt.Println("ESTIMASI:")
	fmt.Scan(&daftar.Estimasi)
	// in := bufio.NewReader(os.Stdin)
	// line, _ := in.ReadString('\n')
	// daftar.Pengirim = line
	// daftar.Penerima = line
	// daftar.Alamat_Penerima = line
	// daftar.Product = line
	// daftar.Product_Type = line
	// daftar.Status_Barang = line

	var query = "Insert into Resi (Nomor, Pengirim, Penerima, Alamat_penerima, Product, Product_type, Status_barang, Estimasi) values (?, ?, ?, ?, ?, ?, ?, ?)"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		fmt.Println("error prepare insert", errPrepare.Error())
	}
	result, errExec := statement.Exec(daftar.Nomor, daftar.Pengirim, daftar.Penerima, daftar.Alamat_Penerima, daftar.Product, daftar.Product_Type, daftar.Status_Barang, daftar.Estimasi)
	if errExec != nil {
		fmt.Println("error exec insert", errExec.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("==========================================")
			fmt.Println("\tSUKSES MELAKUKAN TAMBAH DATA")
			fmt.Println("==========================================")

		} else {
			fmt.Println("gagal insert")
		}
	}
}
