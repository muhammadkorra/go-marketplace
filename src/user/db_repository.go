package user

import (
	"gorm.io/gorm"
)

type UserDBRepository struct {
	db *gorm.DB
}

func NewUserDBRepository(db *gorm.DB) *UserDBRepository {
	return &UserDBRepository{
		db: db,
	}
}

func (udb *UserDBRepository) FindAll() ([]User, error) {
	var users []User
	err := udb.db.Model(&User{}).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (udb *UserDBRepository) FindById(id uint) (*User, error) {
	var user User
	err := udb.db.Model(&User{}).First(&user, "id=?", id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
