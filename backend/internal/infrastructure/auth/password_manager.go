package auth

import (
	"errors"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

// PasswordManager はパスワード関連の機能を提供する
type PasswordManager struct {
	minLength      int
	requireUpper   bool
	requireLower   bool
	requireDigit   bool
	requireSpecial bool
}

// NewPasswordManager は新しいPasswordManagerインスタンスを生成する
func NewPasswordManager(minLength int, requireUpper, requireLower, requireDigit, requireSpecial bool) *PasswordManager {
	return &PasswordManager{
		minLength:      minLength,
		requireUpper:   requireUpper,
		requireLower:   requireLower,
		requireDigit:   requireDigit,
		requireSpecial: requireSpecial,
	}
}

// ValidatePassword はパスワードが設定された要件を満たしているか検証する
func (pm *PasswordManager) ValidatePassword(password string) error {
	if len(password) < pm.minLength {
		return errors.New("password must be at least " + string(rune(pm.minLength)) + " characters")
	}

	var hasUpper, hasLower, hasDigit, hasSpecial bool
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasDigit = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if pm.requireUpper && !hasUpper {
		return errors.New("password must contain at least one uppercase letter")
	}

	if pm.requireLower && !hasLower {
		return errors.New("password must contain at least one lowercase letter")
	}

	if pm.requireDigit && !hasDigit {
		return errors.New("password must contain at least one digit")
	}

	if pm.requireSpecial && !hasSpecial {
		return errors.New("password must contain at least one special character")
	}

	return nil
}

// HashPassword はパスワードをハッシュ化する
func (pm *PasswordManager) HashPassword(password string) (string, error) {
	if err := pm.ValidatePassword(password); err != nil {
		return "", err
	}

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedBytes), nil
}

// VerifyPassword はパスワードとハッシュが一致するか検証する
func (pm *PasswordManager) VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
