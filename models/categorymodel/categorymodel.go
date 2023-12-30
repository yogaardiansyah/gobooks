// Back End dibuat oleh Yoga Ardiansyah

// models/categorymodel/categorymodel.go
package categorymodel

import (
	"go-web-native/config"
	"go-web-native/entities"
)

// GetAll mengambil semua kategori dari database.
func GetAll() []entities.Category {
	rows, err := config.DB.Query(`SELECT * FROM categories`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var categories []entities.Category

	// Iterasi melalui baris hasil query dan membuat objek kategori untuk setiap baris.
	for rows.Next() {
		var category entities.Category
		if err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpddatedAt); err != nil {
			panic(err)
		}
		categories = append(categories, category)
	}

	return categories
}

// Create membuat kategori baru dan menyimpannya ke dalam database.
func Create(category entities.Category) bool {
	result, err := config.DB.Exec(`
		INSERT INTO categories (name, created_at, updated_at) 
		VALUES (?, ?, ?)`,
		category.Name,
		category.CreatedAt,
		category.UpddatedAt,
	)

	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastInsertId > 0
}

// Detail mengembalikan detail kategori berdasarkan ID.
func Detail(id int) entities.Category {
	row := config.DB.QueryRow(`SELECT id, name FROM categories WHERE id = ? `, id)

	var category entities.Category

	if err := row.Scan(&category.Id, &category.Name); err != nil {
		panic(err.Error())
	}

	return category
}

// Update memperbarui informasi kategori berdasarkan ID.
func Update(id int, category entities.Category) bool {
	query, err := config.DB.Exec(`UPDATE categories SET name = ?, updated_at = ? where id = ?`, category.Name, category.UpddatedAt, id)
	if err != nil {
		panic(err)
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0
}

// Delete menghapus kategori berdasarkan ID.
func Delete(id int) error {
	_, err := config.DB.Exec("DELETE FROM categories WHERE id = ?", id)
	return err
}
