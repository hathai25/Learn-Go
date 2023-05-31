package models

type IDCard struct {
	ID        uint   `json:"id" gorm:"primary_key" gorm:"auto_increment"`
	IDNUmber  string `json:"id_number"`
	PersonID  uint   `json:"person_id"`
	Person    Person `gorm:"foreignKey:PersonID"`
	BornAt    string `json:"born_at"`
	ExpiredAt string `json:"expired_at"`
}
