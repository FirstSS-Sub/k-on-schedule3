package usecase

import (
	"github.com/FirstSS-Sub/k-on-schedule2/server/domain/model"
	"github.com/FirstSS-Sub/k-on-schedule2/server/domain/repository"
)

type UserUsecase interface {
	Create(name string) error
	FindById(id uint) (*model.User, error)
	UpdateSchedule(id uint, thu, fri, sat, sun, mon, tue, wed string) error
	ChangeName(id uint, newName string) (bool, error)
	Delete(id uint) error
}

type userUsecase struct {
	UserRepository repository.UserRepository
}

// コンストラクタ
func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{UserRepository: userRepository}
}

func (uu *userUsecase) Create(name string) error {
	return uu.UserRepository.Insert(name)
}

func (uu *userUsecase) FindById(id uint) (*model.User, error) {
	return uu.UserRepository.FindById(id)
}

func (uu *userUsecase) UpdateSchedule(id uint, thu, fri, sat, sun, mon, tue, wed string) error {
	// TODO
	user := new(model.User)
	user.ID = id
	user.Thu = thu
	user.Fri = fri
	user.Sat = sat
	user.Sun = sun
	user.Mon = mon
	user.Tue = tue
	user.Wed = wed

	return uu.UserRepository.Update(user)
}

func (uu *userUsecase) ChangeName(id uint, newName string) (bool, error) {
	// TODO
	if foundSameName := uu.UserRepository.SearchSameName(newName); foundSameName == true {
		return true, nil
	}

	user := new(model.User)
	user.ID = id
	user.Name = newName

	return false, uu.UserRepository.Update(user)
}

func (uu *userUsecase) Delete(id uint) error {
	return uu.UserRepository.Delete(id)
}
