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

func (h handler) AddOrder(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	var basket models.Basket
	err = json.Unmarshal(body, &basket)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	if basket.UserID == 0 || basket.Value == 0 {
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("userID or order value not found"))
		return
	}
	var order models.Order

	order.UserID = basket.UserID
	order.Value = basket.Value
	ordernumber, err := order.CreateOrder(h.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	if ordernumber == 0 {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("order not created"))
		return
	}
	// check nand drop stock
	lenght := len(basket.BasketItems)
	if lenght == 0 {
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("basket items not found"))
		if err := order.DeleteOrder(h.DB); err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		return
	}

	for i := 0; i < lenght; i++ {
		if err := basket.BasketItems[i].CheckStock(h.DB); err != nil {
			responses.ERROR(w, http.StatusNotFound, err)
			if err := order.DeleteOrder(h.DB); err != nil {
				responses.ERROR(w, http.StatusInternalServerError, err)
				return
			}
			return

		}
		if err := basket.BasketItems[i].DropStock(h.DB); err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			if err := order.DeleteOrder(h.DB); err != nil {
				responses.ERROR(w, http.StatusInternalServerError, err)
				return
			}
			return
		}
	}

	var order_items []models.OrderItems
	var order_item models.OrderItems
	for i := 0; i < lenght; i++ {
		order_item.OrderNumber = ordernumber
		order_item.Quantity = basket.BasketItems[i].Quantity
		order_item.Sku = basket.BasketItems[i].Sku
		order_item.Value = basket.BasketItems[i].Value

		order_items = append(order_items, order_item)
	}

	fmt.Println(order_items)
	if err := models.CreateOrderItems(h.DB, order_items); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		if err := order.DeleteOrder(h.DB); err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		return
	}

	var orders *models.Order
	orders, err = order.FindOrder(h.DB)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		if err := order.DeleteOrder(h.DB); err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		return
	}
	//free basket and items
	for i := 0; i < lenght; i++ {
		if err := basket.BasketItems[i].DeleteBasketItem(h.DB); err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			if err := order.DeleteOrder(h.DB); err != nil {
				responses.ERROR(w, http.StatusInternalServerError, err)
				return
			}
			return
		}
	}

	if err := basket.UpdateBasketValueAfterOrder(h.DB); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		if err := order.DeleteOrder(h.DB); err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		return
	}

	responses.JSON(w, http.StatusOK, orders)

}

func (h handler) ListOrder(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	var order *models.Order
	err = json.Unmarshal(body, &order)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	order, err = order.FindOrder(h.DB)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}

	responses.JSON(w, http.StatusOK, order)
}

func (h handler) AddOldOrder(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	var basket models.Basket
	err = json.Unmarshal(body, &basket)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	if basket.UserID == 0 || basket.Value == 0 {
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("userID or order value not found"))
		return
	}
	var order models.Order

	order.UserID = basket.UserID
	order.Value = basket.Value
	ordernumber, err := order.CreateOrderOld(h.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	if ordernumber == 0 {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("order not created"))
		return
	}
	// check nand drop stock
	lenght := len(basket.BasketItems)
	if lenght == 0 {
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("basket items not found"))
		if err := order.DeleteOrder(h.DB); err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		return
	}

	for i := 0; i < lenght; i++ {
		if err := basket.BasketItems[i].CheckStock(h.DB); err != nil {
			responses.ERROR(w, http.StatusNotFound, err)
			if err := order.DeleteOrder(h.DB); err != nil {
				responses.ERROR(w, http.StatusInternalServerError, err)
				return
			}
			return

		}
		if err := basket.BasketItems[i].DropStock(h.DB); err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			if err := order.DeleteOrder(h.DB); err != nil {
				responses.ERROR(w, http.StatusInternalServerError, err)
				return
			}
			return
		}
	}

	var order_items []models.OrderItems
	var order_item models.OrderItems
	for i := 0; i < lenght; i++ {
		order_item.OrderNumber = ordernumber
		order_item.Quantity = basket.BasketItems[i].Quantity
		order_item.Sku = basket.BasketItems[i].Sku
		order_item.Value = basket.BasketItems[i].Value

		order_items = append(order_items, order_item)
	}

	fmt.Println(order_items)
	if err := models.CreateOrderItems(h.DB, order_items); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		if err := order.DeleteOrder(h.DB); err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		return
	}

	var orders *models.Order
	orders, err = order.FindOrder(h.DB)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		if err := order.DeleteOrder(h.DB); err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		return
	}
	//free basket and items
	for i := 0; i < lenght; i++ {
		if err := basket.BasketItems[i].DeleteBasketItem(h.DB); err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			if err := order.DeleteOrder(h.DB); err != nil {
				responses.ERROR(w, http.StatusInternalServerError, err)
				return
			}
			return
		}
	}

	if err := basket.UpdateBasketValueAfterOrder(h.DB); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		if err := order.DeleteOrder(h.DB); err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		return
	}

	responses.JSON(w, http.StatusOK, orders)

}
