package model

import (
	"html"
	"strings"

	"github.com/iqunlim/loginexample/crypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"password"`
}

func (u *User) SaveToDB() error {
	var err error 
	err = DBSingleton.Create(&u).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *User) BeforeSave(*gorm.DB) error {
	hash, err := crypt.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hash

	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	return nil
}



// ID, CreatedAt and UpdatedAt are all handled by GORM
type BlogPost struct {
	gorm.Model
	Title string `json:"title"`
	// Location of the markdown file that is associated with the blog post
	ContentLocation string `json:"content"`//TODO: Path? URL? 
	Category string `json:"category"`
	Tags []string `json:"tags"`
}