package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// jika nama function diawali huruf besar (kapital), maka
// function bersifat public
func GetPostgresDB() (*sql.DB, error) {
	// os.Getenv akan mengambil environtment variabel yang dimasukkan
	// untuk mendapatkannya, saat runningnya harus melakukan :
	// POSTGRES_HOST=value go run main.go
	// jadi, environtmen variable nya dimasukkan sebelum me-running main.go
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	databaseName := os.Getenv("POSTGRES_DB_NAME")

	// sprintF berguna untuk menyimpan string dengan sebuah format
	desc := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, databaseName)

	db, err := createConnection(desc)

	if err != nil {
		return nil, err
	}

	return db, nil
}

// jika nama function diawali dengan lowercase, maka function
// bersifat private

// function ini berguna untuk membuat koneksi ke DB
func createConnection(desc string) (*sql.DB, error) {
	// sql.Open('nama DB', settingDB)
	db, err := sql.Open("postgres", desc)

	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	return db, nil
}
