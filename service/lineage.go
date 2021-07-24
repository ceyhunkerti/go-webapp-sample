package service

import (
	"github.com/Screen17/catalog/appcontext"
	"github.com/Screen17/catalog/model"
)

type LineageService struct {
	context appcontext.Context
}

func NewLineageService(context appcontext.Context) *LineageService {
	return &LineageService{context: context}
}

func (s *LineageService) FindAllLineages() *[]model.Lineage {
	rep := s.context.GetRepository()
	lineage := model.Lineage{}
	result, err := lineage.FindAll(rep)
	if err != nil {
		s.context.GetLogger().GetZapLogger().Errorf(err.Error())
		return nil
	}
	return result
}
