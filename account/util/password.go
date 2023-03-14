package util

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"

	"golang.org/x/crypto/scrypt"
)


func HashPassword(password string) (string, error) {
	
	// 加盐（随机数）
	salt := make([]byte, 32)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	// 哈希（明文密码 plainText + 随机数）
	shash, err := scrypt.Key([]byte(password), salt, 32768, 8, 1, 32)
	if err != nil {
		return "", err
	}
	
	// 双因素
	hashedPW := fmt.Sprintf("%s.%s", hex.EncodeToString(shash), hex.EncodeToString(salt))

	return hashedPW, nil 
}


func comparePasswords(storedPassword string, suppliedPassword string) (bool, error) {
	pwsalt := strings.Split(storedPassword, ".")

	salt, err := hex.DecodeString(pwsalt[1])
	if err != nil {
		return false, fmt.Errorf("Unable to verify user password")
	}

	shash, err := scrypt.Key([]byte(suppliedPassword), salt, 32768, 8, 1, 32)
	
	return hex.EncodeToString(shash) == pwsalt[0], nil 
}