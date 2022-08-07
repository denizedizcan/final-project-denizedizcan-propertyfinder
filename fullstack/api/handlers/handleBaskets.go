//basket handle add update remove basket item and calculate the value of basket(calcute in different pakage)

package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/denizedizcan/final-project-denizedizcan-propertyfinder/api/models"
	"github.com/denizedizcan/final-project-denizedizcan-propertyfinder/api/responses"
)

func (h handler) GetBasket(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	err = json.Unmarshal(body, &user)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	if user.Email == "" {
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("email not found"))
		return
	}
	basket, err := user.FindUserBasketbyUser(h.DB)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}

	responses.JSON(w, http.StatusOK, &basket)
}

func (h handler) AddBasketItem(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	var basket_item models.BasketItems
	err = json.Unmarshal(body, &basket_item)

	fmt.Println(basket_item)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	if basket_item.BasketID == 0 || basket_item.Sku == 0 {
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("BasketID or Sku not found"))
		return
	}
	// check stock

	if err := basket_item.CheckStock(h.DB); err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	fmt.Println("stock checked")
	if err := basket_item.InsertOneBasketItem(h.DB); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	fmt.Println("insert basket item")
	// find price and update basket value find stock
	if err := basket_item.UpdateBasketItemsValue(h.DB); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	fmt.Println("update basketitem value")

	//updateBasket
	if err := basket_item.UpdateBasketValue(h.DB); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	fmt.Println("update basket value")

	var basket models.Basket

	basket, err = basket_item.FindUserBasketbyBasketitem(h.DB)
	fmt.Println("find user")

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, &basket)
}

func (h handler) UpdateBasketItem(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	var basket_item models.BasketItems
	err = json.Unmarshal(body, &basket_item)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	if basket_item.BasketID == 0 || basket_item.Sku == 0 {
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("BasketID or Sku not found"))
		return
	}
	// check stock

	if err := basket_item.CheckStock(h.DB); err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}

	if err := basket_item.UpdateBasketItem(h.DB); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	// find price and update basket value find stock
	if err := basket_item.UpdateBasketItemsValue(h.DB); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	//updateBasket
	if err := basket_item.UpdateBasketValue(h.DB); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	var basket models.Basket

	basket, err = basket_item.FindUserBasketbyBasketitem(h.DB)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, &basket)
}

func (h handler) DeleteOneItem(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	var basket_item models.BasketItems
	err = json.Unmarshal(body, &basket_item)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	if basket_item.BasketID == 0 || basket_item.Sku == 0 {
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("BasketID or Sku not found"))
		return
	}

	if err := basket_item.DeleteBasketItem(h.DB); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	// find price and update basket value find stock
	if err := basket_item.UpdateBasketItemsValue(h.DB); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	//updateBasket
	if err := basket_item.UpdateBasketValue(h.DB); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	var basket models.Basket

	basket, err = basket_item.FindUserBasketbyBasketitem(h.DB)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, &basket)

}
