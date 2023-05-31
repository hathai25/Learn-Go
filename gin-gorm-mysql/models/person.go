package models

type Person struct {
	ID   uint   `json:"id" gorm:"primary_key" gorm:"auto_increment"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
