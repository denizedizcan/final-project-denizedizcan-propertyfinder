//basket handle add update remove basket item and calculate the value of basket(calcute in different pakage)

package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/denizedizcan/final-project-denizedizcan-propertyfinder/api/models"
	"github.com/denizedizcan/final-project-denizedizcan-propertyfinder/api/responses"
)

// get basket of the user
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

// add basket item to basket
func (h handler) AddBasketItem(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	// get the basket item from body
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
	// insert basket item
	if err := basket_item.InsertOneBasketItem(h.DB); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	// find price and update basket value find stock
	if err := basket_item.UpdateBasketItemsValue(h.DB); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	//updateBasket value
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
	// check for discount
	discount := basket.FindDiscount(h.DB)
	basket.Value -= discount

	//update after discount
	if err := basket.UpdateBasketValue(h.DB); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	basket, err = basket_item.FindUserBasketbyBasketitem(h.DB)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, &basket)
}

// update the basket item quantity
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
	// update the basket item
	if err := basket_item.UpdateBasketItem(h.DB); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	// find price and update basket value find stock
	if err := basket_item.UpdateBasketItemsValue(h.DB); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	//updateBasket value
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
	// check for discount
	discount := basket.FindDiscount(h.DB)
	basket.Value -= discount

	//update after discount
	if err := basket.UpdateBasketValue(h.DB); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	basket, err = basket_item.FindUserBasketbyBasketitem(h.DB)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, &basket)
}

// delete one item from basket
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
	// after delete check for discount
	discount := basket.FindDiscount(h.DB)
	basket.Value -= discount

	//update after discount
	if err := basket.UpdateBasketValue(h.DB); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	basket, err = basket_item.FindUserBasketbyBasketitem(h.DB)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, &basket)

}
