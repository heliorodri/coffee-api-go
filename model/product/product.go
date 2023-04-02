package entity

//the gorm.Model struct provides some common fields like ID, CreatedAt, UpdatedAt, and DeletedAt
import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
