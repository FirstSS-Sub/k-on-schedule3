package infrastructure

import (
	"github.com/FirstSS-Sub/k-on-schedule2/server/domain/model"
	"github.com/FirstSS-Sub/k-on-schedule2/server/domain/repository"
	"github.com/jinzhu/gorm"
)

type GroupRepository struct {
	DB *gorm.DB
}

// コンストラクタ
func NewGroupRepository(db *gorm.DB) repository.GroupRepository {
	return &GroupRepository{DB: db}
}

func (ur *GroupRepository) Insert(name string) error {
	// groupという箱の中に、参照したいmodel.Group型の構造体のアドレスを入れる
	group := new(model.Group)
	group.Name = name

	// &groupにしてしまうと、groupという箱自体のアドレスになってしまう
	if err := ur.DB.Create(group).Error; err != nil {
		return err
	}

	return nil
}

func (ur *GroupRepository) FindById(id uint) (*model.Group, error) {
	group := new(model.Group)
	group.ID = id

	if err := ur.DB.First(group).Error; err != nil {
		return nil, err
	}

	return group, nil
}

func (ur *GroupRepository) Update(group *model.Group) error {
	if err := ur.DB.Save(group).Error; err != nil {
		return err
	}

	return nil
}

func (ur *GroupRepository) Delete(id uint) error {
	group := new(model.Group)
	group.ID = id

	if err := ur.DB.Delete(group).Error; err != nil {
		return err
	}

	return nil
}
