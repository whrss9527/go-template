package service

import (
	"context"
	v1 "go-template/api/template_proj/v1"
)

// Login 登录
func (s *TemplateService) Login(ctx context.Context, _ *v1.LoginReq) (*v1.LoginReply, error) {
	res, err := s.uc.Login()
	if err != nil {
		return nil, err
	}
	return &v1.LoginReply{
		Id: res,
	}, nil
}
