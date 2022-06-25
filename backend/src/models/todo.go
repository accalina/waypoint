package models

type Todo struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Content  string `json:"content"`
	IsDone   bool   `json:"is_done" gorm:"default:false"`
	IsDelete bool   `json:"is_delete" gorm:"default:false"`
}
