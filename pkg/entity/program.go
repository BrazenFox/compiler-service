package entity

type Program struct {
	Language string `json:"language" binding:"required"`
	Code     string `json:"code"  binding:"required"`
}
