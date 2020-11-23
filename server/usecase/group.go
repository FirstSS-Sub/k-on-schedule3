package usecase

import (
	"github.com/FirstSS-Sub/k-on-schedule3/server/domain/model"
	"github.com/FirstSS-Sub/k-on-schedule3/server/domain/repository"
)

type GroupUsecase interface {
	Create(name string) error
	FindById(id uint) (*model.Group, error)
	ChangeName(id uint, newName string) (bool, error)
	AddUser(groupId uint, userUid string) error
	Leave(groupId uint, userUid string) error
	Delete(id uint) error
}

type groupUsecase struct {
	GroupRepository repository.GroupRepository
}

// コンストラクタ
func NewGroupUsecase(groupRepository repository.GroupRepository) GroupUsecase {
	return &groupUsecase{GroupRepository: groupRepository}
}

func (gu *groupUsecase) Create(name string) error {
	return gu.GroupRepository.Insert(name)
}

func (gu *groupUsecase) FindById(id uint) (*model.Group, error) {
	return gu.GroupRepository.FindById(id)
}

func (gu *groupUsecase) ChangeName(id uint, newName string) (bool, error) {
	// TODO
	if foundSameName := gu.GroupRepository.SearchSameName(newName); foundSameName == true {
		return true, nil
	}

	group := new(model.Group)
	group.ID = id
	group.Name = newName

	return false, gu.GroupRepository.Update(group)
}

func (gu *groupUsecase) AddUser(groupId uint, userUid string) error {
	group := new(model.Group)
	user := new(model.User)
	group.ID = groupId
	user.UserUID = userUid

	return gu.GroupRepository.AddAssociation(group, user)
}

func (gu *groupUsecase) Leave(groupId uint, userUid string) error {
	group := new(model.Group)
	user := new(model.User)
	group.ID = groupId
	user.UserUID = userUid

	return gu.GroupRepository.DeleteAssociation(group, user)
}

func (gu *groupUsecase) Delete(id uint) error {
	return gu.GroupRepository.Delete(id)
}
