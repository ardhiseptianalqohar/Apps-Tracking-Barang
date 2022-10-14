package caribarang

import (
	"Aplikasi/entities"
	// "bufio"
	"database/sql"
	"fmt"
	// "os"
)

func SearchBarang(db *sql.DB) {
	fmt.Println("======================================")
	fmt.Println("\tTRACKING BARANG ANDA")
	fmt.Println("======================================")

	cari := entities.Resi{}
	fmt.Println("Masukan Nomor Resi Anda :")
	fmt.Scan(&cari.Nomor)
	// in := bufio.NewReader(os.Stdin)
	// line, _ := in.ReadString('\n')
	// cari.Nomor = line

	row, err := Search(db, cari)
	if err != nil {
		fmt.Println("No Resi Tidak Di Temukan", err)
	}
	// fmt.Printf("No Resi : %s\nPengirim : %s\nPenerima : %s\nProduct : %s\nProduct Type : %s\nStatus Barang : %s\nEstimasi Barang : %s\n", row.Nomor, row.Pengirim, row.Penerima, row.Product, row.Product_Type, row.Status_Barang, row.Estimasi)
	fmt.Println("Nomor Resi Anda : ", row.Nomor)
	fmt.Println("Pengirim        : ", row.Pengirim)
	fmt.Println("Penerima        : ", row.Penerima)
	fmt.Println("Product         : ", row.Product)
	fmt.Println("Product Type    : ", row.Product_Type)
	fmt.Println("Status Barang   : ", row.Status_Barang)
	fmt.Println("Estimasi Barang : ", row.Estimasi)
}

func Search(db *sql.DB, search entities.Resi) (entities.Resi, error) {
	stat, err := db.Prepare("SELECT nomor, pengirim, penerima, product, product_type, status_barang, estimasi FROM resi WHERE nomor = ?")
	if err != nil {
		return entities.Resi{}, err
	}

	var info entities.Resi
	errs := stat.QueryRow(search.Nomor).Scan(&info.Nomor, &info.Pengirim, &info.Penerima, &info.Product, &info.Product_Type, &info.Status_Barang, &info.Estimasi)
	if errs != nil {
		return entities.Resi{}, err
	}
	return info, nil
}
