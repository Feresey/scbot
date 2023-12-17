package tests

import (
	"context"
	"testing"

	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/Feresey/scbot/internal/reposiory"
	"github.com/Feresey/scbot/tests/cleaner"
	"github.com/stretchr/testify/suite"
)

func TestSuite(t *testing.T) {
	suite.Run(t, &Suite{})
}

type Suite struct {
	suite.Suite

	db      *gorm.DB
	cleaner *cleaner.Cleaner

	rep *reposiory.Repository
}

func (s *Suite) SetupSuite() {
	r := s.Require()

	db, err := gorm.Open(
		postgres.Open("postgresql://postgres:pass@localhost:5432/postgres?sslmode=disable"),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: "bot.",
			},
		},
	)
	r.NoError(err)

	exp, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	r.NoError(err)
	bsp := sdktrace.NewBatchSpanProcessor(exp)
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithSpanProcessor(bsp),
	)

	s.db = db
	s.cleaner = cleaner.New(s.db)
	s.rep = reposiory.New(db, tp)
}

func (s *Suite) SetupTest() {
	r := s.Require()
	ctx := context.Background()
	err := s.cleaner.CleanAll(ctx)
	r.NoError(err)
}
