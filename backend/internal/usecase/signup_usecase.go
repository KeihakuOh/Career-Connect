package usecase

import (
	"context"
)

// SignupInput はユーザー登録のための入力データ
type SignupInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	UserType string `json:"user_type" validate:"required,oneof=student company"`
	Name     string `json:"name" validate:"required"`
	// 学生用フィールド
	University     string `json:"university,omitempty" validate:"required_if=UserType student"`
	GraduationYear int    `json:"graduation_year,omitempty" validate:"required_if=UserType student"`
	Major          string `json:"major,omitempty" validate:"required_if=UserType student"`
	// 企業用フィールド
	Industry string `json:"industry,omitempty" validate:"required_if=UserType company"`
	Location string `json:"location,omitempty" validate:"required_if=UserType company"`
	Website  string `json:"website,omitempty" validate:"required_if=UserType company"`
}

// SignupOutput はユーザー登録の結果
type SignupOutput struct {
	UserID  int64  `json:"user_id"`
	Message string `json:"message"`
	Token   string `json:"token,omitempty"`
}

// SignupUseCase はユーザー登録のビジネスロジックを定義するインターフェース
type SignupUseCase interface {
	// Signup はユーザー登録処理を実行する
	Signup(ctx context.Context, input *SignupInput) (*SignupOutput, error)
}
