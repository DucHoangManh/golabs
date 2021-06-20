package models

type Category struct {
	Id uint `json:"id"`
	Name string `json:"name"`
	Products []Product `json:"-" gorm:"foreignKey:CategoryId"`
}