package model

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model
	UserUID string
	GroupID uint
	Content string
}
