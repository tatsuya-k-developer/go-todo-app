package models

// Todo の構造体
type Todo struct {
	Id     uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Title  string `json:"title"`
	Status string `json:"status"`
}
