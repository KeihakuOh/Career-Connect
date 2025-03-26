package repository

import (
	"context"

	"github.com/KeihakuOh/career-connect/internal/domain/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) (int64, error)
	GetByID(ctx context.Context, id int64) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id int64) error
}
