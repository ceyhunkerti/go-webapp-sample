package service

import (
	"github.com/Screen17/catalog/appcontext"
	"github.com/Screen17/catalog/model"
)

type SearchService struct {
	context appcontext.Context
}

func NewSearchService(context appcontext.Context) *SearchService {
	return &SearchService{context: context}
}

func (m *SearchService) Search() *[]model.Dataset {
	rep := m.context.GetRepository()
	dataset := model.Dataset{}
	result, err := dataset.FindAll(rep)
	if err != nil {
		m.context.GetLogger().GetZapLogger().Errorf(err.Error())
		return nil
	}
	return result
}
