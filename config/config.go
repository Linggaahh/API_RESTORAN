package config

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Konek() {
	var err error
	DB, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/restoran")
	if err != nil {
		panic(err)
	}

	if err := DB.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Database berhasil terhubung")
}