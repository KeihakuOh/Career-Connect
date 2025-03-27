package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	UserID   int64  `json:"user_id"`
	Email    string `json:"email"`
	UserType string `json:"user_type"`
	jwt.RegisteredClaims
}

type JWTManager struct {
	secretKey string
	expiry    time.Duration
}

func NewJWTManager(secretKey string, expiry time.Duration) *JWTManager {
	return &JWTManager{
		secretKey: secretKey,
		expiry:    expiry,
	}
}

func (m *JWTManager) GenerateToken(userID int64, email, userType string) (string, error) {
	expirationTime := time.Now().Add(m.expiry)
	claims := JWTClaims{
		UserID:   userID,
		Email:    email,
		UserType: userType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	// トークンを作成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 署名して文字列に変換
	tokenString, err := token.SignedString([]byte(m.secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Verify はトークンを検証してクレームを返す
func (m *JWTManager) Verify(tokenString string) (*JWTClaims, error) {
	// トークンをパース
	var claims JWTClaims
	token, err := jwt.ParseWithClaims(
		tokenString,
		&claims,
		func(token *jwt.Token) (interface{}, error) {
			// 署名方法を確認
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}

			return []byte(m.secretKey), nil
		},
	)

	// パースエラーをチェック
	if err != nil {
		return nil, err
	}

	// トークンの有効性をチェック
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return &claims, nil
}
