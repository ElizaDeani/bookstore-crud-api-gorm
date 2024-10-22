package main

import (
	"log"
	"net/http"

	"post-test-mikti/pkg/config"
	"post-test-mikti/pkg/routes/bookstoreroutes"

	"github.com/gorilla/mux"
)

func main() {
	// Inisialisasi database
	db, err := config.SetupDatabase()
	if err != nil {
		log.Fatalf("Gagal menghubungkan ke database: %v", err)
	}
	defer db.Close()

	// Inisialisasi router
	router := mux.NewRouter()

	// Inisialisasi rute buku
	bookstoreroutes.RegisterBookRoutes(router, db)

	// Jalankan server
	log.Println("Server berjalan di port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
