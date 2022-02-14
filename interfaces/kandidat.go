package interfaces

import "PilEk/model"

type IKandidat interface {
	GetByIdKandidat(int) (*model.Kandidat, error)
	GetByNomorUrut(int) (*model.Kandidat, error)
	CreateKandidat(kandidat *model.Kandidat) error
	UpdateKandidat(kandidat *model.Kandidat) error
	DeleteKandidat(kandidat *model.Kandidat) error
	List(offset, limit int) ([]model.Kandidat, int, error)
}
