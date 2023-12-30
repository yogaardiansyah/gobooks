// Back End dibuat oleh Yoga Ardiansyah

// main.go
package main

import (
	"go-web-native/config"
	"go-web-native/controllers/categorycontroller"
	"go-web-native/controllers/homecontroller"
	"go-web-native/controllers/productcontroller"
	"log"
	"net/http"
)

func main() {
	// Koneksi ke database
	config.ConnectDB()

	// Konfigurasi rute
	setupRoutes()

	// Layanan file statis
	serveStaticFiles()

	// Jalankan server
	addr := "localhost:8080"
	log.Printf("Server berjalan di http://%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func setupRoutes() {
	// Rute halaman utama
	http.HandleFunc("/", homecontroller.Welcome)

	// Rute kategori
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	// Rute produk
	http.HandleFunc("/products", productcontroller.Index)
	http.HandleFunc("/products/add", productcontroller.Add)
	http.HandleFunc("/products/detail", productcontroller.Detail)
	http.HandleFunc("/products/edit", productcontroller.Edit)
	http.HandleFunc("/products/delete", productcontroller.Delete)
}

func serveStaticFiles() {
	// Layanan unggah
	http.Handle("/uploads/images/", http.StripPrefix("/uploads/images/", http.FileServer(http.Dir("./uploads/images"))))
}
