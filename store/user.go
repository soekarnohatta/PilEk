package store

import (
	"PilEk/model"
	"github.com/jinzhu/gorm"
)

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (us *UserStore) GetByUsername(username string) (*model.User, error) {
	var m model.User
	if err := us.db.Where(&model.User{Username: username}).First(&m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (us *UserStore) CreateUser(u *model.User) error {
	return us.db.Create(u).Error
}

func (us *UserStore) UpdateUser(u *model.User) error {
	return us.db.Model(u).Update(u).Error
}

func (us *UserStore) DeleteUser(u *model.User) error {
	return us.db.Model(u).Delete(u).Error
}

func (us *UserStore) List(offset, limit int) ([]model.User, int, error) {
	var (
		Users []model.User
		count int
	)

	us.db.Model(&Users).Count(&count)
	us.db.Offset(offset).
		Limit(limit).
		Order("created_at desc").Find(&Users)

	return Users, count, nil
}
