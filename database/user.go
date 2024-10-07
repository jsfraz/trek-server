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
	if err := tx.Where("superuser = ?", true).Delete(&models.User{}).Error; err != nil {
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

// Get user by username.
//
//	@param username
//	@return *models.User
//	@return error
func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := utils.GetSingleton().PostgresDb.Model(&models.User{}).Where("username = ?", username).Attrs(models.User{}).FirstOrInit(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Check if user with given ID exists.
//
//	@param id
//	@return bool
//	@return error
func UserExistsById(id uint64) (bool, error) {
	var count int64
	err := utils.GetSingleton().PostgresDb.Model(&models.User{}).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count == 1, nil
}

// Get user by ID.
//
//	@param id
//	@return *models.User
//	@return error
func GetUserById(id uint64) (*models.User, error) {
	var user models.User
	err := utils.GetSingleton().PostgresDb.Model(&models.User{}).Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Get all users.
//
//	@return *[]models.User
//	@return error
func GetAllUsers() (*[]models.User, error) {
	var users []models.User = []models.User{}
	err := utils.GetSingleton().PostgresDb.Model(&models.User{}).Order("id ASC").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}

// Delete user with given ID.
//
//	@param id
//	@return error
func DeleteUser(id uint64) error {
	return utils.GetSingleton().PostgresDb.Where("id = ?", id).Delete(&models.User{}).Error
}

// Update user.
//
//	@param id
//	@param username
//	@param password
//	@return error
func UpdateUser(id uint64, username string, password string) error {
	var user models.User
	err := utils.GetSingleton().PostgresDb.Model(&models.User{}).Where("id = ?", id).First(&user).Error
	if err != nil {
		return err
	}
	// username
	user.Username = username
	// password
	if password != "" {
		err := user.SetPassword(password)
		// error
		if err != nil {
			return err
		}
	}
	// update
	return utils.GetSingleton().PostgresDb.Save(&user).Error
}
