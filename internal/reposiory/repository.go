package reposiory

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

type Repository struct {
	db *gorm.DB
	tr trace.Tracer
}

func New(db *gorm.DB, tp *sdktrace.TracerProvider) *Repository {
	return &Repository{
		db: db,
		tr: tp.Tracer("repository"),
	}
}

func (r *Repository) Players(db *gorm.DB) *PlayerRepository {
	if db == nil {
		db = r.db
	}
	return NewPlayerRepository(db, r.tr)
}

func (r *Repository) Corporations(db *gorm.DB) *CorporationRepository {
	if db == nil {
		db = r.db
	}
	return NewCorporationRepository(db, r.tr)
}

func (r *Repository) AddPlayerInfo(ctx context.Context, playerInfo PlayerInfo) error {
	ctx, span := r.tr.Start(ctx, "rep: AddPlayerInfo")
	defer span.End()

	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		playersRep := r.Players(tx)
		if _, err := playersRep.GetPlayer(ctx, playerInfo.UserID); err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}

			if err := playersRep.CreatePlayer(ctx, &Player{
				UserID: playerInfo.UserID,
			}); err != nil {
				return err
			}
		}

		if playerInfo.CorpID.Valid {
			corpsRep := r.Corporations(tx)
			if _, err := corpsRep.GetCorp(ctx, int(playerInfo.CorpID.Int32)); err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					return err
				}

				if err := corpsRep.CreateCorp(ctx, &Corporation{
					CorpID: int(playerInfo.CorpID.Int32),
				}); err != nil {
					return err
				}
			}
		}

		if err := tx.Save(&playerInfo).Error; err != nil {
			return fmt.Errorf("create player: %w", err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("gorm.Transaction: %w", err)
	}
	return nil
}

func (r *Repository) GetPlayerHistory(ctx context.Context, playerID int) (out *PlayerHistory, err error) {
	ctx, span := r.tr.Start(ctx, "rep: GetPlayerHistory")
	defer span.End()

	var res []*PlayerHistoryItem
	if err := r.db.WithContext(ctx).Debug().
		Model(&PlayerInfo{}).
		Joins("corporation_info").
		Where("id = ?", playerID).
		Find(&res).Error; err != nil {
	}
	return &PlayerHistory{
		History: res,
	}, nil
}
