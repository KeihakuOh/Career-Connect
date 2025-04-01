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

type StudentRepository interface {
	Create(ctx context.Context, student *entity.Student) (int64, error)
	GetByUserID(ctx context.Context, userID int64) (*entity.Student, error)
	Update(ctx context.Context, student *entity.Student) error
	Delete(ctx context.Context, userID int64) error
}

type CompanyRepository interface {
	Create(ctx context.Context, company *entity.Company) (int64, error)
	GetByUserID(ctx context.Context, userID int64) (*entity.Company, error)
	Update(ctx context.Context, company *entity.Company) error
	Delete(ctx context.Context, userID int64) error
}
