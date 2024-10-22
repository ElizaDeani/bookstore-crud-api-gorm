package bookcontrollers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"post-test-mikti/pkg/models"
	"post-test-mikti/pkg/utils"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// GetAllBooks menangani permintaan untuk mendapatkan semua buku
func GetAllBooks(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var books []models.Book
		result := db.Find(&books)
		if result.Error != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Gagal mengambil data buku")
			return
		}
		utils.RespondWithJSON(w, http.StatusOK, books)
	}
}

// GetBookByID menangani permintaan untuk mendapatkan buku berdasarkan ID
func GetBookByID(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "ID tidak valid")
			return
		}

		var book models.Book
		result := db.First(&book, id)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				utils.RespondWithError(w, http.StatusNotFound, "Buku tidak ditemukan")
			} else {
				utils.RespondWithError(w, http.StatusInternalServerError, "Gagal mengambil data buku")
			}
			return
		}
		utils.RespondWithJSON(w, http.StatusOK, book)
	}
}

// CreateBook menangani permintaan untuk membuat buku baru
func CreateBook(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&book); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Data tidak valid")
			return
		}
		defer r.Body.Close()

		result := db.Create(&book)
		if result.Error != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Gagal membuat buku")
			return
		}
		utils.RespondWithJSON(w, http.StatusCreated, book)
	}
}

// UpdateBook menangani permintaan untuk memperbarui buku
func UpdateBook(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "ID tidak valid")
			return
		}

		var book models.Book
		result := db.First(&book, id)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				utils.RespondWithError(w, http.StatusNotFound, "Buku tidak ditemukan")
			} else {
				utils.RespondWithError(w, http.StatusInternalServerError, "Gagal mengambil data buku")
			}
			return
		}

		var updatedData models.Book
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&updatedData); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Data tidak valid")
			return
		}
		defer r.Body.Close()

		book.Title = updatedData.Title
		book.Author = updatedData.Author
		book.Description = updatedData.Description
		book.Price = updatedData.Price

		result = db.Save(&book)
		if result.Error != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Gagal memperbarui buku")
			return
		}
		utils.RespondWithJSON(w, http.StatusOK, book)
	}
}

// DeleteBook menangani permintaan untuk menghapus buku
func DeleteBook(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "ID tidak valid")
			return
		}

		var book models.Book
		result := db.First(&book, id)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				utils.RespondWithError(w, http.StatusNotFound, "Buku tidak ditemukan")
			} else {
				utils.RespondWithError(w, http.StatusInternalServerError, "Gagal mengambil data buku")
			}
			return
		}

		result = db.Delete(&book)
		if result.Error != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Gagal menghapus buku")
			return
		}
		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Buku berhasil dihapus"})
	}
}
