package model

import "github.com/jinzhu/gorm"

type Kandidat struct {
	gorm.Model
	Nama        string `gorm:"unique_index;not null"`
	NomorUrut   int    `gorm:"column:nomorurut"`
	JumlahSuara int    `gorm:"column:jumlahsuara"`
	Visi        string
	Misi        string
	Status      bool
	Image       string
}
