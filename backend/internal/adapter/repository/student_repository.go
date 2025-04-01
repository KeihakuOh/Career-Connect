// backend/internal/adapter/repository/postgres/student_repository.go
package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/KeihakuOh/career-connect/internal/domain/entity"
)

type StudentRepository struct {
	db *sql.DB
}

func NewStudentRepository(db *sql.DB) *StudentRepository {
	return &StudentRepository{
		db: db,
	}
}

// Create は新しい学生プロフィールを作成する
func (r *StudentRepository) Create(ctx context.Context, student *entity.Student) (int64, error) {
	result, err := r.withTransaction(ctx, func(tx *sql.Tx) (interface{}, error) {
		query := `
            INSERT INTO students (
                user_id, university, graduation_year, major, created_at, updated_at
            ) VALUES ($1, $2, $3, $4, $5, $6)
            RETURNING id
        `

		var id int64
		err := tx.QueryRowContext(
			ctx,
			query,
			student.UserID,
			student.University,
			student.GraduationYear,
			student.Major,
			student.CreatedAt,
			student.UpdatedAt,
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

// GetByUserID はユーザーIDに基づいて学生プロフィールを取得する
func (r *StudentRepository) GetByUserID(ctx context.Context, userID int64) (*entity.Student, error) {
	result, err := r.withTransaction(ctx, func(tx *sql.Tx) (interface{}, error) {
		query := `
            SELECT id, user_id, university, graduation_year, major, created_at, updated_at
            FROM students
            WHERE user_id = $1
        `

		var student entity.Student
		err := tx.QueryRowContext(ctx, query, userID).Scan(
			&student.ID,
			&student.UserID,
			&student.University,
			&student.GraduationYear,
			&student.Major,
			&student.CreatedAt,
			&student.UpdatedAt,
		)

		if err != nil {
			if err == sql.ErrNoRows {
				return nil, errors.New("student profile not found")
			}
			return nil, err
		}

		return &student, nil
	})

	if err != nil {
		return nil, err
	}

	return result.(*entity.Student), nil
}

// Update は学生プロフィールを更新する
func (r *StudentRepository) Update(ctx context.Context, student *entity.Student) error {
	_, err := r.withTransaction(ctx, func(tx *sql.Tx) (interface{}, error) {
		query := `
            UPDATE students
            SET university = $1, graduation_year = $2, major = $3, updated_at = $4
            WHERE user_id = $5
        `

		_, err := tx.ExecContext(
			ctx,
			query,
			student.University,
			student.GraduationYear,
			student.Major,
			student.UpdatedAt,
			student.UserID,
		)

		return nil, err
	})

	return err
}

// Delete は学生プロフィールを削除する
func (r *StudentRepository) Delete(ctx context.Context, userID int64) error {
	_, err := r.withTransaction(ctx, func(tx *sql.Tx) (interface{}, error) {
		query := "DELETE FROM students WHERE user_id = $1"

		_, err := tx.ExecContext(ctx, query, userID)
		return nil, err
	})

	return err
}

// トランザクション処理を行うヘルパー関数
func (r *StudentRepository) withTransaction(ctx context.Context, fn func(tx *sql.Tx) (interface{}, error)) (interface{}, error) {
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
