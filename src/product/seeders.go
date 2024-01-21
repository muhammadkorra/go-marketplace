package product

import (
	"github.com/go-faker/faker/v4"
	"gorm.io/gorm"
)

type ProductSeeder struct {
	db *gorm.DB
}

func NewProductsSeeder(db *gorm.DB) *ProductSeeder {
	return &ProductSeeder{
		db: db,
	}
}

func (ps *ProductSeeder) GenerateProducts() error {
	for i := 0; i < 10; i++ {
		err := ps.db.Create(&Product{
			Name:        faker.Name(),
			Description: faker.Sentence(),
			Price:       faker.AmountWithCurrency(),
		}).Error

		if err != nil {
			return err
		}
	}

	return nil
}
