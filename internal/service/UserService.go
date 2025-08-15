package service

import (
	"Rent-And-Buy-App/internal/dtos"
	"Rent-And-Buy-App/internal/entity"
	"Rent-And-Buy-App/internal/repository"
	"errors"
	"strconv"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (us *UserService) GetAll() ([]entity.User, error) {
	users, err := us.repo.GetAll()
	return users, err
}

func (us *UserService) GetByID(id string) (*entity.User, error) {
	//var user *entity.User
	ID, err := strconv.ParseUint(id, 10, 32)
	user, err := us.repo.GetById(uint(ID))
	if err != nil || user == nil {
		errors.New("User not found")
	}

	return user, err
}

func (us *UserService) UpdateUser(user *entity.User, updateDto dtos.UpdateUserDto) (*entity.User, error) {
	id := user.ID
	isExist, err := us.repo.GetById(id)
	if err != nil || isExist == nil {
		return nil, errors.New("User not found")
	}

	user.Name = updateDto.Name
	user.Surname = updateDto.Surname
	user.Email = updateDto.Email
	user.IsSeller = updateDto.IsSeller

	err = us.repo.Update(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *UserService) DeleteUser(Id string) error {
	user, err := us.GetByID(Id)
	if err != nil {
		return errors.New("could not convert into uint")
	}
	err = us.repo.Delete(user)
	return err
}
