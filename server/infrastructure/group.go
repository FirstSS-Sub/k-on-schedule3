package infrastructure

import (
	"github.com/FirstSS-Sub/k-on-schedule3/server/domain/model"
	"github.com/FirstSS-Sub/k-on-schedule3/server/domain/repository"
	"github.com/jinzhu/gorm"
)

type GroupRepository struct {
	DB *gorm.DB
}

// コンストラクタ
func NewGroupRepository(db *gorm.DB) repository.GroupRepository {
	return &GroupRepository{DB: db}
}

func (gr *GroupRepository) Insert(name string) error {
	// groupという箱の中に、参照したいmodel.Group型の構造体のアドレスを入れる
	group := new(model.Group)
	group.Name = name

	// &groupにしてしまうと、groupという箱自体のアドレスになってしまう
	if err := gr.DB.Create(group).Error; err != nil {
		return err
	}

	return nil
}

func (gr *GroupRepository) FindById(id uint) (*model.Group, error) {
	group := new(model.Group)
	group.ID = id

	if err := gr.DB.First(group).Error; err != nil {
		return nil, err
	}

	return group, nil
}

func (gr *GroupRepository) SearchSameName(name string) bool {
	group := new(model.Group)

	// 同じ名前が既に存在する場合はtrue
	foundSameName := !gr.DB.Where("name = ?", name).First(group).RecordNotFound()
	return foundSameName
}

func (gr *GroupRepository) Update(group *model.Group) error {
	if err := gr.DB.Save(group).Error; err != nil {
		return err
	}

	return nil
}

func (gr *GroupRepository) AddAssociation(group *model.Group, user *model.User) error {
	if err := gr.DB.Model(group).Association("Users").Append(user).Error; err != nil {
		return err
	}

	return nil
}

func (gr *GroupRepository) DeleteAssociation(group *model.Group, user *model.User) error {
	if err := gr.DB.Model(group).Association("Users").Delete(user).Error; err != nil {
		return err
	}

	return nil
}

func (gr *GroupRepository) Delete(id uint) error {
	group := new(model.Group)
	group.ID = id

	if err := gr.DB.Delete(group).Error; err != nil {
		return err
	}

	return nil
}
