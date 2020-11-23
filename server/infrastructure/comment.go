package infrastructure

import (
	"github.com/FirstSS-Sub/k-on-schedule3/server/domain/model"
	"github.com/FirstSS-Sub/k-on-schedule3/server/domain/repository"
	"github.com/jinzhu/gorm"
)

type CommentRepository struct {
	DB *gorm.DB
}

// コンストラクタ
func NewCommentRepository(db *gorm.DB) repository.CommentRepository {
	return &CommentRepository{DB: db}
}

func (cr *CommentRepository) Insert(userUid string, groupId uint, content string) error {
	// commentという箱の中に、参照したいmodel.Comment型の構造体のアドレスを入れる
	comment := new(model.Comment)
	comment.UserUID = userUid
	comment.GroupID = groupId
	comment.Content = content

	// &commentにしてしまうと、commentという箱自体のアドレスになってしまう
	if err := cr.DB.Create(comment).Error; err != nil {
		return err
	}

	return nil
}

func (cr *CommentRepository) FindById(id uint) (*model.Comment, error) {
	comment := new(model.Comment)
	comment.ID = id

	if err := cr.DB.First(comment).Error; err != nil {
		return nil, err
	}

	return comment, nil
}

func (cr *CommentRepository) Update(comment *model.Comment) error {
	if err := cr.DB.Save(comment).Error; err != nil {
		return err
	}

	return nil
}

func (cr *CommentRepository) Delete(id uint) error {
	comment := new(model.Comment)
	comment.ID = id

	if err := cr.DB.Delete(comment).Error; err != nil {
		return err
	}

	return nil
}
