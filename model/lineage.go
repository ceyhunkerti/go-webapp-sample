package model

type Lineage struct {
	ID              int                    `json:"id" gorm:"primary_key"`
	DagID           string                 `json:"name"`
	TaskID          string                 `json:"description"`
	SourceDatasetID *uint                  `json:"source_dataset_id"`
	SourceDataset   *Dataset               `json:"source_dataset"`
	TargetDatasetID *uint                  `json:"target_dataset_id"`
	TargetDataset   *Dataset               `json:"target_dataset"`
	Mapping         map[string]interface{} `json:"mapping"`
	Meta            map[string]interface{} `json:"meta"`
	CreatedOn       string                 `json:"created_on"`
	UpdatedOn       string                 `json:"updated_on"`
}
