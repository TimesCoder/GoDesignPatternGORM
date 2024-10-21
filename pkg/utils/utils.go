package utils

import (
	"errors"
	"log"
	"strconv"
	"unicode/utf8"
)

// ValidateBookTitle memeriksa apakah judul buku valid
// Judul buku harus memiliki panjang minimal 3 karakter
func ValidateBookTitle(title string) error {
	if utf8.RuneCountInString(title) < 3 {
		return errors.New("Judul buku harus memiliki minimal 3 karakter")
	}
	return nil
}

// ValidateBookAuthor memeriksa apakah nama penulis valid
// Nama penulis harus memiliki panjang minimal 3 karakter
func ValidateBookAuthor(author string) error {
	if utf8.RuneCountInString(author) < 3 {
		return errors.New("Nama penulis harus memiliki minimal 3 karakter")
	}
	return nil
}

// ValidatePrice memeriksa apakah harga tidak bernilai negatif
func ValidatePrice(price float64) error {
	if price < 0 {
		return errors.New("Harga tidak boleh negatif")
	}
	return nil
}

// LogInfo mencatat pesan informasi ke dalam log
func LogInfo(message string) {
	log.Println("INFO: " + message)
}

// LogWarning mencatat pesan peringatan ke dalam log
func LogWarning(message string) {
	log.Println("WARNING: " + message)
}

// LogError mencatat pesan kesalahan (error) ke dalam log
func LogError(err error) {
	if err != nil {
		log.Println("ERROR: " + err.Error())
	}
}

// ParsePrice mengonversi string menjadi float64 dan memeriksa kesalahan
// Memastikan format harga valid dan tidak negatif
func ParsePrice(priceStr string) (float64, error) {
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return 0, errors.New("Format harga tidak valid")
	}
	if price < 0 {
		return 0, errors.New("Harga tidak boleh negatif")
	}
	return price, nil
}

// ParseID mengonversi string menjadi integer dan memeriksa kesalahan
// Memastikan ID yang diberikan adalah angka positif
func ParseID(idStr string) (int, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		return 0, errors.New("ID tidak valid")
	}
	return id, nil
}
