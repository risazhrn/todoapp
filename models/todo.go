package models

type ToDo struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Task   string `json:"task"`
	Status string `json:"status"`
}
