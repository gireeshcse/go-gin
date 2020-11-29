package entity

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gireeshcse/go-gin/config"

	"golang.org/x/crypto/argon2"
)

// Address : stores address of user
type Address struct {
	Street  string `json:"street"`
	HouseNo string `json:"house_no"`
	City    string `json:"city" binding:"required"`
	State   string `json:"state" binding:"required"`
	Country string `json:"country" binding:"required"`
}

// User : Defines user entity
type User struct {
	Name       string `json:"name" xml:"title" form:"title"`
	Username   string `json:"username" binding:"min=2,max=10" validate:"containsUser"`
	Email      string `json:"email" xml:"email" form:"email" binding:"required,email" `
	ProfileURL string `json:"profile_url" binding:"required,url"`
	Age        int8   `json:"age" binding:"gte=1,lte=130"`
	Address    string `json:"address" binding:"required"`
	Password   string `json:"password" binding:"required"`
}

// Login : Auth username and password
type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// PasswordConfig the configuration for storing passwords
type PasswordConfig struct {
	time    uint32
	memory  uint32
	threads uint8
	keyLen  uint32
}

// DefaultPasswordConfig : Default password configuration
var DefaultPasswordConfig = &PasswordConfig{
	time:    1,
	memory:  64 * 1024,
	threads: 4,
	keyLen:  32,
}

// GeneratePassword is used to generate a new password hash for storing and
// comparing at a later date.
func GeneratePassword(c *PasswordConfig, password string) (string, error) {

	// Generate a Salt
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, c.time, c.memory, c.threads, c.keyLen)

	// Base64 encode the salt and hashed password.
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	format := "$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s"
	full := fmt.Sprintf(format, argon2.Version, c.memory, c.time, c.threads, b64Salt, b64Hash)
	return full, nil
}

// ComparePassword is used to compare a user-inputted password to a hash to see
// if the password matches or not.
func ComparePassword(password, hash string) (bool, error) {

	parts := strings.Split(hash, "$")

	c := &PasswordConfig{}
	_, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &c.memory, &c.time, &c.threads)
	if err != nil {
		return false, err
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, err
	}

	decodedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, err
	}
	c.keyLen = uint32(len(decodedHash))

	comparisonHash := argon2.IDKey([]byte(password), salt, c.time, c.memory, c.threads, c.keyLen)

	return (subtle.ConstantTimeCompare(decodedHash, comparisonHash) == 1), nil
}

// ValidateToken to validate the json token
func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// signing method validation
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", token.Header["alg"])
		}

		// return the secret signing key
		return []byte(config.SECRET), nil
	})
}
