package models

import "gorm.io/gorm"

type Product struct {
	Id uint `json:"id"`
	CategoryId uint `json:"categoryId"`
	Images []Image `json:"images" gorm:"foreignKey:ProductId"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	IsSale bool `json:"isSale"`
	Rating float32 `json:"rating"`
	Reviews []Review `json:"-" gorm:"foreignKey:ProductId"`
	CreatedAt int64 `json:"createdAt" gorm:"AutoCreateTime"`
	ModifiedAt int64 `json:"modifiedAt" gorm:"AutoUpdateTime"`
}

func (product *Product) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Product{}).Count(&total)
	return total
}
