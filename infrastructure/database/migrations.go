package database

import (
	"book-api-cleanarc/internal/domain"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "20251215_create_authors_and_books",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.AutoMigrate(&domain.Author{}); err != nil {
					return err
				}
				if err := tx.AutoMigrate(&domain.Book{}); err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Migrator().DropTable(&domain.Book{}, &domain.Author{}); err != nil {
					return err
				}
				return nil
			},
		},
	})

	return m.Migrate()
}
