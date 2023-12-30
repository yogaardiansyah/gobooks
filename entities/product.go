// Back End dibuat oleh Yoga Ardiansyah

// entities/product.go
package entities

import "time"

type Product struct {
	Id          uint
	Name        string
	Category    Category
	ISSN        int64
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	ImageName   string
	ImagePath   string
}
