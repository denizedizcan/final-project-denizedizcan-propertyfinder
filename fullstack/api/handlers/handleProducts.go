package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/denizedizcan/final-project-denizedizcan-propertyfinder/api/models"
	"github.com/denizedizcan/final-project-denizedizcan-propertyfinder/api/responses"
)

func (h handler) GetProducts(w http.ResponseWriter, r *http.Request) {

	var products []models.Product
	products, err := models.FindAllProducts(h.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, &products)
}

func (h handler) AddProducts(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	var products []models.Product
	err = json.Unmarshal(body, &products)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	if err := models.InsertProducts(h.DB); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, &products)
}
