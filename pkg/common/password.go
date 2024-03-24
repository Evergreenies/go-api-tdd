package common

import "golang.org/x/crypto/bcrypt"

func PasswordHash(txt string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(txt), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func CheckPassword(txt string, hashed string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(txt))
}
