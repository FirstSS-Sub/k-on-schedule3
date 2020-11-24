package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserUID string
	Name    string
	Schedule
	Comments []Comment `gorm:"foreignKey:UserUID;references:UserUID;"`
	Groups   []*Group  `gorm:"many2many:user_group;references:UserUID;"`
}

type Schedule struct {
	Thu string
	Fri string
	Sat string
	Sun string
	Mon string
	Tue string
	Wed string
}
