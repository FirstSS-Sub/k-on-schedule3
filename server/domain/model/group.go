package model

import (
	"github.com/jinzhu/gorm"
)

type Group struct {
	gorm.Model
	Name  string
	Users []*User `gorm:"many2many:user_group;"`
}
