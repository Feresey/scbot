package reposiory

import (
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	_ "gorm.io/plugin/dbresolver"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

// func (r *Repository) GetUserHistory(userID int) (*UserHistory, error) {
// }
