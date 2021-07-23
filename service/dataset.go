package service

import (
	"github.com/Screen17/catalog/appcontext"
	"github.com/Screen17/catalog/model"
)

type DatasetService struct {
	context appcontext.Context
}

func NewDatasetService(context appcontext.Context) *DatasetService {
	return &DatasetService{context: context}
}

func (m *DatasetService) FindAllDatasets() *[]model.Dataset {
	rep := m.context.GetRepository()
	dataset := model.Dataset{}
	result, err := dataset.FindAll(rep)
	if err != nil {
		m.context.GetLogger().GetZapLogger().Errorf(err.Error())
		return nil
	}
	return result
}
