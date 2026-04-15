package repository

import (
	"golang-relational-db/internal/dto"
	"golang-relational-db/internal/model"
)

type UserRepository interface {
	// Mengembalikan pointer (*model.Users) agar:
	// 1. Fungsi bisa nge-return 'nil' jika data user sama sekali tidak ketemu.
	// 2. Lebih hemat memori karena tidak mengkopi seluruh panjang objek saat oper-operan fungsi.
	GetByUsername(username string) (*model.Users, error)
	IsUserExist(username string) bool
	Insert(registerReq dto.RegisterReq) error
	UpdatePassword(username, password string) error
	GetProfile(id int, username string) (int, string, string, error)
}
