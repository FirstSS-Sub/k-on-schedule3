package repository

import (
	"github.com/FirstSS-Sub/k-on-schedule2/server/domain/model"
)

type UserRepository interface {
	Insert(name string) error
	FindById(id uint) (*model.User, error)
	SearchSameName(name string) bool
	Update(user *model.User) error
	Delete(id uint) error
	// GetComments(id uint) ([]*domain.Comment, error)
	// GetGroups(id uint) ([]*domain.Group, error)

	// AddComment(DB *gorm.DB, name, thu, fri, sat, sun, mon, tue, wed string, comments []*domain.Comment, groups []*domain.Group) error
	// AddGroup(DB *gorm.DB, name, thu, fri, sat, sun, mon, tue, wed string, comments []*domain.Comment, groups []*domain.Group) error
	// GetByUserID(DB *gorm.DB, id uint) (*domain.User, error)
}
