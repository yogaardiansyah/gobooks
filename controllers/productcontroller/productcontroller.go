// Back End dibuat oleh Yoga Ardiansyah

// controllers/productcontroller/productcontroller.go
package productcontroller

import (
	"fmt"
	"go-web-native/entities"
	"go-web-native/models/categorymodel"
	"go-web-native/models/productmodel"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"text/template"
	"time"
)

// Index menangani permintaan untuk halaman indeks produk.
func Index(w http.ResponseWriter, r *http.Request) {
	// Mengambil semua produk dari model
	products := productmodel.Getall()

	// Membuat peta data yang akan digunakan dalam template
	data := map[string]interface{}{
		"products": products,
	}

	// Menganalisis template HTML
	temp, err := template.ParseFiles("views/product/index.html")
	if err != nil {
		panic(err)
	}

	// Menjalankan template dengan data yang telah disiapkan
	temp.Execute(w, data)
}

// Add menangani permintaan untuk menambahkan produk baru.
func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Jika metode HTTP adalah GET, tampilkan formulir tambah produk
		temp, err := template.ParseFiles("views/product/create.html")
		if err != nil {
			panic(err)
		}

		// Mengambil daftar kategori untuk ditampilkan dalam formulir
		categories := categorymodel.GetAll()
		data := map[string]interface{}{
			"categories": categories,
		}

		// Menjalankan template dengan data yang telah disiapkan
		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		// Jika metode HTTP adalah POST, proses formulir tambah produk yang dikirimkan
		var product entities.Product

		// Mendapatkan ID kategori dari formulir
		categoryId, err := strconv.Atoi(r.FormValue("category_id"))
		if err != nil {
			panic(err)
		}

		// Mendapatkan ISSN dari formulir
		issn, err := strconv.Atoi(r.FormValue("issn"))
		if err != nil {
			panic(err)
		}

		// Mengisi data produk dengan nilai dari formulir
		product.Name = r.FormValue("name")
		product.Category.Id = uint(categoryId)
		product.ISSN = int64(issn)
		product.Description = r.FormValue("description")
		product.CreatedAt = time.Now()
		product.UpdatedAt = time.Now()

		// Menangani unggah gambar produk
		file, handler, err := r.FormFile("image")
		if err == nil {
			// Menutup file setelah selesai
			defer file.Close()

			// Menggunakan handler.Filename untuk mendapatkan nama file asli
			originalFilename := handler.Filename
			fmt.Println("Original Filename:", originalFilename)

			// Menggunakan handler.Header.Get("Content-Type") untuk mendapatkan tipe konten
			contentType := handler.Header.Get("Content-Type")
			fmt.Println("Content Type:", contentType)

			// Membuat direktori uploads/images jika belum ada
			if err := os.MkdirAll("uploads/images", os.ModePerm); err != nil {
				panic(err)
			}

			// Menghasilkan timestamp
			timestamp := time.Now().UnixNano() / int64(time.Millisecond)

			// Menyimpan file yang diunggah ke direktori uploads/images
			imageName := fmt.Sprintf("%d_%s_image.jpg", timestamp, product.Name)
			imagePath := "uploads/images/" + imageName
			outFile, err := os.Create(imagePath)
			if err != nil {
				panic(err)
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, file)
			if err != nil {
				panic(err)
			}

			// Mengatur detail gambar dalam struktur produk
			product.ImageName = imageName
			product.ImagePath = imagePath
		}

		// Menyimpan informasi produk ke dalam database
		if ok := productmodel.Create(product); !ok {
			// Mengalihkan kembali jika gagal menyimpan ke database
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		// Mengalihkan ke halaman produk setelah berhasil menambahkan produk
		http.Redirect(w, r, "/products", http.StatusSeeOther)
	}
}

// Handler adalah tipe fungsi untuk menangani permintaan HTTP
type Handler func(http.ResponseWriter, *http.Request)

// AddHandler adalah variabel dengan tipe Handler yang merujuk ke fungsi Add
var AddHandler Handler = Add

// Detail menangani permintaan untuk halaman detail produk.
func Detail(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan ID produk dari parameter URL
	idString := r.URL.Query().Get("id")

	// Mengkonversi ID string menjadi tipe data integer
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	// Mendapatkan detail produk dari model berdasarkan ID
	product := productmodel.Detail(id)

	// Menyiapkan data untuk digunakan dalam template
	data := map[string]interface{}{
		"product": product,
	}

	// Menganalisis template HTML
	temp, err := template.ParseFiles("views/product/detail.html")
	if err != nil {
		panic(err)
	}

	// Menjalankan template dengan data yang telah disiapkan
	temp.Execute(w, data)
}

