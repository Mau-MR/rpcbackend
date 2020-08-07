package service

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username      string
	HashedPasword string
	Role          string
	Db            string
}

//returns a new User
func NewUser(username string, password string, role string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("cannot hash password: %w", err)
	}
	user := &User{
		Username:      username,
		HashedPasword: string(hashedPassword),
		Role:          role,
	}
	return user, nil

}
func (user *User) IsCorrectPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.HashedPasword), []byte(password))
	return err == nil
}
func (user *User) Clone() *User {

	return &User{
		Username:      user.Username,
		Role:          user.Role,
		HashedPasword: user.HashedPasword,
	}
}
