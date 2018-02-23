package ry

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User represents and application user
type User struct {
	ID        int    `storm:"id,increment"`
	Email     string `storm:"unique"`
	Password  string `json:"-"`
	Name      string
	LastLogin time.Time
	Created   time.Time
}

type UserGroup struct {
	ID      int
	UserID  int
	GroupID int
}

// UserInfo represents user information safe for showing to other users
type UserInfo struct {
	ID   int
	Name string
}

// HashPassword takes a plaintext password string and returns a hash from bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash compares the hash of a plain password with a stored
// hash, returning a bool match result
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
