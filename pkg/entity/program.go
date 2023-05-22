package entity

type Program struct {
	Language string `json:"language" binding:"required"`
	Code     string `json:"code"  binding:"required"`
	Name     string `json:"name"  binding:"required"`
	UserId   string `json:"user_id"  binding:"required"`
}
