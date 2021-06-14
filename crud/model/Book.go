package model

type Book struct {
	Id      int64
	Title   string   `json:"title"`
	Authors []Author `json:"authors"`
	Rating  float32  `json:"rating"`
}
