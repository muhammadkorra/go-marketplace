package abstractions

import (
	"gorm.io/gorm"
)

type Queryable[T interface{}] struct {
	db    *gorm.DB
	model T
}

func NewQueryable[T interface{}](db *gorm.DB) *Queryable[T] {
	return &Queryable[T]{
		db: db,
	}
}

func (q *Queryable[T]) FindAll() ([]T, error) {
	var results []T
	err := q.db.Model(q.model).Find(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (q *Queryable[T]) FindByID(id uint) (*T, error) {
	var result T
	err := q.db.Model(q.model).First(&result, "id=?", id).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}
