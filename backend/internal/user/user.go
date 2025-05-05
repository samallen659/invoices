package user

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

const (
	TIME    = 1
	MEMORY  = 64 * 1024
	THREADS = 4
	KEYLEN  = 32
	SALTLEN = 16
)

type User struct {
	UserName  string `json:"userName"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func NewUser(userName string, firstName string, lastName string, email string, password string) (*User, error) {
	if firstName == "" {
		return nil, errors.New("firstName cannot be empty")
	}
	if lastName == "" {
		return nil, errors.New("lastName cannot be empty")
	}
	if email == "" {
		return nil, errors.New("email cannot be empty")
	}
	if userName == "" {
		return nil, errors.New("userName cannot by empty")
	}

	user := &User{
		UserName:  userName,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}

	password, err := HashPassword(password)
	if err != nil {
		return nil, err
	}
	user.Password = password

	return user, nil
}

func HashPassword(password string) (string, error) {
	salt, err := GenerateSalt()
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, TIME, MEMORY, THREADS, KEYLEN)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash := fmt.Sprintf("$argon2id$v=19$m=%d,t=%d,p=%d$%s$%s", MEMORY, TIME, THREADS, b64Salt, b64Hash)
	return encodedHash, nil
}

func VerifyPassword(password string, encodedHash string) (bool, error) {
	parts := strings.Split(encodedHash, "$")
	if len(parts) != 6 {
		return false, errors.New("Invalid hash format")
	}

	var memory uint32
	var time uint32
	var threads uint8
	_, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &memory, &time, &threads)
	if err != nil {
		return false, err
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, err
	}

	expectedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, err
	}

	hash := argon2.IDKey([]byte(password), salt, time, memory, threads, uint32(len(expectedHash)))

	if subtle.ConstantTimeCompare(hash, expectedHash) == 1 {
		return true, nil
	}

	return false, nil
}

func GenerateSalt() ([]byte, error) {
	salt := make([]byte, SALTLEN)
	_, err := rand.Read(salt)
	return salt, err
}
