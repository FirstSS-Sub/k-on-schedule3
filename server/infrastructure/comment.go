package infrastructure

import (
	"github.com/FirstSS-Sub/k-on-schedule2/server/domain/model"
	"github.com/FirstSS-Sub/k-on-schedule2/server/domain/repository"
	"github.com/jinzhu/gorm"
)

type CommentRepository struct {
	DB *gorm.DB
}

// コンストラクタ
func NewCommentRepository(db *gorm.DB) repository.CommentRepository {
	return &CommentRepository{DB: db}
}

func (ur *CommentRepository) Insert(userId uint, content string) error {
	// commentという箱の中に、参照したいmodel.Comment型の構造体のアドレスを入れる
	comment := new(model.Comment)
	comment.UserID = userId
	comment.Content = content

	// &commentにしてしまうと、commentという箱自体のアドレスになってしまう
	if err := ur.DB.Create(comment).Error; err != nil {
		return err
	}

	return nil
}

func (ur *CommentRepository) FindById(id uint) (*model.Comment, error) {
	comment := new(model.Comment)
	comment.ID = id

	if err := ur.DB.First(comment).Error; err != nil {
		return nil, err
	}

	return comment, nil
}

func (ur *CommentRepository) Update(comment *model.Comment) error {
	if err := ur.DB.Save(comment).Error; err != nil {
		return err
	}

	return nil
}

func (ur *CommentRepository) Delete(id uint) error {
	comment := new(model.Comment)
	comment.ID = id

	if err := ur.DB.Delete(comment).Error; err != nil {
		return err
	}

	return nil
}
