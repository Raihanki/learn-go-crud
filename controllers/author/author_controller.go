package author

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
	var authors []models.Author

	results := config.DB.Find(&authors)
	if results.Error != nil {
		helpers.Response(writer, http.StatusInternalServerError, results.Error.Error(), nil)
		return
	}

	helpers.Response(writer, http.StatusOK, "Successfully get authors", authors)
}

func Store(writer http.ResponseWriter, request *http.Request) {
	author := models.Author{}

	err := json.NewDecoder(request.Body).Decode(&author)
	if err != nil {
		helpers.Response(writer, http.StatusBadRequest, err.Error(), nil)
		return
	}
	defer request.Body.Close()

	newAuthor := config.DB.Create(&author)
	if newAuthor.Error != nil {
		helpers.Response(writer, http.StatusInternalServerError, newAuthor.Error.Error(), nil)
		return
	}

	helpers.Response(writer, http.StatusCreated, "Successfully create author", author)
}

func Update(writer http.ResponseWriter, request *http.Request) {
	idParams := mux.Vars(request)["id"]
	id, _ := strconv.Atoi(idParams)
	author := models.Author{}

	getAuthor := config.DB.First(&author, id)
	if errors.Is(getAuthor.Error, gorm.ErrRecordNotFound) {
		helpers.Response(writer, http.StatusNotFound, "Author not found", nil)
		return
	}

	if getAuthor.Error != nil {
		helpers.Response(writer, http.StatusInternalServerError, getAuthor.Error.Error(), nil)
		return
	}

	err := json.NewDecoder(request.Body).Decode(&author)
	if err != nil {
		helpers.Response(writer, http.StatusBadRequest, err.Error(), nil)
		return
	}
	defer request.Body.Close()

	result := config.DB.Where("id = ?", author.ID).Updates(&author)
	if result.Error != nil {
		helpers.Response(writer, http.StatusInternalServerError, result.Error.Error(), nil)
		return
	}

	helpers.Response(writer, http.StatusOK, "Successfully update author", author)
}

func Show(writer http.ResponseWriter, request *http.Request) {
	idParams := mux.Vars(request)["id"]
	id, _ := strconv.Atoi(idParams)
	author := models.Author{}

	getAuthor := config.DB.First(&author, id)
	if errors.Is(getAuthor.Error, gorm.ErrRecordNotFound) {
		helpers.Response(writer, http.StatusNotFound, "Author not found", nil)
		return
	}

	if getAuthor.Error != nil {
		helpers.Response(writer, http.StatusInternalServerError, getAuthor.Error.Error(), nil)
		return
	}

	helpers.Response(writer, http.StatusOK, "Successfully get author", author)
}

func Destroy(writer http.ResponseWriter, request *http.Request) {
	idParams := mux.Vars(request)["id"]
	id, _ := strconv.Atoi(idParams)

	result := config.DB.Where("id = ?", id).Delete(&models.Author{})
	if result.RowsAffected == 0 {
		helpers.Response(writer, http.StatusInternalServerError, "Author Not Found", nil)
		return
	}

	if result.Error != nil {
		helpers.Response(writer, http.StatusInternalServerError, result.Error.Error(), nil)
		return
	}

	helpers.Response(writer, http.StatusOK, "Successfully delete author", nil)
}
