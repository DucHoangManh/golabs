package models

type Image struct {
	Id uint `json:"id"`
	ProductId uint `json:"productId"`
	Url string `json:"url"`
}