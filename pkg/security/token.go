package security

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

const minSecretKeySize = 32

// token 错误类型
var (
	ErrInvalidSecretKey = fmt.Errorf("密钥长度必须 > %d", minSecretKeySize-1)
	ErrInvalidToken     = fmt.Errorf("Token 无效")
	ErrExpiredToken     = fmt.Errorf("Token 超时")
)

// TokenPayload 负载数据
type TokenPayload struct {
	ID        uuid.UUID              `json:"id"`
	Unique    string                 `json:"unique"` // 唯一标识
	Data      map[string]interface{} // 额外用户自定义数据
	IssuedAt  time.Time              `json:"issued_at"`
	ExpiredAt time.Time              `json:"expired_at"`
}

// NewTokenPayload 根据 TokenString 获取一个TokenPayload
func NewTokenPayload(duration time.Duration, unique string, data map[string]interface{}) (*TokenPayload, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := &TokenPayload{
		ID:        id,
		Unique:    unique,
		Data:      data,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

// GenerateToken 生成token
func (payload *TokenPayload) GenerateToken(secretKey string) (string, error) {
	if len(secretKey) < minSecretKeySize {
		return "", ErrInvalidSecretKey
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(secretKey))
}

// VerifyToken 验证 token 是否有效, 并返回 TokenPayload
func (payload *TokenPayload) VerifyToken(token string, secretKey string) (*TokenPayload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(secretKey), nil
	}

	parsedToken, err := jwt.ParseWithClaims(token, &TokenPayload{}, keyFunc)
	if err != nil {
		return nil, err
	}

	tokenPayload, ok := parsedToken.Claims.(*TokenPayload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return tokenPayload, nil
}

// Valid 验证token是否过时
func (payload *TokenPayload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
