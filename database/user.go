package database

import (
	"jsfraz/trek-server/models"
	"jsfraz/trek-server/utils"
)

// Inserts user to database.
//
//	@param user
//	@return error
func InsertUser(user models.User) error {
	err := utils.GetSingleton().PostgresDb.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

// Check if user with given username exists.
//
//	@param username
//	@return bool
//	@return error
func UserExistsByUsername(username string) (bool, error) {
	var count int64
	err := utils.GetSingleton().PostgresDb.Model(&models.User{}).Where("username = ?", username).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count == 1, nil
}

// Insert new superuser and delete old one.
//
//	@param user
//	@return error
func CreateSuperuser(user models.User) error {
	// transaction
	tx := utils.GetSingleton().PostgresDb.Begin()
	// delete superusers
	if err := tx.Table("users").Where("superuser = ?", true).Delete(&models.User{}).Error; err != nil {
		tx.Rollback() // rollback the transaction if an error occurs
		return err
	}
	// insert user
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback() // rollback the transaction if an error occurs
		return err
	}
	// commit transaction
	return tx.Commit().Error
}

func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := utils.GetSingleton().PostgresDb.Model(&models.User{}).Where("username = ?", username).Attrs(models.User{}).FirstOrInit(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
