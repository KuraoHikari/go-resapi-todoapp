package dto

type CreateProductRequest struct {
	Name   string `json:"name" form:"name" binding:"required,min=1"`
	Bounty uint64 `json:"price" form:"price" binding:"required"`
	Image  string `json:"image" form:"image"`
}

type UpdateProductRequest struct {
	ID     int64  `json:"id" form:"id"`
	Name   string `json:"name" form:"name" binding:"required,min=1"`
	Bounty uint64 `json:"price" form:"price" binding:"required"`
	Image  string `json:"image" form:"image"`
}
