package store

import (
	"PilEk/model"
	"github.com/jinzhu/gorm"
)

type KelasStore struct {
	db *gorm.DB
}

func NewKelasStore(db *gorm.DB) *KelasStore {
	return &KelasStore{
		db: db,
	}
}

func (us *KelasStore) GetByKelasId(KelasId uint) (*model.Kelas, error) {
	var m model.Kelas
	if err := us.db.Where(&model.Kelas{Model: gorm.Model{ID: KelasId}}).First(&m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (us *KelasStore) GetByKelasname(Kelasname string) (*model.Kelas, error) {
	var m model.Kelas
	if err := us.db.Where(&model.Kelas{Kelas: Kelasname}).First(&m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (us *KelasStore) CreateKelas(u *model.Kelas) error {
	return us.db.Create(u).Error
}

func (us *KelasStore) UpdateKelas(u *model.Kelas) error {
	return us.db.Model(u).Update(u).Error
}

func (us *KelasStore) DeleteKelas(u *model.Kelas) error {
	return us.db.Model(u).Delete(u).Error
}

func (us *KelasStore) List(offset, limit int) ([]model.Kelas, int, error) {
	var (
		Kelas []model.Kelas
		count int
	)

	us.db.Model(&Kelas).Count(&count)
	us.db.Offset(offset).
		Limit(limit).
		Order("created_at desc").Find(&Kelas)

	return Kelas, count, nil
}
