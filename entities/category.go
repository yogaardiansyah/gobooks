// Back End dibuat oleh Yoga Ardiansyah

// entities/category.go
package entities

import "time"

type Category struct {
	Id         uint
	Name       string
	CreatedAt  time.Time
	UpddatedAt time.Time
}
