package repository

import (
	"context"

	"github.com/KeihakuOh/career-connect/internal/domain/entity"
)

type CompanyRepository interface {
	Create(ctx context.Context, company *entity.Company) (int64, error)
	GetByUserID(ctx context.Context, userID int64) (*entity.Company, error)
	Update(ctx context.Context, company *entity.Company) error
	Delete(ctx context.Context, userID int64) error
}
