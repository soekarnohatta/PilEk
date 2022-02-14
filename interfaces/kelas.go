package interfaces

import "PilEk/model"

type IKelas interface {
	GetByKelasId(uint) (*model.Kelas, error)
	GetByKelasname(string) (*model.Kelas, error)
	CreateKelas(Kelas *model.Kelas) error
	UpdateKelas(Kelas *model.Kelas) error
	DeleteKelas(Kelas *model.Kelas) error
	List(offset, limit int) ([]model.Kelas, int, error)
}
