package config

import (
	"design_pattern/pkg/models"
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("tokoBuku.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal menghubungkan ke Database!", err)
	}
	// Migrasi schema model Book
	err = DB.AutoMigrate(&models.Book{})
	if err != nil {
		log.Fatal("Gagal melakukan migrasi: ", err)
	}

	fmt.Println("Koneksi ke database berhasil dan migrasi selesai!")
}
