package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTProvider struct {
	signingKey []byte
	issuer     string
	expiry     time.Duration
}

type Claims struct {
	UserID      string   `json:"user_id"`
	UserName    string   `json:"user_name"`
	Permissions []string `json:"permissions"`
	jwt.RegisteredClaims
}

type SignedToken struct {
	token  string
	claims Claims
}

func (s SignedToken) String() string {
	return s.token
}

func (s SignedToken) Claims() Claims {
	return s.claims
}

type ErrInvalidToken struct{}

func (e ErrInvalidToken) Error() string {
	return "invalid token"
}

type ErrInvalidSigningKey struct{}

func (e ErrInvalidSigningKey) Error() string {
	return "invalid signing key"
}

func NewJWTProvider(signingKey string, issuer string, expiry time.Duration) (*JWTProvider, error) {
	if signingKey == "" {
		return nil, ErrInvalidSigningKey{}
	}
	return &JWTProvider{
		signingKey: []byte(signingKey),
		issuer:     issuer,
		expiry:     expiry,
	}, nil
}

func (p *JWTProvider) CreateToken(userID, userName string, permissions []string) (SignedToken, error) {
	claims := Claims{
		UserID:      userID,
		UserName:    userName,
		Permissions: permissions,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    p.issuer,
			Subject:   userID,
			Audience:  jwt.ClaimStrings{p.issuer},
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(p.expiry)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	signed, err := token.SignedString(p.signingKey)
	if err != nil {
		return SignedToken{}, err
	}
	return SignedToken{
		token:  signed,
		claims: claims,
	}, nil
}

func (p *JWTProvider) ParseToken(tokenString string) (*Claims, error) {
	c := Claims{}
	_, err := jwt.ParseWithClaims(tokenString, &c, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return p.signingKey, nil
	})

	if err != nil {
		return nil, err
	}

	return &c, nil
}
