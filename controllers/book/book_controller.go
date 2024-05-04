package book

import (
	"crud_app/config"
	"crud_app/helpers"
	"crud_app/models"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Index(writer http.ResponseWriter, request *http.Request) {
	var books []models.Book

	result := config.DB.Joins("Author").Order("created_at desc").Find(&books)
	if result.Error != nil {
		helpers.Response(writer, http.StatusInternalServerError, "Failed to get books", nil)
		return
	}

	helpers.Response(writer, http.StatusOK, "Success get books", books)
}

func Store(writer http.ResponseWriter, request *http.Request) {
	book := models.Book{}
	err := json.NewDecoder(request.Body).Decode(&book)
	if err != nil {
		helpers.Response(writer, http.StatusBadRequest, "Failed to decode request body", nil)
		return
	}
	defer request.Body.Close()

	author := config.DB.First(&models.Author{}, book.AuthorID)
	if errors.Is(author.Error, gorm.ErrRecordNotFound) {
		helpers.Response(writer, http.StatusNotFound, "Author not found", nil)
		return
	}

	result := config.DB.Omit("Author").Create(&book)
	if result.Error != nil {
		helpers.Response(writer, http.StatusInternalServerError, "Failed to create book", nil)
		return
	}

	bookResource := map[string]interface{}{
		"id":         book.ID,
		"title":      book.Title,
		"author_id":  book.AuthorID,
		"created_at": book.CreatedAt,
	}

	helpers.Response(writer, http.StatusCreated, "Success create book", bookResource)
}

func Update(writer http.ResponseWriter, request *http.Request) {
	idParams := mux.Vars(request)["id"]
	id, _ := strconv.Atoi(idParams)
	book := models.Book{}

	checkBook := config.DB.Where("id = ?", id).First(&book)
	if errors.Is(checkBook.Error, gorm.ErrRecordNotFound) {
		helpers.Response(writer, http.StatusNotFound, "Book not found", nil)
		return
	}

	err := json.NewDecoder(request.Body).Decode(&book)
	if err != nil {
		helpers.Response(writer, http.StatusBadRequest, "Failed to decode request body", nil)
		return
	}
	defer request.Body.Close()

	updateBook := config.DB.Where("id = ?", id).Updates(&book)
	if updateBook.Error != nil {
		helpers.Response(writer, http.StatusInternalServerError, "Failed to update book", nil)
		return
	}

	bookResource := map[string]interface{}{
		"id":         book.ID,
		"title":      book.Title,
		"author_id":  book.AuthorID,
		"updated_at": book.UpdatedAt,
	}

	helpers.Response(writer, http.StatusOK, "Success update book", bookResource)
}

func Show(writer http.ResponseWriter, request *http.Request) {
	idParams := mux.Vars(request)["id"]
	id, _ := strconv.Atoi(idParams)

	book := models.Book{}

	result := config.DB.Joins("Author").First(&book, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		helpers.Response(writer, http.StatusNotFound, "Book not found", nil)
		return
	}

	helpers.Response(writer, http.StatusOK, "Success get book", book)
}

func Destroy(writer http.ResponseWriter, request *http.Request) {
	idParams := mux.Vars(request)["id"]
	id, _ := strconv.Atoi(idParams)

	deleteBook := config.DB.Delete(&models.Book{}, id)
	if deleteBook.RowsAffected == 0 {
		helpers.Response(writer, http.StatusNotFound, "Book not found", nil)
		return
	}

	if deleteBook.Error != nil {
		helpers.Response(writer, http.StatusInternalServerError, "Failed to delete book", nil)
		return
	}

	helpers.Response(writer, http.StatusOK, "Success delete book", nil)
}
