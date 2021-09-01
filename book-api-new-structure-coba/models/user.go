package models

import (
	"book-api-mvc/api/middlewares"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Token    string `json:"token" form:"token"`
}

type GormUserModel struct {
	db *gorm.DB
}

func NewUserModel(db *gorm.DB) *GormUserModel {
	return &GormUserModel{db: db}
}

// Interface User

type UserModel interface {
	GetAll() ([]User, error)
	Get(userId int) (User, error)
	Insert(User) (User, error)
	Delete(userId int) (User, error)
	Edit(user User, userId int) (User, error)
	Login(email, password string) (User, error)
}

func (m *GormUserModel) GetAll() ([]User, error) {
	var users []User
	if err := m.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (m *GormUserModel) Get(userId int) (User, error) {
	var user User
	if err := m.db.First(&user, userId).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (m *GormUserModel) Insert(user User) (User, error) {
	if err := m.db.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (m *GormUserModel) Delete(userId int) (User, error) {
	var user User
	if err := m.db.First(&user, userId).Error; err != nil {
		return user, err
	}
	if err := m.db.Delete(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (m *GormUserModel) Edit(userTemp User, userId int) (User, error) {
	var user User
	if err := m.db.First(&user, userId).Error; err != nil {
		return user, err
	}
	if userTemp.Name != "" {
		user.Name = userTemp.Name
	}
	if userTemp.Email != "" {
		user.Email = userTemp.Email
	}
	if userTemp.Password != "" {
		user.Password = userTemp.Password
	}
	if err := m.db.Save(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (m *GormUserModel) Login(email, password string) (User, error) {
	var user User
	var err error
	if err = m.db.Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
		return user, err
	}

	user.Token, err = middlewares.CreateToken(int(user.ID))

	if err != nil {
		return user, err
	}
	if err := m.db.Save(user).Error; err != nil {
		return user, err
	}

	return user, nil
}
