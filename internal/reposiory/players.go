package reposiory

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

type PlayerRepository struct {
	db *gorm.DB
	tr trace.Tracer
}

func NewPlayerRepository(db *gorm.DB, tr trace.Tracer) *PlayerRepository {
	return &PlayerRepository{
		db: db,
		tr: tr,
	}
}

func (r *PlayerRepository) NewStorage(db *gorm.DB) *PlayerRepository {
	return &PlayerRepository{db: db, tr: r.tr}
}

func (r *PlayerRepository) GetPlayer(ctx context.Context, playerID int) (out *Player, err error) {
	ctx, span := r.tr.Start(ctx, "GetPlayer")
	defer span.End()

	out = &Player{}
	if err := r.db.WithContext(ctx).Where("user_id = ?", playerID).First(out).Error; err != nil {
		return nil, fmt.Errorf("get player: %w", err)
	}
	return out, nil
}

func (r *PlayerRepository) CreatePlayer(ctx context.Context, player *Player) error {
	ctx, span := r.tr.Start(ctx, "CreatePlayer")
	defer span.End()

	if err := r.db.WithContext(ctx).Save(&player).Error; err != nil {
		return fmt.Errorf("create player: %w", err)
	}
	return nil
}
