package utils

import "golang.org/x/crypto/bcrypt"

func HashingPassword(pass string) (string, error) {
	hashedByte, errHashed := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if errHashed != nil {
		return "", errHashed
	}

	return string(hashedByte), nil
}

func CheckPasswordHash(hashedPassword, pass string) bool {
	errCheckPassword := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(pass))
	return errCheckPassword == nil
}
