package mysql

import (
	"github.com/go-kratos/kratos/v2/log"
	"go-template/internal/data/model"
	"go-template/internal/pkg/db"

	"github.com/go-kratos/kratos/v2/errors"

	"gorm.io/gorm"
)

type TemplateRepo struct {
	log *log.Helper
}

func NewTemplateRepo(logger log.Logger) *TemplateRepo {
	uc := &TemplateRepo{log: log.NewHelper(logger)}
	return uc
}

func (rp *TemplateRepo) CreateUser(tx *gorm.DB, user model.User) (id string, error error) {
	err := tx.Create(&user).Error
	if err != nil {
		return "", err
	}
	return user.Id, nil
}

func (rp *TemplateRepo) FindUserById(id string) (user model.User, err error) {
	result := db.DB.First(&user, "id = ?", id)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		err = result.Error
		return
	}
	return
}
