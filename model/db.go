package model

import (
	"log"

	"github.com/iqunlim/loginexample/crypt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


var DBSingleton *gorm.DB


func CreateDB() (*gorm.DB, error) {
	DB, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	DB.AutoMigrate(&User{})
	log.Print("Connected to sqlite database")
	return DB, nil
}

func VerifyLogin(u *User) (*User, error) {
	var err error 

	UserInformation := User{}

	if err = DBSingleton.Model(User{}).Where("username = ?", u.Username).Take(&UserInformation).Error; err != nil {
		return nil, err
	}

	if !crypt.CheckPasswordHash(u.Password, UserInformation.Password) {
		return nil, bcrypt.ErrMismatchedHashAndPassword
	}

	return &UserInformation, nil
}