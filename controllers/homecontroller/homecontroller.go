// Back End dibuat oleh Yoga Ardiansyah

// controllers/homecontroller/homecontroller.go
package homecontroller

import (
	"net/http"
	"text/template"
)

// Welcome menangani permintaan untuk halaman utama
func Welcome(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "views/home/index.html", nil)
}

// Fungsi helper untuk merender template
func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	temp, err := template.ParseFiles(tmpl)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = temp.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
