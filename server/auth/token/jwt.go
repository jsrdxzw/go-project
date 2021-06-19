package token

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWTTokenGen struct {
	issue      string
	privateKey *rsa.PrivateKey
	nowFunc    func() time.Time
}

func (t *JWTTokenGen) GenerateToken(accountID string, expire time.Duration) (string, error) {
	now := t.nowFunc().Unix()
	tkn := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.StandardClaims{
		ExpiresAt: now + int64(expire.Seconds()),
		IssuedAt:  now,
		Issuer:    t.issue,
		Subject:   accountID,
	})
	return tkn.SignedString(t.privateKey)
}

func NewJWTTokenGen(issuer string, privateKey *rsa.PrivateKey) *JWTTokenGen {
	return &JWTTokenGen{
		issue:      issuer,
		nowFunc:    time.Now,
		privateKey: privateKey,
	}
}
