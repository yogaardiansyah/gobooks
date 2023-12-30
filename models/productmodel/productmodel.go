// Back End dibuat oleh Yoga Ardiansyah

// models/productmodel/productmodel.go
package productmodel

import (
	"go-web-native/config"
	"go-web-native/entities"
)

// Getall mengambil semua produk dari database.
func Getall() []entities.Product {
	rows, err := config.DB.Query(`
		SELECT 
			products.id, 
			products.name, 
			categories.name as category_name,
			products.issn, 
			products.description, 
			products.created_at, 
			products.updated_at,
			products.image_name,
			products.image_path
		FROM products
		JOIN categories ON products.category_id = categories.id
	`)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var products []entities.Product

	// Iterasi melalui baris hasil query dan membuat objek produk untuk setiap baris.
	for rows.Next() {
		var product entities.Product
		if err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Category.Name,
			&product.ISSN,
			&product.Description,
			&product.CreatedAt,
			&product.UpdatedAt,
			&product.ImageName,
			&product.ImagePath,
		); err != nil {
			panic(err)
		}

		products = append(products, product)
	}

	return products
}

// Create membuat produk baru dan menyimpannya ke dalam database.
func Create(product entities.Product) bool {
	result, err := config.DB.Exec(`
		INSERT INTO products(
			name, category_id, issn, description, created_at, updated_at, image_name, image_path
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		product.Name,
		product.Category.Id,
		product.ISSN,
		product.Description,
		product.CreatedAt,
		product.UpdatedAt,
		product.ImageName,
		product.ImagePath,
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

// Detail mengembalikan detail produk berdasarkan ID.
func Detail(id int) entities.Product {
	row := config.DB.QueryRow(`
		SELECT 
			products.id, 
			products.name, 
			categories.name as category_name,
			products.issn, 
			products.description, 
			products.created_at, 
			products.updated_at,
			products.image_name,
			products.image_path
		FROM products
		JOIN categories ON products.category_id = categories.id
		WHERE products.id = ?
	`, id)

	var product entities.Product

	err := row.Scan(
		&product.Id,
		&product.Name,
		&product.Category.Name,
		&product.ISSN,
		&product.Description,
		&product.CreatedAt,
		&product.UpdatedAt,
		&product.ImageName,
		&product.ImagePath,
	)

	if err != nil {
		panic(err)
	}

	return product
}

// Update memperbarui informasi produk berdasarkan ID.
func Update(id int, product entities.Product) bool {
	query, err := config.DB.Exec(`
        UPDATE products SET 
            name = ?, 
            category_id = ?,
            issn = ?,
            description = ?,
            updated_at = ?,
            image_name = ?,
            image_path = ?
        WHERE id = ?`,
		product.Name,
		product.Category.Id,
		product.ISSN,
		product.Description,
		product.UpdatedAt,
		product.ImageName,
		product.ImagePath,
		id,
	)

	if err != nil {
		panic(err)
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0
}

// Delete menghapus produk berdasarkan ID.
func Delete(id int) error {
	_, err := config.DB.Exec("DELETE FROM products WHERE id = ?", id)
	return err
}
