package utils

import "golang.org/x/crypto/bcrypt"

//EncryptPassword 密码加密
func EncryptPassword(source string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)
	return string(hashPassword), err
}

//ComparePassword 密码对比
func ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
