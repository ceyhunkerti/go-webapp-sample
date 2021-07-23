package service

import (
	"github.com/Screen17/catalog/appcontext"
	"github.com/Screen17/catalog/model"
)

type CategoryService struct {
	context appcontext.Context
}

func NewCategoryService(context appcontext.Context) *CategoryService {
	return &CategoryService{context: context}
}

func (m *CategoryService) FindAllCategories() *[]model.Category {
	rep := m.context.GetRepository()
	category := model.Category{}
	result, err := category.FindAll(rep)
	if err != nil {
		m.context.GetLogger().GetZapLogger().Errorf(err.Error())
		return nil
	}
	return result
}
