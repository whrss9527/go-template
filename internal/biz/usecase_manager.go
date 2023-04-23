package biz

type UsecaseManager struct {
	UserUsecase *TemplateUsecase
}

func NewUsecaseManager() *UsecaseManager {
	return &UsecaseManager{}
}
