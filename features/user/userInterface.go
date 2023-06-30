package user

import "belajar/bareng/features"

type UserData interface {
	Insert(user features.UserEntity) (uint,error)
	Login(email string, password string) (uint, error)
	SelectById(id uint) (features.UserEntity,error)
	Update(id uint, user features.UserEntity) (uint,error)
	Delete(id uint)error
}

type UserService interface{
	Add(user features.UserEntity)(uint,error)
	Login(email string, password string) (uint, error)
	GetById(id uint) (features.UserEntity,error)
	Edit(id uint, user features.UserEntity) (uint,error)
	Delete(id uint) error
}