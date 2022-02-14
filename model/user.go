package model

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Username   string `gorm:"unique_index;not null"`
	Password   string `gorm:"not null"`
	Nama       string `gorm:"unique_index;not null"`
	IdKelas    int
	IdKandidat int
	Status     int `gorm:"not null"`
	Aktif      bool
	IsAdmin    bool
}

type Kelas struct {
	gorm.Model
	Kelas  string `gorm:"unique_index;not null"`
	Jumlah int
}

const (
	SudahMemilih = 1
	BelumMemilih = 0
)

func (u *User) HashPassword(plain string) (string, error) {
	if len(plain) == 0 {
		return "", errors.New("password should not be empty")
	}
	h, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(h), err
}

func (u *User) CheckPassword(plain string) bool {
	fmt.Println(u.Password, plain)
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plain))
	return err == nil
}
