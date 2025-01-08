package service

import (
	"bufio"
	"errors"
	"flag"
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
	// Extremely basic admin handling.
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


	var userToBe *model.User
	name := flag.String("username", "", "Intitial Admin Username")
	pwd := flag.String("pwd", "", "Initial Admin Password")
	flag.Parse()
	if *name == "" || *pwd == "" {
		fmt.Println("No or wrong flags detected. Entering manual entry mode")
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Admin Username: ")
		newName, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		newName = strings.Trim(newName, " \t\r\n")
		fmt.Print("Admin Password: ")
		newPwd, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		newPwd = strings.Trim(newPwd, " \t\r\n")
		userToBe = &model.User{
			Username: newName,
			Password: newPwd,
		}
	} else {
		userToBe = &model.User{
			Username: *name,
			Password: *pwd,
		}
	}
	if userToBe.Username == "" || userToBe.Password == "" {
		panic("Username and password must be input")

	}
	if err := ur.Register(userToBe); err != nil {
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