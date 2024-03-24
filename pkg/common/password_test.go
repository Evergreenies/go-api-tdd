package common

import "testing"

func TestPasswordHash(t *testing.T) {
	password := "password"
	hashed, err := PasswordHash(password)
	if err != nil {
		t.Fatal(err)
	}

	if len(hashed) == 0 {
		t.Errorf("want hash, got %q\n", hashed)
	}

	if hashed == password {
		t.Error("password not hashed")
	}
}

func TestPasswordHashError(t *testing.T) {
	longPassword := make([]byte, 73)
	_, err := PasswordHash(string(longPassword))
	if err == nil {
		t.Error("Expected an error; got nil")
	}
}

func TestCheckPassword(t *testing.T) {
	password := "password"
	hashed, err := PasswordHash(password)
	if err != nil {
		t.Fatal(err)
	}

	err = CheckPassword(password, hashed)
	if err != nil {
		t.Errorf("Password verification failed; got %q\n", err)
	}
}
