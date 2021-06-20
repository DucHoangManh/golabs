package models


type Price struct {
	Id uint `json:"id"`
	ProductId uint `json:"productId"` 
	Value float64 `json:"value"`
	CreatedAt uint64 `json:"createdAt" gorm:"AutoCreateTime"`
	Product Product `json:"-" gorm:"foreignKey:ProductId"`
}