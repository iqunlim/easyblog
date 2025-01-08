package service

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/iqunlim/easyblog/crypt"
	"github.com/iqunlim/easyblog/model"
	"github.com/iqunlim/easyblog/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(u *model.User) error
	Verify(u *model.User) (*model.User, error)
	// Delete
	FirstRun()
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

func (ur *UserServiceImpl) FirstRun() {
	fmt.Println(" _____                _     _             ")
	fmt.Println("| ____|__ _ ___ _   _| |__ | | ___   __ _ ")
	fmt.Println("|  _| / _` / __| | | | '_ \\| |/ _ \\ / _` |")
	fmt.Println("| |__| (_| \\__ \\ |_| | |_) | | (_) | (_| |")
	fmt.Println("|_____\\__,_|___/\\__, |_.__/|_|\\___/ \\__, |")
	fmt.Println("                 |___/               |___/ ")

	_, err := ur.repository.GetUserConfig()
	if err == nil {
		return
	}
	// Check user repository for database connection
	// Create settings table
	// Bufio reader that goes through some config options. Figure them out!
	// Admin username
	// Admin password
	// Apply this to the users table with userservice.Register()
	// JWT key?
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Admin Username: ")
	name, _ := reader.ReadString('\n')
	name = strings.Trim(name, "\n")
	fmt.Print("Admin Password: ")
	pwd, _ := reader.ReadString('\n')
	pwd = strings.Trim(pwd, "\n")
	fmt.Printf("Username: %s, Password: %s", name, pwd)
	registerUser := &model.User{
		Username: name,
		Password: name,
	}
	if err := ur.Register(registerUser); err != nil {
		log.Fatalf("Error occured in registration: %v", err)
	}
	if err := ur.repository.PutUserConfig(&model.UserConfig{ FirstRunCompleted: true }); err != nil {
		log.Fatalf("Error occured in setting up config: %v", err)
	}
}


type UserExistsError struct {
}

func (e UserExistsError) Error() string {
	return "Username Exists"
}