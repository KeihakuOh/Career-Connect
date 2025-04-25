package repository

import (
	"context"

	"github.com/KeihakuOh/career-connect/internal/domain/entity"
)

type StudentRepository interface {
	Create(ctx context.Context, student *entity.Student) (int64, error)
	GetByUserID(ctx context.Context, userID int64) (*entity.Student, error)
	Update(ctx context.Context, student *entity.Student) error
	Delete(ctx context.Context, userID int64) error
}
