package model

import (
	"encoding/json"

	"github.com/Screen17/catalog/repository"
	"github.com/lib/pq"
	"gorm.io/datatypes"
)

type Dataset struct {
	ID          int            `json:"id" gorm:"primary_key"`
	Name        string         `json:"name" validate:"required"`
	Description string         `json:"description"`
	Columns     pq.StringArray `json:"columns" gorm:"type:varchar(128)[]"`
	Meta        datatypes.JSON `json:"meta"`
	CreatedOn   string         `json:"created_on"`
	UpdatedOn   string         `json:"updated_on"`
}

func (Dataset) TableName() string {
	return "datasets"
}

func NewDataset(name string, description string, columns pq.StringArray, meta datatypes.JSON) *Dataset {
	return &Dataset{Name: name, Description: description, Columns: columns, Meta: meta}
}

func (d *Dataset) FindAll(rep repository.Repository) (*[]Dataset, error) {
	var datasets []Dataset

	if err := rep.Find(&datasets).Error; err != nil {
		return nil, err
	}
	return &datasets, nil
}

// ToString is return string of object
func (d *Dataset) ToString() (string, error) {
	bytes, error := json.Marshal(d)
	return string(bytes), error
}
