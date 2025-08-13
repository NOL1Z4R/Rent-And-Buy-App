package hash

import "golang.org/x/crypto/bcrypt"

//type password interface {
//	HashPassword(password string) (string, error)
//	VerifyPassword(password, hash string) bool
//}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func VerifyPassword(hash, pw string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
	return err == nil
}
