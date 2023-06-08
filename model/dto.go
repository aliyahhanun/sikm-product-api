package model

type ProductDetail struct {
	ID       uint     `json:"id"`
	Name     string   `json:"name"`
	Price    int      `json:"price"`
	User     User     `json:"poster"`
	Category Category `json:"category"`
}
type ProductRequest struct {
	Name       string `json:"name" binding:"required,min=8"`          // validator read more : https://github.com/go-playground/validator#other
	Price      int    `json:"price" binding:"required,gt=0,numeric"`  // read more : https://github.com/go-playground/validator#comparisons
	CategoryID int    `json:"category_id" binding:"required,numeric"` // if we dont use Gin, validate : "required"
	PostedBy   int    `json:"posted_by"`
}

type UpdateProductRequest struct {
	Name       string `json:"name" binding:"required,min=8"`
	Price      int    `json:"price" binding:"required,gt=0"`
	CategoryID int    `json:"category_id" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}
