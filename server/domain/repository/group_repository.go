package repository

import (
	"github.com/FirstSS-Sub/k-on-schedule3/server/domain/model"
)

type GroupRepository interface {
	Insert(name string) error
	FindById(id uint) (*model.Group, error)
	SearchSameName(name string) bool
	Update(group *model.Group) error
	AddAssociation(group *model.Group, user *model.User) error
	DeleteAssociation(group *model.Group, user *model.User) error
	Delete(id uint) error
}
