package repository

import (
	"github.com/FirstSS-Sub/k-on-schedule3/server/domain/model"
)

type CommentRepository interface {
	Insert(userUid string, groupId uint, content string) error
	FindById(id uint) (*model.Comment, error) // 使うか？これ
	Update(comment *model.Comment) error
	Delete(id uint) error
}
