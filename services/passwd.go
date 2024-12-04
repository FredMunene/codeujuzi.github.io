package services

import (
	"crypto/sha256"
	"fmt"
)

func GetHashedPasswd(passwd string) string {
	hasher := sha256.New()
	hasher.Write([]byte(passwd))
	fmt.Printf("%x", hasher.Sum(nil))

	return string(hasher.Sum(nil))
}
