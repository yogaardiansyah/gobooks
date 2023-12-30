// Back End dibuat oleh Yoga Ardiansyah

// config/database.go
package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// DB adalah pool koneksi database
var DB *sql.DB

// ConnectDB membuat koneksi ke database
func ConnectDB() {
	// Menggunakan "root:@/gobooks?parseTime=true" sebagai parameter koneksi.
	// Parameter tersebut bisa disesuaikan dengan kebutuhan konfigurasi database.
	db, err := sql.Open("mysql", "root:@/gobooks?parseTime=true")
	if err != nil {
		// Jika terjadi error saat koneksi, kita menggunakan log.Fatal
		// untuk mengakhiri program dan mencetak pesan error ke log.
		log.Fatal("Error connecting to the database:", err)
	}

	// Menyimpan koneksi database ke variabel DB untuk digunakan di seluruh aplikasi.
	DB = db

	// Mencetak pesan bahwa database terhubung ke log.
	log.Println("Database connected")
}
