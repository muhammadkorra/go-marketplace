package user

import (
	"go-marketplace/src/abstractions"

	"gorm.io/gorm"
)

type UserDBRepository struct {
	*abstractions.Queryable[User]
}

func NewUserDBRepository(db *gorm.DB) *UserDBRepository {
	return &UserDBRepository{
		Queryable: abstractions.NewQueryable[User](db),
	}
}
