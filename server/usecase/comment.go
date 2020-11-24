package usecase

import (
	"github.com/FirstSS-Sub/k-on-schedule3/server/domain/model"
	"github.com/FirstSS-Sub/k-on-schedule3/server/domain/repository"
)

type CommentUsecase interface {
	Create(userUid string, groupId uint, content string) error
	FindById(id uint) (*model.Comment, error) // 使うか？これ
	Update(id uint, content string) error
	Delete(id uint) error
}

type commentUsecase struct {
	CommentRepository repository.CommentRepository
}

// コンストラクタ
func NewCommentUsecase(commentRepository repository.CommentRepository) CommentUsecase {
	return &commentUsecase{CommentRepository: commentRepository}
}

func (cu *commentUsecase) Create(userUid string, groupId uint, content string) error {
	comment, err := cu.CommentRepository.Insert(userUid, groupId, content)
	if err != nil {
		return err
	}

	user := new(model.User)
	user.UserUID = userUid

	return cu.CommentRepository.AddAssociation(comment, user)
}

func (cu *commentUsecase) FindById(id uint) (*model.Comment, error) {
	return cu.CommentRepository.FindById(id)
}

func (cu *commentUsecase) Update(id uint, content string) error {
	// TODO

	comment := new(model.Comment)
	comment.ID = id
	comment.Content = content

	return cu.CommentRepository.Update(comment)
}

func (cu *commentUsecase) Delete(id uint) error {
	return cu.CommentRepository.Delete(id)
}
