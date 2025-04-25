package interactor

import (
	"context"
	"errors"

	"github.com/KeihakuOh/career-connect/internal/domain/entity"
	"github.com/KeihakuOh/career-connect/internal/domain/repository"
	"github.com/KeihakuOh/career-connect/internal/infrastructure/auth"
	"github.com/KeihakuOh/career-connect/internal/usecase"
	"golang.org/x/crypto/bcrypt"
)

// SignupInteractor はユーザー登録に特化したユースケースの実装
type SignupInteractor struct {
	userRepo    repository.UserRepository
	studentRepo repository.StudentRepository
	companyRepo repository.CompanyRepository
	jwtManager  *auth.JWTManager
}

// NewSignupInteractor は新しいSignupInteractorを作成する
func NewSignupInteractor(
	userRepo repository.UserRepository,
	studentRepo repository.StudentRepository,
	companyRepo repository.CompanyRepository,
	jwtManager *auth.JWTManager,
) usecase.SignupUseCase {
	return &SignupInteractor{
		userRepo:    userRepo,
		studentRepo: studentRepo,
		companyRepo: companyRepo,
		jwtManager:  jwtManager,
	}
}

// Signup はユーザー登録処理を実行する
func (i *SignupInteractor) Signup(ctx context.Context, input *usecase.SignupInput) (*usecase.SignupOutput, error) {
	// メールアドレスの重複チェック
	exists := false
	user, err := i.userRepo.GetByEmail(ctx, input.Email)
	if err == nil && user != nil {
		exists = true
	}

	if exists {
		return nil, errors.New("email already exists")
	}

	// パスワードをハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// ユーザーエンティティ作成
	user = entity.NewUser(
		input.Email,
		string(hashedPassword),
		input.UserType,
		input.Name,
	)

	// トランザクション処理
	var userID int64
	var message string

	// ユーザー登録
	userID, err = i.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	// ユーザータイプに応じた追加情報登録
	switch input.UserType {
	case "student":
		student := entity.NewStudent(
			userID,
			input.University,
			input.GraduationYear,
			input.Major,
		)
		_, err = i.studentRepo.Create(ctx, student)
		message = "Student registered successfully"
	case "company":
		company := entity.NewCompany(
			userID,
			"", // Description is empty initially
			input.Industry,
			input.Location,
			input.Website,
		)
		_, err = i.companyRepo.Create(ctx, company)
		message = "Company registered successfully"
	default:
		return nil, errors.New("invalid user type")
	}

	if err != nil {
		// ロールバックが必要な場合はここで処理
		// 本来ならばトランザクションマネージャーを使ってロールバック
		_ = i.userRepo.Delete(ctx, userID)
		return nil, err
	}

	// JWTトークン生成
	token, err := i.jwtManager.GenerateToken(userID, input.Email, input.UserType)
	if err != nil {
		return nil, err
	}

	return &usecase.SignupOutput{
		UserID:  userID,
		Message: message,
		Token:   token,
	}, nil
}
