package handler

import (
	_ "OrderService/docs"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Get All Orders
// @Tags orders
// @Description get all orders
// @ID get-all-orders
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.Order
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/orders/all [get]
func (h *Handler) getAll(c *gin.Context) {
	orders, err := h.service.Orders.GetAll()
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, orders)
}

// @Summary Get Order By Uid From Cache
// @Tags orders
// @Description Get order by uid from cache
// @ID get-order-by-uid-from-cache
// @Accept json
// @Produce json
// @Param        uid  path      string        true "order's uid" minlength(19)  maxlength(36)
// @Success      200  {object}  models.Order
// @Failure      400  {object}  errorResponse
// @Failure      404  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /api/orders/{uid} [get]
func (h *Handler) GetOrderByUIDFromCache(c *gin.Context) {
	uid := c.Param("uid")

	order, err := h.service.GetByUID(uid)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, order)
}
