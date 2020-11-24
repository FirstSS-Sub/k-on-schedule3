package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserUID string
	Name    string
	Schedule []Schedule
	Comments []Comment `gorm:"foreignKey:UserUID;references:UserUID;"`
	Groups   []*Group  `gorm:"many2many:user_group;references:UserUID;"`
}

type Schedule struct {
	Flags string
}
