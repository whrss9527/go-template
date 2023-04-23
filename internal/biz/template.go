package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	"go-template/internal/conf"
	"go-template/internal/data/mysql"
)

type UserStatus int8

type TemplateUsecase struct {
	log   *log.Helper
	ucMgr *UsecaseManager
	conf  *conf.Biz
	repo  *mysql.TemplateRepo
}

func NewTemplateUsecase(logger log.Logger, ucMgr *UsecaseManager, biz *conf.Biz, repo *mysql.TemplateRepo) *TemplateUsecase {
	uc := &TemplateUsecase{log: log.NewHelper(logger)}
	ucMgr.UserUsecase = uc
	uc.ucMgr = ucMgr
	uc.conf = biz
	uc.repo = repo
	return uc
}

func (uc *TemplateUsecase) Login() (string, error) {

	userId := "123"
	return userId, nil

}
