package security

import (
	"testing"
	"time"

	"github.com/evergreenies/go-api-tdd/pkg/common"
	"github.com/evergreenies/go-api-tdd/pkg/domain"
)

func TestJWTToken(t *testing.T) {
	_, err := NewJWT("short-key")
	if err == nil {
		t.Error("expecting error as key too short.")
	}
}

func TestJwtToken(t *testing.T) {
	key := common.RandomString(32)
	newJWT, err := NewJWT(key)
	if err != nil {
		t.Fatal(err)
	}

	user := domain.User{
		ID:    1,
		Email: "test@test.com",
	}

	payload, err := newJWT.CreateToken(user, 1*time.Minute)
	if err != nil {
		t.Fatal(err)
	}

	if payload.UserID != user.ID {
		t.Errorf("want %q; got %q", payload.UserID, user.ID)
	}

	if len(payload.Token) == 0 {
		t.Error("token string is empty")
	}

	if payload.ExpiresAt.Before(time.Now()) {
		t.Error("token is expired")
	}

	_, err = newJWT.VerifyToken(payload.Token + "invalid")
	if err == nil {
		t.Error("invalid token")
	}

	expiredToken, err := newJWT.CreateToken(user, -1*time.Minute)
	if err != nil {
		t.Fatal("some error while creating token")
	}

	_, err = newJWT.VerifyToken(expiredToken.Token)
	if err == nil {
		t.Error("invalid token or expired")
	}
}
