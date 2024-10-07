package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ItemName    string  `gorm:"not null" json:"item_name"`
	Description string  `json:"description"`
	Quantity    int     `gorm:"not null" json:"quantity"`
	Code        float64 `json:"code"`
	Cartegory   string  `json:"cartegory"`
	Caterlogue  string  `json:"caterlogue"`
}
