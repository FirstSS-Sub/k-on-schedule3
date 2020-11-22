package infrastructure

import (
	"github.com/FirstSS-Sub/k-on-schedule2/server/domain/model"
	"github.com/FirstSS-Sub/k-on-schedule2/server/domain/repository"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

// コンストラクタ
func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &UserRepository{DB: db}
}

func (ur *UserRepository) Insert(name string) error {
	// userという箱の中に、参照したいmodel.User型の構造体のアドレスを入れる
	user := new(model.User)
	user.Name = name

	// &userにしてしまうと、userという箱自体のアドレスになってしまう
	if err := ur.DB.Create(user).Error; err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) FindById(id uint) (*model.User, error) {
	user := new(model.User)
	user.ID = id

	if err := ur.DB.First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *UserRepository) SearchSameName(name string) bool {
	user := new(model.User)

	// 同じ名前が既に存在する場合はtrue
	foundSameName := !ur.DB.Where("name = ?", name).First(user).RecordNotFound()
	return foundSameName
}

func (ur *UserRepository) Update(user *model.User) error {
	if err := ur.DB.Save(user).Error; err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) Delete(id uint) error {
	user := new(model.User)
	user.ID = id

	if err := ur.DB.Delete(user).Error; err != nil {
		return err
	}

	return nil
}
