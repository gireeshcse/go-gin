package service

import (
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWTService interface defines the methods required by jwtservice
type JWTService interface {
	GenerateToken(name string, roles []string) string
	ValidateToken(tokenString string) (*jwt.Token, error)
}

// jwtCustomClaims are custom claims extending the default ones
type jwtCustomClaims struct {
	Name  string `json:"name"`
	Roles string `json:"roles"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey      string
	issuer         string
	expirationTime time.Duration
}

// NewJWTService to create jwt tokens and validate them
func NewJWTService(issuer string, secret string, expirationTime time.Duration) JWTService {
	return &jwtService{
		secretKey:      secret,
		issuer:         issuer,
		expirationTime: expirationTime,
	}
}

func (jwtSrv *jwtService) GenerateToken(username string, roles []string) string {
	var userRoles string = strings.Join(roles, ":")

	// set custom and standard claims
	claims := &jwtCustomClaims{
		username,
		userRoles,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * jwtSrv.expirationTime).Unix(),
			Issuer:    jwtSrv.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	// create a token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(jwtSrv.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (jwtSrv *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// signing method validation
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", token.Header["alg"])
		}

		// return the secret signing key
		return []byte(jwtSrv.secretKey), nil
	})
}
