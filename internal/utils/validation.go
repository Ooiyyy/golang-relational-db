package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// FormatValidationError mengubah error validasi Gin yang kaku menjadi pesan yang ramah pengguna
func FormatValidationError(err error) []string {
	var errors []string

	// Pastikan error berasal dari tag struct validator
	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrs {
			switch e.Tag() {
			case "required":
				errors = append(errors, fmt.Sprintf("Kolom %s wajib diisi", e.Field()))
			case "email":
				errors = append(errors, fmt.Sprintf("Kolom %s harus berformat email yang valid", e.Field()))
			case "min":
				errors = append(errors, fmt.Sprintf("Kolom %s minimal %s karakter", e.Field(), e.Param()))
			default:
				errors = append(errors, fmt.Sprintf("Kolom %s tidak valid", e.Field()))
			}
		}
	} else {
		// Error bentuk lain (misal field json salah tipe data)
		errors = append(errors, "Struktur JSON salah atau ada tipe data yang tidak sesuai")
	}

	return errors
}
