package models

type Cart struct {
	Id uint `json:"-"`
	ProductId uint `json:"productId"`
	Quantity uint `json:"quantity"`
	Product Product `json:"product" gorm:"foreignKey:ProductId"`
	CreatedAt int64 `json:"createdAt" gorm:"AutoCreateTime"`
	ModifiedAt int64 `json:"modifiedAt" gorm:"AutoUpdateTime"`
}