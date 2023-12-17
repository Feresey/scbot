package reposiory

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

type CorporationRepository struct {
	db *gorm.DB
	tr trace.Tracer
}

func NewCorporationRepository(db *gorm.DB, tr trace.Tracer) *CorporationRepository {
	return &CorporationRepository{
		db: db,
		tr: tr,
	}
}

func (r *CorporationRepository) NewStorage(db *gorm.DB) *CorporationRepository {
	return &CorporationRepository{db: db, tr: r.tr}
}

func (r *CorporationRepository) GetCorp(ctx context.Context, corpID int) (out *Corporation, err error) {
	ctx, span := r.tr.Start(ctx, "corp: GetCorp")
	defer span.End()

	out = &Corporation{}
	if err := r.db.WithContext(ctx).Where("corp_id = ?", corpID).First(out).Error; err != nil {
		return nil, fmt.Errorf("get corp: %w", err)
	}
	return out, nil
}

func (r *CorporationRepository) CreateCorp(ctx context.Context, corp *Corporation) error {
	ctx, span := r.tr.Start(ctx, "corp: CreateCorp")
	defer span.End()

	if err := r.db.WithContext(ctx).Save(&corp).Error; err != nil {
		return fmt.Errorf("create corp: %w", err)
	}
	return nil
}
