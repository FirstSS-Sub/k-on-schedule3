package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserUID string
	Name    string
	Schedule
	Comments []Comment `gorm:"foreignKey:UserUID"` // ; はいらない？
	Groups   []*Group  `gorm:"many2many:user_group;"`
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
