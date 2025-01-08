package repository

import (
	"github.com/iqunlim/easyblog/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(u *model.User) error
	GetByUsername(username string) (*model.User, error)
	GetUserConfig() (*model.UserConfig, error)
	PutUserConfig(userConfig *model.UserConfig) error
}

type UserRepositoryStandard struct {
	DB *gorm.DB
}


func NewUserRepository(DB *gorm.DB) UserRepository {
	return &UserRepositoryStandard{
		DB: DB,
	}
}

func (ur *UserRepositoryStandard) GetByUsername(username string) (*model.User, error) {

	var dbUserInfo model.User

	if err := ur.DB.Where("username = ?", username).First(&dbUserInfo).Error; err != nil {
		return nil, err
	}


	return &dbUserInfo, nil
}

func (ur *UserRepositoryStandard) Create(u *model.User) error {
	if err := ur.DB.Create(u).Error; err != nil {
		return err
	}
	return nil
}

func (ur *UserRepositoryStandard) PutUserConfig(userConfig *model.UserConfig) error {

	if err := ur.DB.Create(userConfig).Error; err != nil {
		return err
	}
	return nil
}

func (ur *UserRepositoryStandard) GetUserConfig() (*model.UserConfig, error) {

	var config *model.UserConfig
	if err := ur.DB.Order("created_at desc").First(&config).Error; err != nil {
		return nil, err
	}
	return config, nil
}