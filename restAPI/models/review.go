package models

type Review struct {
	Id      uint   `json:"id"`
	ProductId uint `json:"productId"`
	Content string `json:"content"`
	Rating  int    `json:"rating"`
}