// Edit menangani permintaan untuk halaman edit produk.
func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Jika metode HTTP adalah GET, tampilkan formulir edit produk
		temp, err := template.ParseFiles("views/product/edit.html")
		if err != nil {
			panic(err)
		}

		// Mendapatkan ID produk dari parameter URL
		idString := r.URL.Query().Get("id")
		// Mengkonversi ID string menjadi tipe data integer
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		// Mendapatkan detail produk dari model berdasarkan ID
		product := productmodel.Detail(id)
		// Mendapatkan daftar kategori untuk ditampilkan dalam formulir
		categories := categorymodel.GetAll()

		// Menyiapkan data untuk digunakan dalam template
		data := map[string]interface{}{
			"product":    product,
			"categories": categories,
		}

		// Menjalankan template dengan data yang telah disiapkan
		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		// Jika metode HTTP adalah POST, proses formulir edit produk yang dikirimkan
		var product entities.Product

		// Mendapatkan ID produk dari formulir
		idString := r.FormValue("id")
		// Mengkonversi ID string menjadi tipe data integer
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		// Mendapatkan ID kategori dari formulir
		categoryId, err := strconv.Atoi(r.FormValue("category_id"))
		if err != nil {
			panic(err)
		}

		// Mendapatkan ISSN dari formulir
		issn, err := strconv.Atoi(r.FormValue("issn"))
		if err != nil {
			panic(err)
		}

		// Mengisi data produk dengan nilai dari formulir
		product.Name = r.FormValue("name")
		product.Category.Id = uint(categoryId)
		product.ISSN = int64(issn)
		product.Description = r.FormValue("description")
		product.UpdatedAt = time.Now()

		// Menangani unggah gambar produk
		file, handler, err := r.FormFile("image")
		if err == nil {
			// Menutup file setelah selesai
			defer file.Close()

			// Menggunakan handler.Filename untuk mendapatkan nama file asli
			originalFilename := handler.Filename
			fmt.Println("Original Filename:", originalFilename)

			// Menggunakan handler.Header.Get("Content-Type") untuk mendapatkan tipe konten
			contentType := handler.Header.Get("Content-Type")
			fmt.Println("Content Type:", contentType)

			// Membuat direktori uploads/images jika belum ada
			if err := os.MkdirAll("uploads/images", os.ModePerm); err != nil {
				panic(err)
			}

			// Menghasilkan timestamp
			timestamp := time.Now().UnixNano() / int64(time.Millisecond)

			// Menyimpan file yang diunggah ke direktori uploads/images
			imageName := fmt.Sprintf("%d_%s_image.jpg", timestamp, product.Name)
			imagePath := "uploads/images/" + imageName
			outFile, err := os.Create(imagePath)
			if err != nil {
				panic(err)
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, file)
			if err != nil {
				panic(err)
			}

			// Mengatur detail gambar dalam struktur produk
			product.ImageName = imageName
			product.ImagePath = imagePath
		}

		// Menyimpan informasi produk yang telah diedit ke dalam database
		if ok := productmodel.Update(id, product); !ok {
			// Mengalihkan kembali jika gagal menyimpan ke database
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		// Mengalihkan ke halaman produk setelah berhasil mengedit produk
		http.Redirect(w, r, "/products", http.StatusSeeOther)
	}
}

// Delete menangani permintaan untuk menghapus produk.
func Delete(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan ID produk dari parameter URL
	idString := r.URL.Query().Get("id")

	// Mengkonversi ID string menjadi tipe data integer
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	// Mendapatkan informasi produk sebelum dihapus
	product := productmodel.Detail(id)

	// Menghapus produk dari database
	if err := productmodel.Delete(id); err != nil {
		panic(err)
	}

	// Menghapus file gambar terkait, jika ada
	if product.ImagePath != "" {
		imagePath := filepath.Join("uploads", "images", product.ImageName)

		err := os.Remove(imagePath)
		if err != nil {
			// Mencatat kesalahan (bisa menggunakan log daripada fmt.Println)
			fmt.Println("Error deleting image file:", err)
		}
	}

	// Mengalihkan ke halaman produk setelah berhasil menghapus produk
	http.Redirect(w, r, "/products", http.StatusSeeOther)
}
