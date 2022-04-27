package security

import (
	"time"
)

// TokenPayload 负载数据
type TokenPayload struct {
	UserID    string
	CreatedAt time.Time
	ExpiresAt time.Time
}

// NewTokenPayload 根据 TokenString 获取一个TokenPayload
func NewTokenPayload(tokenStr string) {
	// token, err := jwt.Parse(tokenStr, parseJwtCallback)
}
