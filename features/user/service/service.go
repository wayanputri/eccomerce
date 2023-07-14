package service

import (
	"belajar/bareng/features"
	"belajar/bareng/features/user"

	"github.com/go-playground/validator/v10"
)

type userService struct {
	userData user.UserData
	validate *validator.Validate
}

// Upgrade implements user.UserService.
func (service *userService) Upgrade(userId uint, user features.UserEntity) (uint, error) {
	id, err:=service.userData.Upgrade(userId,user)
	if err != nil{
		return 0,err
	}
	return id,nil
}

// Delete implements user.UserService
func (servise *userService) Delete(id uint) error {
	err := servise.userData.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

// Edit implements user.UserService
func (servis *userService) Edit(id uint, user features.UserEntity) (uint, error) {
	id, err := servis.userData.Update(id, user)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// GetById implements user.UserService
func (servis *userService) GetById(id uint) (features.UserEntity, error) {
	data, err := servis.userData.SelectById(id)
	if err != nil {
		return features.UserEntity{}, err
	}
	return data, nil
}

// Login implements user.UserService
func (servis *userService) Login(email string, password string) (uint, error) {
	data := features.LoginUser{
		Email:    email,
		Password: password,
	}
	errValidate := servis.validate.Struct(data)
	if errValidate != nil {
		return 0, errValidate
	}

	id, err := servis.userData.Login(email, password)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// Add implements user.UserService
func (servis *userService) Add(user features.UserEntity) (uint, error) {

	id, err := servis.userData.Insert(user)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func New(repo user.UserData) user.UserService {
	return &userService{
		userData: repo,
		validate: validator.New(),
	}
}
