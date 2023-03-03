package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/north-hascii/crm-planning/models"
	"net/http"
)

func (h *Handler) getOperation(context *gin.Context) {

}

func (h *Handler) getAllOperation(c *gin.Context) {
	operations, err := h.services.GetAllOperations()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, operations)
}

func (h *Handler) createOperation(c *gin.Context) {
	var input models.Operation

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateOperation(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) updateOperation(c *gin.Context) {

}

func (h *Handler) deleteOperation(c *gin.Context) {

}
