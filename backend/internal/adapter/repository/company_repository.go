// backend/internal/adapter/repository/postgres/company_repository.go
package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/KeihakuOh/career-connect/internal/domain/entity"
)

type CompanyRepository struct {
	db *sql.DB
}

func NewCompanyRepository(db *sql.DB) *CompanyRepository {
	return &CompanyRepository{
		db: db,
	}
}

// Create は新しい企業プロフィールを作成する
func (r *CompanyRepository) Create(ctx context.Context, company *entity.Company) (int64, error) {
	result, err := r.withTransaction(ctx, func(tx *sql.Tx) (interface{}, error) {
		query := `
            INSERT INTO companies (
                user_id, description, industry, location, website, created_at, updated_at
            ) VALUES ($1, $2, $3, $4, $5, $6, $7)
            RETURNING id
        `

		var id int64
		err := tx.QueryRowContext(
			ctx,
			query,
			company.UserID,
			company.Description,
			company.Industry,
			company.Location,
			company.Website,
			company.CreatedAt,
			company.UpdatedAt,
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

// GetByUserID はユーザーIDに基づいて企業プロフィールを取得する
func (r *CompanyRepository) GetByUserID(ctx context.Context, userID int64) (*entity.Company, error) {
	result, err := r.withTransaction(ctx, func(tx *sql.Tx) (interface{}, error) {
		query := `
            SELECT id, user_id, description, industry, location, website, created_at, updated_at
            FROM companies
            WHERE user_id = $1
        `

		var company entity.Company
		err := tx.QueryRowContext(ctx, query, userID).Scan(
			&company.ID,
			&company.UserID,
			&company.Description,
			&company.Industry,
			&company.Location,
			&company.Website,
			&company.CreatedAt,
			&company.UpdatedAt,
		)

		if err != nil {
			if err == sql.ErrNoRows {
				return nil, errors.New("company profile not found")
			}
			return nil, err
		}

		return &company, nil
	})

	if err != nil {
		return nil, err
	}

	return result.(*entity.Company), nil
}

// Update は企業プロフィールを更新する
func (r *CompanyRepository) Update(ctx context.Context, company *entity.Company) error {
	_, err := r.withTransaction(ctx, func(tx *sql.Tx) (interface{}, error) {
		query := `
            UPDATE companies
            SET description = $1, industry = $2, location = $3, website = $4, updated_at = $5
            WHERE user_id = $6
        `

		_, err := tx.ExecContext(
			ctx,
			query,
			company.Description,
			company.Industry,
			company.Location,
			company.Website,
			company.UpdatedAt,
			company.UserID,
		)

		return nil, err
	})

	return err
}

// Delete は企業プロフィールを削除する
func (r *CompanyRepository) Delete(ctx context.Context, userID int64) error {
	_, err := r.withTransaction(ctx, func(tx *sql.Tx) (interface{}, error) {
		query := "DELETE FROM companies WHERE user_id = $1"

		_, err := tx.ExecContext(ctx, query, userID)
		return nil, err
	})

	return err
}

// トランザクション処理を行うヘルパー関数
func (r *CompanyRepository) withTransaction(ctx context.Context, fn func(tx *sql.Tx) (interface{}, error)) (interface{}, error) {
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
