package dao

import (
	"sub-manager/model"

	"gorm.io/gorm"
)

var (
	UserDao UserDAO
)

type UserDAO interface {
	Add(user *model.User) error
	Update(user *model.User) error
	DeleteByName(username string) error
	FindByName(username string) (*model.User, error)
}

type UserDAOImpl struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAOImpl {
	return &UserDAOImpl{db: db}
}

func (dao *UserDAOImpl) Add(user *model.User) error {
	return dao.db.Create(user).Error
}

func (dao *UserDAOImpl) Update(user *model.User) error {
	oldUser, err := dao.FindByName(user.Username)
	if err != nil {
		return err
	}
	return dao.db.Model(oldUser).Updates(user).Error
}

func (dao *UserDAOImpl) DeleteByName(username string) error {
	user, err := dao.FindByName(username)
	if err != nil {
		return err
	}
	return dao.db.Delete(user, 1).Error
}

func (dao *UserDAOImpl) FindByName(username string) (*model.User, error) {
	user := model.User{}
	err := dao.db.First(&user, "username = ?", username).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
