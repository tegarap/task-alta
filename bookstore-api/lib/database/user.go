package database

import (
	"book-api-mvc/config"
	"book-api-mvc/lib/middleware"
	"book-api-mvc/models"
)

func LoginUser(user *models.User) (interface{}, error) {
	var err error
	if err = config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(user).Error; err != nil {
		return nil, err
	}
	user.Token, err = middleware.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}
	if err = config.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetAllUser() (interface{}, error) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUser(id int) (interface{}, error) {
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func CreateUser(user *models.User) (interface{}, error) {
	if err := config.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func DeleteUser(id int) (interface{}, error) {
	var user models.User
	if err := config.DB.Delete(&user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(user interface{}, newUser *models.User) (interface{}, error) {
	if err := config.DB.Model(user).Updates(newUser).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetDetailUsers(userId int) (interface{}, error) {
	var user models.User
	if err := config.DB.Find(&user, userId).Error; err != nil {
		return nil, err
	}
	return user, nil
}