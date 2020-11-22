package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name string
	Schedule
	Comments []Comment `gorm:"foreignKey:UserID"` // ; はいらない？
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
