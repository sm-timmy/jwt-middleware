package model

import (
	"golang.org/x/crypto/bcrypt"
	"local/database"
)

//go:generate reform

// User represents a row in users table.
//
//reform:users
type User struct {
	ID       int32  `param:"id" query:"id" form:"id" reform:"id,pk"`
	Name     string `param:"name" query:"name" form:"name" reform:"name"`
	Email    string `param:"email" query:"email" form:"email" reform:"email"`
	Age      *int32 `param:"age" query:"age" form:"age" reform:"age"`
	Password string `param:"password" query:"password" form:"password"  reform:"password"`
}

// CreateUserRecord creates a user record in the database
// CreateUserRecord takes a pointer to a User struct and creates a user record in the database
// It returns an error if there is an issue creating the user record
func (user *User) CreateUserRecord() error {
	if err := database.GlobalDB.Save(user); err != nil {
		return err
	}
	return nil
}

// HashPassword encrypts user password
// HashPassword takes a string as a parameter and encrypts it using bcrypt
// It returns an error if there is an issue encrypting the password
func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// CheckPassword checks user password
// CheckPassword takes a string as a parameter and compares it to the user's encrypted password
// It returns an error if there is an issue comparing the passwords
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
