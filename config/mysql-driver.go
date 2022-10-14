package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDB() *sql.DB {
	dbConnection := os.Getenv("DB_CONNECTION")
	db, err := sql.Open("mysql", dbConnection)

	if err != nil {
		log.Fatal("gagal membuka sql", err.Error())
	}

	errPing := db.Ping()
	if errPing != nil {
		log.Fatal("Gagal menghubungkan DB", errPing.Error())
	} else {
		fmt.Println("===================================")
		fmt.Println("\tSELAMAT DATANG BOS")
		fmt.Println("===================================")
	}

	return db
}
