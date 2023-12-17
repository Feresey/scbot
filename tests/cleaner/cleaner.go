package cleaner

import (
	"context"
	"errors"

	"github.com/Feresey/scbot/internal/reposiory"
	"gorm.io/gorm"
)

type Cleaner struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Cleaner {
	return &Cleaner{db: db}
}

func (c *Cleaner) CleanAll(ctx context.Context) error {
	db := c.db.WithContext(ctx)
	return errors.Join(
		db.Where("1=1").Delete(&reposiory.CorporationInfo{}).Error,
		db.Where("1=1").Delete(&reposiory.PlayerInfo{}).Error,
		db.Where("1=1").Delete(&reposiory.Corporation{}).Error,
		db.Where("1=1").Delete(&reposiory.Player{}).Error,
	)
}
