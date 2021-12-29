package repos

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type PostgresDetailRepository struct {
	logger *zap.Logger
	db     *gorm.DB
}

func (p *PostgresDetailRepository) FindDetailById(ID uint64) *DetailRecord {
	var detail DetailRecord
	result := p.db.First(&detail, ID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &detail
}

func NewPostgresDetailsRepository(logger *zap.Logger, db *gorm.DB) *PostgresDetailRepository {
	p := &PostgresDetailRepository{
		logger: logger.With(zap.String("type", "PostgresDetailRepository")),
		db:     db,
	}
	return p
}
