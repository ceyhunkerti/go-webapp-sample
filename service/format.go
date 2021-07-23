package service

import (
	"github.com/Screen17/catalog/appcontext"
	"github.com/Screen17/catalog/model"
)

// FormatService is a service for managing master data such as format and category.
type FormatService struct {
	context appcontext.Context
}

// NewFormatService is constructor.
func NewFormatService(context appcontext.Context) *FormatService {
	return &FormatService{context: context}
}

// FindAllFormats returns the list of all formats.
func (m *FormatService) FindAllFormats() *[]model.Format {
	rep := m.context.GetRepository()
	format := model.Format{}
	result, err := format.FindAll(rep)
	if err != nil {
		m.context.GetLogger().GetZapLogger().Errorf(err.Error())
		return nil
	}
	return result
}
