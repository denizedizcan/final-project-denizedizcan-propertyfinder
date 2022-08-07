package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/denizedizcan/final-project-denizedizcan-propertyfinder/api/models"
	"github.com/denizedizcan/final-project-denizedizcan-propertyfinder/api/responses"
)

func (h handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	var user models.User
	err = json.Unmarshal(body, &user)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err := user.Validate("Create"); err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user.Prepare()
	if err := user.SaveUser(h.DB); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	var basket models.Basket
	basket.UserID = user.UserID

	if err := basket.InsertBasket(h.DB); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	if err := user.FindUser(h.DB); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, &user)
}

func (h handler) Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	var user models.User
	err = json.Unmarshal(body, &user)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err := user.Validate("Login"); err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()

	if err := user.LoginUser(h.DB); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, &user)
}
