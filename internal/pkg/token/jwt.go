package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var ErrInvalidToken = errors.New("invalid token")

type Manager struct {
	secret []byte
	expire time.Duration
}

func New(secret string, expireMinutes int) *Manager {
	return &Manager{
		secret: []byte(secret),
		expire: time.Duration(expireMinutes) * time.Minute,
	}
}

type Claims struct {
	UserID uint64 `json:"user_id`
	jwt.RegisteredClaims
}

func (m *Manager) Sign(userID uint64) (string, error) {
	now := time.Now()
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(m.expire)),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return t.SignedString(m.secret)
}

func (m *Manager) Parse(tokenStr string) (uint64, error) {
	t, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (any, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, ErrInvalidToken
		}
		return m.secret, nil
	})
	if err != nil {
		return 0, ErrInvalidToken
	}
	claims, ok := t.Claims.(*Claims)
	if !ok || !t.Valid {
		return 0, ErrInvalidToken
	}
	return claims.UserID, nil
}
