package model

import (
	"database/sql/driver"
	"errors"
	"html"
	"strings"

	"github.com/iqunlim/easyblog/crypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username" form:"username"`
	Password string `gorm:"size:255;not null;" json:"password" form:"password"`
	Role string `gorm:"size:255;not null;"` 
}

// BeforeSave likely needs to stay with the models
func (u *User) BeforeSave(tx *gorm.DB) error {
	hash, err := crypt.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hash

	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	return nil
}


// https://raaaaaaaay86.medium.com/how-to-store-plain-string-slice-by-using-gorm-f855602013e6
type Tags []string

// This turns the VARCHAR string from the db in to the Tags
func (t *Tags) Scan(src any) error {
	str, ok := src.(string)
	if !ok {
		return errors.New("src value cannot cast to string")
	}
	*t = strings.Split(str, ",")
	return nil
}

// This turns the Tags []string in to a regular string for the driver
// Implements the driver.Valuer interface
func (t Tags) Value() (driver.Value, error) {
	if len(t) == 0 {
		return nil, nil
	}
	return strings.Join(t, ","), nil
}

// ID, CreatedAt and UpdatedAt are all handled by GORM
type BlogPost struct {
	gorm.Model
	Title string `gorm:"type:VARCHAR(255);not null" json:"title" form:"title"`
	ImageURL string `gorm:"type:VARCHAR(255);not null" json:"imgurl" form:"imgurl"`
	Content string `gorm:"type:TEXT;not null" json:"content" form:"content"` 
	Summary string `gorm:"type:TEXT;not null" json:"summary" form:"summary"`
	Tags Tags `gorm:"type:VARCHAR(255);" json:"tags" form:"tags"`
}

/*
func (b *BlogPost) BeforeSave(tx *gorm.DB) error {

	// Escape incoming 
	b.Title = html.EscapeString(b.Title)
	b.Category = html.EscapeString(b.Category)
	b.Content = html.EscapeString(b.Content)
	for _, t := range b.Tags {
		t = html.EscapeString(t)
	}
	return nil
}
	*/