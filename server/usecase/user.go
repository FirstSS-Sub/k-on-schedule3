package usecase

import (
	"github.com/FirstSS-Sub/k-on-schedule3/server/domain/model"
	"github.com/FirstSS-Sub/k-on-schedule3/server/domain/repository"
)

type UserUsecase interface {
	Create(name string) error
	FindByUid(uid string) (*model.User, error)
	UpdateSchedule(uid string, thu, fri, sat, sun, mon, tue, wed string) error
	ChangeName(uid string, newName string) (bool, error)
	Delete(uid string) error
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

func (uu *userUsecase) FindByUid(uid string) (*model.User, error) {
	return uu.UserRepository.FindByUid(uid)
}

func (uu *userUsecase) UpdateSchedule(uid string, thu, fri, sat, sun, mon, tue, wed string) error {
	// TODO
	user := new(model.User)
	user.UserUID = uid
	user.Thu = thu
	user.Fri = fri
	user.Sat = sat
	user.Sun = sun
	user.Mon = mon
	user.Tue = tue
	user.Wed = wed

	return uu.UserRepository.Update(user)
}

func (uu *userUsecase) ChangeName(uid string, newName string) (bool, error) {
	// TODO
	if foundSameName := uu.UserRepository.SearchSameName(newName); foundSameName == true {
		return true, nil
	}

	user := new(model.User)
	user.UserUID = uid
	user.Name = newName

	return false, uu.UserRepository.Update(user)
}

func (uu *userUsecase) Delete(uid string) error {
	return uu.UserRepository.Delete(uid)
}
