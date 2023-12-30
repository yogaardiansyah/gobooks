// Back End dibuat oleh Yoga Ardiansyah

// controllers/categorycontroller/categorycontroller.go
package categorycontroller

import (
	"go-web-native/entities"
	"go-web-native/models/categorymodel"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan semua kategori dari model
	categories := categorymodel.GetAll()

	// Membuat map data untuk disertakan dalam template
	data := map[string]interface{}{
		"categories": categories,
	}

	// Parsing dan mengeksekusi template
	renderTemplate(w, "views/category/index.html", data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Menampilkan halaman form untuk menambahkan kategori
		renderTemplate(w, "views/category/create.html", nil)
	}

	if r.Method == http.MethodPost {
		// Menangani penambahan kategori
		handleCategoryAdd(w, r)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Menampilkan halaman form untuk mengedit kategori
		handleCategoryEditGet(w, r)
	}

	if r.Method == http.MethodPost {
		// Menangani pengiriman form edit kategori
		handleCategoryEditPost(w, r)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	// Menghapus kategori
	handleCategoryDelete(w, r)
}

// Fungsi helper untuk merender template
func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	temp, err := template.ParseFiles(tmpl)
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

// Fungsi helper untuk menangani penambahan kategori
func handleCategoryAdd(w http.ResponseWriter, r *http.Request) {
	var category entities.Category

	category.Name = r.FormValue("name")
	category.CreatedAt = time.Now()
	category.UpddatedAt = time.Now()

	// Menambahkan kategori
	if ok := categorymodel.Create(category); !ok {
		renderTemplate(w, "views/category/create.html", nil)
		return
	}

	// Redirect ke halaman kategori setelah penambahan berhasil
	http.Redirect(w, r, "/categories", http.StatusSeeOther)
}

// Fungsi helper untuk menangani pengambilan data kategori untuk form edit
func handleCategoryEditGet(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/category/edit.html")
	if err != nil {
		panic(err)
	}

	// Mengambil ID kategori dari parameter URL
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	// Mendapatkan detail kategori berdasarkan ID
	category := categorymodel.Detail(id)
	data := map[string]interface{}{
		"category": category,
	}

	temp.Execute(w, data)
}

// Fungsi helper untuk menangani pengiriman form edit kategori
func handleCategoryEditPost(w http.ResponseWriter, r *http.Request) {
	var category entities.Category

	// Mendapatkan data dari form
	idString := r.FormValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	category.Name = r.FormValue("name")
	category.UpddatedAt = time.Now()

	// Memperbarui kategori berdasarkan ID
	if ok := categorymodel.Update(id, category); !ok {
		// Jika gagal, redirect ke halaman sebelumnya
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
		return
	}

	// Redirect ke halaman kategori setelah perubahan berhasil
	http.Redirect(w, r, "/categories", http.StatusSeeOther)
}

// Fungsi helper untuk menangani penghapusan kategori
func handleCategoryDelete(w http.ResponseWriter, r *http.Request) {
	// Mengambil ID kategori dari parameter URL
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	// Menghapus kategori berdasarkan ID
	if err := categorymodel.Delete(id); err != nil {
		panic(err)
	}

	// Redirect ke halaman kategori setelah penghapusan berhasil
	http.Redirect(w, r, "/categories", http.StatusSeeOther)
}
