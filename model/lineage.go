package model

import (
	"encoding/json"

	"github.com/Screen17/catalog/repository"
	"gorm.io/datatypes"
)

type Lineage struct {
	ID              int            `json:"id" gorm:"primary_key"`
	DagID           string         `json:"name"`
	TaskID          string         `json:"description"`
	SourceDatasetID *uint          `json:"source_dataset_id"`
	SourceDataset   *Dataset       `json:"source_dataset"`
	TargetDatasetID *uint          `json:"target_dataset_id"`
	TargetDataset   *Dataset       `json:"target_dataset"`
	Mapping         datatypes.JSON `json:"mapping"`
	Meta            datatypes.JSON `json:"meta"`
	CreatedOn       string         `json:"created_on"`
	UpdatedOn       string         `json:"updated_on"`
}

func (Lineage) TableName() string {
	return "lineages"
}

func NewLineage(dagID string, taskID string, sourceDatasetID uint, targetDatasetID uint, mapping datatypes.JSON, meta datatypes.JSON) *Lineage {
	return &Lineage{DagID: dagID, TaskID: taskID, SourceDatasetID: &sourceDatasetID, TargetDatasetID: &targetDatasetID, Mapping: mapping, Meta: meta}
}

func (l *Lineage) FindAll(rep repository.Repository) (*[]Lineage, error) {
	var lineages []Lineage

	if err := rep.Preload("TargetDataset").Preload("SourceDataset").Find(&lineages).Error; err != nil {
		return nil, err
	}
	return &lineages, nil
}

func (l *Lineage) ToString() (string, error) {
	bytes, error := json.Marshal(l)
	return string(bytes), error
}
