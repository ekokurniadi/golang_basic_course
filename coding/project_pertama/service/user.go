package service

import (
	"project_pertama/entity"
	"project_pertama/input"
	"project_pertama/repository"
)

type UserService interface {
	Save(input input.UserInput) (entity.User, error)
	UploadProfilePicture(ID int, fileImage string) (entity.User, error)
	FindByID(ID int) (entity.User, error)
	Update(inputID input.UserInputID, inputData input.UserInput) (entity.User, error)
	FindAll(search string, page int, size int) ([]entity.User, error)
	TotalFetchData(search string, page int, size int) (int, error)
	DeleteUser(ID int) (bool, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *userService {
	return &userService{userRepository}
}

func (s *userService) Save(input input.UserInput) (entity.User, error) {
	var user entity.User
	user.NamaLengkap = input.NamaLengkap
	user.JenisKelamin = input.JenisKelamin
	user.UserName = input.UserName
	user.Password = input.Password
	user.Role = input.Role

	newUser, err := s.userRepository.Save(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (s *userService) UploadProfilePicture(ID int, fileImage string) (entity.User, error) {
	user, err := s.userRepository.FindByID(ID)
	if err != nil {
		return user, err
	}
	user.ProfilePicture = fileImage

	foundUser, err := s.userRepository.UploadProfilePicture(user)
	if err != nil {
		return foundUser, err
	}
	return foundUser, nil

}

func (s *userService) Update(inputID input.UserInputID, inputData input.UserInput) (entity.User, error) {
	user, err := s.userRepository.FindByID(inputID.ID)
	if err != nil {
		return user, err
	}
	user.NamaLengkap = inputData.NamaLengkap
	user.JenisKelamin = inputData.JenisKelamin
	user.Password = inputData.Password
	user.Role = inputData.Role

	updatedUser, err := s.userRepository.Update(user)
	if err != nil {
		return updatedUser, err
	}
	return updatedUser, nil
}

func (s *userService) FindByID(ID int) (entity.User, error) {
	user, err := s.userRepository.FindByID(ID)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *userService) FindAll(search string, page int, size int) ([]entity.User, error) {
	users, err := s.userRepository.FindAll(search, page, size)
	if err != nil {
		return users, err
	}
	return users, nil
}

func (s *userService) TotalFetchData(search string, page int, size int) (int, error) {
	totalData, err := s.userRepository.TotalFetchData(search, page, size)
	if err != nil {
		return totalData, err
	}
	return totalData, nil
}

func (s *userService) DeleteUser(ID int) (bool, error) {
	deleteUser, err := s.userRepository.DeleteUser(ID)
	if err != nil {
		return deleteUser, err
	}
	return deleteUser, nil
}
