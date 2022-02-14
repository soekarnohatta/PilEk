package store

import (
	"PilEk/model"
	"github.com/jinzhu/gorm"
)

type KandidatStore struct {
	db *gorm.DB
}

func NewKandidatStore(db *gorm.DB) *KandidatStore {
	return &KandidatStore{
		db: db,
	}
}

func (ks *KandidatStore) GetByIdKandidat(id int) (*model.Kandidat, error) {
	var m model.Kandidat
	if err := ks.db.Where(&model.Kandidat{Model: gorm.Model{ID: uint(id)}}).First(&m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (ks *KandidatStore) GetByNomorUrut(nourut int) (*model.Kandidat, error) {
	var m model.Kandidat
	if err := ks.db.Where(&model.Kandidat{NomorUrut: nourut}).First(&m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (ks *KandidatStore) CreateKandidat(u *model.Kandidat) (err error) {
	return ks.db.Create(u).Error
}

func (ks *KandidatStore) UpdateKandidat(u *model.Kandidat) error {
	return ks.db.Model(u).Update(u).Error
}
func (ks *KandidatStore) DeleteKandidat(a *model.Kandidat) error {
	return ks.db.Delete(a).Error
}

func (ks *KandidatStore) List(offset, limit int) ([]model.Kandidat, int, error) {
	var tags []model.Kandidat
	var count int
	if err := ks.db.Find(&tags).Count(&count).Offset(offset).Limit(limit).Error; err != nil {
		return nil, 0, err
	}

	return tags, count, nil

}
