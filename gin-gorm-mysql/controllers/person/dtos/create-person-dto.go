package dtos

type CreatePersonDTO struct {
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required"`
}
