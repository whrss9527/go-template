package service

import (
	v1 "go-template/api/template_proj/v1"
	"go-template/internal/biz"
)

// TemplateService service.
type TemplateService struct {
	v1.UnimplementedTemplateProjServer

	uc *biz.TemplateUsecase
}

// NewTemplateService new a greeter service.
func NewTemplateService(uc *biz.TemplateUsecase) *TemplateService {
	return &TemplateService{uc: uc}
}
