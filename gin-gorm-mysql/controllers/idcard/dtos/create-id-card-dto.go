package dtos

type CreateIDCardDTO struct {
	IDNumber  string `json:"id_number" binding:"required"`
	PersonID  uint   `json:"person_id" binding:"required"`
	BornAt    string `json:"born_at"`
	ExpiredAt string `json:"expired_at"`
}
