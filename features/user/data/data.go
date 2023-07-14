package data

import (
	"belajar/bareng/features"
	"belajar/bareng/features/user"
	"belajar/bareng/helper"
	"errors"

	"gorm.io/gorm"
)

type userData struct {
	db *gorm.DB
}

// Upgrade implements user.UserData.
func (data *userData) Upgrade(userId uint, user features.UserEntity) (uint, error) {
	var userModel features.User
	tx:=data.db.First(&userModel,userId)
	if tx.Error != nil{
		return 0,tx.Error
	}

	txx:= data.db.Model(&userModel).Where("id=?",userId).Updates(features.UserEntityToModel(user))
	if txx.Error != nil{
		return 0,tx.Error
	}
	return userId,nil
}

// Delete implements user.UserData
func (repo *userData) Delete(id uint) error {
	var user features.User

	tx := repo.db.Delete(&user, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Update implements user.UserData
func (repo *userData) Update(id uint, user features.UserEntity) (uint, error) {
	var userModel features.User
	tx := repo.db.Model(&userModel).Where("id=?", id).Updates(features.UserEntityToModel(user))
	if tx.Error != nil {
		return 0, tx.Error
	}
	return id, nil
}

// SelectById implements user.UserData
func (repo *userData) SelectById(id uint) (features.UserEntity, error) {

	var user features.User
	tx := repo.db.First(&user, id)
	if tx.Error != nil {
		return features.UserEntity{}, errors.New("data tidak ditemukan")
	}
	data := features.UserModelToEntity(user)
	return data, nil
}

// Login implements user.UserData
func (repo *userData) Login(email string, password string) (uint, error) {
	var user features.User
	tx := repo.db.Where("email=?", email).First(&user)
	if tx.Error != nil {
		return 0, errors.New("email tidak ditemukan" + tx.Error.Error())
	}

	macth := helper.CheckPassword(password, user.Password)
	if !macth {
		return 0, errors.New("isi password sesuai data yang ada")
	}
	return user.ID, nil
}

// insert implements user.UserData
func (repo *userData) Insert(user features.UserEntity) (uint, error) {
	hashPassword, err := helper.HasPassword(user.Password)
	if err != nil {
		return 0, errors.New("error hashing password: " + err.Error())
	}
	user.Password = hashPassword
	userModel := features.UserEntityToModel(user)
	tx := repo.db.Create(&userModel)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("insert failes, row affected = 0")
	}

	return userModel.ID, nil
}

func New(data *gorm.DB) user.UserData {
	return &userData{
		db: data,
	}
}
