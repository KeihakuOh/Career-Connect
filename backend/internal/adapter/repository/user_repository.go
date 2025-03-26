// backend/internal/adapter/repository/postgres/user_repository.go
package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/KeihakuOh/career-connect/internal/domain/entity"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// トランザクション処理を行うヘルパー関数
func (r *UserRepository) withTransaction(ctx context.Context, fn func(tx *sql.Tx) (interface{}, error)) (interface{}, error) {
	// トランザクションを開始
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	// 関数実行後に自動的にロールバックまたはコミット
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		}
	}()

	result, err := fn(tx)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *UserRepository) Create(ctx context.Context, user *entity.User) (int64, error) {
	result, err := r.withTransaction(ctx, func(tx *sql.Tx) (interface{}, error) {
		query := `
            INSERT INTO users (
                email, password_hash, user_type, name, profile_image, created_at, updated_at
            ) VALUES ($1, $2, $3, $4, $5, $6, $7)
            RETURNING id
        `

		var id int64
		err := tx.QueryRowContext(
			ctx,
			query,
			user.Email,
			user.PasswordHash,
			user.UserType,
			user.Name,
			user.ProfileImage,
			user.CreatedAt,
			user.UpdatedAt,
		).Scan(&id)

		if err != nil {
			return 0, err
		}

		return id, nil
	})

	if err != nil {
		return 0, err
	}

	return result.(int64), nil
}

func (r *UserRepository) GetByID(ctx context.Context, id int64) (*entity.User, error) {
	result, err := r.withTransaction(ctx, func(tx *sql.Tx) (interface{}, error) {
		query := `
            SELECT id, email, password_hash, user_type, name, profile_image, created_at, updated_at
            FROM users
            WHERE id = $1
        `

		var user entity.User
		err := tx.QueryRowContext(ctx, query, id).Scan(
			&user.ID,
			&user.Email,
			&user.PasswordHash,
			&user.UserType,
			&user.Name,
			&user.ProfileImage,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

		if err != nil {
			if err == sql.ErrNoRows {
				return nil, errors.New("user not found")
			}
			return nil, err
		}

		return &user, nil
	})

	if err != nil {
		return nil, err
	}

	return result.(*entity.User), nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	result, err := r.withTransaction(ctx, func(tx *sql.Tx) (interface{}, error) {
		query := `
            SELECT id, email, password_hash, user_type, name, profile_image, created_at, updated_at
            FROM users
            WHERE email = $1
        `

		var user entity.User
		err := tx.QueryRowContext(ctx, query, email).Scan(
			&user.ID,
			&user.Email,
			&user.PasswordHash,
			&user.UserType,
			&user.Name,
			&user.ProfileImage,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

		if err != nil {
			if err == sql.ErrNoRows {
				return nil, errors.New("user not found")
			}
			return nil, err
		}

		return &user, nil
	})

	if err != nil {
		return nil, err
	}

	return result.(*entity.User), nil
}

func (r *UserRepository) Update(ctx context.Context, user *entity.User) error {
	_, err := r.withTransaction(ctx, func(tx *sql.Tx) (interface{}, error) {
		query := `
            UPDATE users
            SET name = $1, profile_image = $2, updated_at = $3
            WHERE id = $4
        `

		_, err := tx.ExecContext(
			ctx,
			query,
			user.Name,
			user.ProfileImage,
			user.UpdatedAt,
			user.ID,
		)

		return nil, err
	})

	return err
}

func (r *UserRepository) Delete(ctx context.Context, id int64) error {
	_, err := r.withTransaction(ctx, func(tx *sql.Tx) (interface{}, error) {
		query := "DELETE FROM users WHERE id = $1"

		_, err := tx.ExecContext(ctx, query, id)
		return nil, err
	})

	return err
}
