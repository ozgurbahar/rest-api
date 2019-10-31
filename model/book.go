package model

type Book struct {
	Id    int    `json:"id" gorm:"primary_key"`
	Name  string `json:"name" gorm:"column:name"`
	Title string `json:"title" gorm:"column:title"`
}

func (Book) TableName() string {
	return "book"
}
