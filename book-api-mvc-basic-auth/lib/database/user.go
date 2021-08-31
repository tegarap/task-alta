package database

import (
	"book-api-mvc/config"
	"book-api-mvc/models"
)

func GetUsers() (interface{}, error) {
	var users []models.User
	if e := config.DB.Find(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}

func GetUser(id int) (interface{}, error) {
	user := models.User{}
	if e := config.DB.First(&user, id).Error; e != nil {
		return nil, e
	}
	return user, nil
}

func CreateUser(user *models.User) (interface{}, error) {
	if e := config.DB.Create(&user).Error; e != nil {
		return nil, e
	}
	return user, nil
}

func DeleteUser(id int) (interface{}, error) {
	user := models.User{}
	if e := config.DB.Delete(&user, id).Error; e != nil {
		return nil, e
	}
	return user, nil
}

func UpdateUser(id int, usr *models.User) (interface{}, error) {
	user := models.User{}
	e := config.DB.First(&user, id)
	if e.Error != nil {
		return nil, e.Error
	}
	if e.RowsAffected > 0 {
		if usr.Name != "" {
			user.Name = usr.Name
		}
		if usr.Email != "" {
			user.Email = usr.Email
		}
		if usr.Password != "" {
			user.Password = usr.Password
		}
		config.DB.Save(&user)
		return &user, nil
	}
	return nil, e.Error
}