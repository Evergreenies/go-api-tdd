package security

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/evergreenies/go-api-tdd/pkg/domain"
)

type jwtToken struct {
	key string
}

func NewJWT(key string) (domain.JWT, error) {
	if len(key) < 32 {
		return nil, errors.New("key too short. key must be atleast 32 characters long")
	}

	return &jwtToken{key: key}, nil
}

func (j *jwtToken) CreateToken(user domain.User, duration time.Duration) (*domain.JWTPayload, error) {
	now := time.Now()
	payload := &domain.JWTPayload{
		UserID:    user.ID,
		ExpiresAt: now.Add(duration),
		IssuedAt:  now,
	}
	tokn := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenString, err := tokn.SignedString([]byte(j.key))
	if err != nil {
		return nil, err
	}

	payload.Token = tokenString

	return payload, nil
}

func (j *jwtToken) VerifyToken(tokn string) (*domain.JWTPayload, error) {
	tkn, err := jwt.ParseWithClaims(tokn, &domain.JWTPayload{}, j.keyFunc)
	if err != nil {
		var terr *jwt.ValidationError
		ok := errors.As(err, &terr)
		if ok && errors.Is(terr.Inner, domain.ErrExpiredToken) {
			return nil, domain.ErrExpiredToken
		}

		return nil, domain.ErrInvalidToken
	}

	payload, ok := tkn.Claims.(*domain.JWTPayload)
	if !ok {
		return nil, domain.ErrInvalidToken
	}

	payload.Token = tokn

	return payload, nil
}

func (j *jwtToken) keyFunc(tkn *jwt.Token) (interface{}, error) {
	_, ok := tkn.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		return nil, domain.ErrInvalidToken
	}

	return []byte(j.key), nil
}
