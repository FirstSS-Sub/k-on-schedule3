package repository

import (
	"github.com/FirstSS-Sub/k-on-schedule2/server/domain/model"
)

type GroupRepository interface {
	Insert(name string) error
	FindById(id uint) (*model.Group, error)
	// AddUser(userId, groupId uint) error
	Update(group *model.Group) error
	Delete(id uint) error
}
