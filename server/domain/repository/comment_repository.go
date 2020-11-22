package repository

import (
	"github.com/FirstSS-Sub/k-on-schedule2/server/domain/model"
)

type CommentRepository interface {
	Insert(userId uint, content string) error
	FindById(id uint) (*model.Comment, error) // 使うか？これ
	Update(comment *model.Comment) error
	Delete(id uint) error
}
