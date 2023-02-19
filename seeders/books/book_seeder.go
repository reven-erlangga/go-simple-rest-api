package seederBook

import (
	"gorm.io/gorm"
)

type Seeder struct {
	Seeder interface{}
}

func RegisterSeeders(db *gorm.DB, total int) []Seeder {
	seeds := []Seeder{}

	for i := 0; i < total; i++ {
		seeds = append(seeds, Seeder{BookFaker(db)})
	}

	return seeds
}

func DBSeed(db *gorm.DB, total int) error {
	if total == 0 {
		total = 1
	}

	for _, seeder := range RegisterSeeders(db, total) {
		err := db.Debug().Create(seeder.Seeder).Error

		if err != nil {
			return err
		}
	}

	return nil
}