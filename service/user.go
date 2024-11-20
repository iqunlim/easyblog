package service

import (
	"errors"

	"github.com/iqunlim/easyblog/crypt"
	"github.com/iqunlim/easyblog/model"
	"github.com/iqunlim/easyblog/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(u *model.User) error
	Verify(u *model.User) (*model.User, error)
	// Delete
}

type UserServiceImpl struct {
	repository repository.UserRepository
}

func NewUserService(u repository.UserRepository) UserService {
	return &UserServiceImpl{
		repository: u,
	}
}

func (ur *UserServiceImpl) Verify(u *model.User) (*model.User, error) {


	dbUserInfo, err := ur.repository.GetByUsername(u.Username)
	if err != nil {
		return nil, err
	}

	if u.Username != dbUserInfo.Username {
		return nil, errors.New("Username and database user do not match! This is a terrible bug!")
	}

	if !crypt.CheckPasswordHash(u.Password, dbUserInfo.Password) {
		return nil, bcrypt.ErrMismatchedHashAndPassword
	}

	return dbUserInfo, nil
}

func (ur *UserServiceImpl) Register(u *model.User) error {

	u.Role = "admin" 

	usr, _ := ur.repository.GetByUsername(u.Username)
	if usr != nil {
		return &UserExistsError{}
	}

	if err := ur.repository.Create(u); err != nil {
		return err

	}
	return nil
}


type UserExistsError struct {
}

func (e UserExistsError) Error() string {
	return "Username Exists"
}