package services

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func GetHashedPasswd(passwd string) (string, error) {
	// hasher := sha256.New()
	// hasher.Write([]byte(passwd))

	hash, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %v", err)
	}

	return string(hash), nil
}

func ComparePasswds(dbPasswd, userPasswd string) bool {
	hashedPasswd, err := GetHashedPasswd(userPasswd)
	if err != nil {
		return false
	}
	fmt.Printf("%v,\n%v",dbPasswd,hashedPasswd)
	return (hashedPasswd == dbPasswd)
}
