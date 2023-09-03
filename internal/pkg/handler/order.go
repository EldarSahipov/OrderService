package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getOrderByUID(c *gin.Context) {
	uid := c.Param("uid")

	order, err := h.service.Orders.GetByUID(uid)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, order)
}

func (h *Handler) getAll(c *gin.Context) {
	orders, err := h.service.Orders.GetAll()
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (h *Handler) GetOrderByUIDFromCache(c *gin.Context) {
	uid := c.Param("uid")

	order, err := h.service.GetByUID(uid)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, order)
}
