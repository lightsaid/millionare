package security

import "golang.org/x/crypto/bcrypt"

// HashPassword 密码加密
func HashPassword(password string) (string, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashPass), nil
}

// VerifyPassword 验证密码是否正确
func VerifyPassword(hashPass, plainText string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(plainText))
}
