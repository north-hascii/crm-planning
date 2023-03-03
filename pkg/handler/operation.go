package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/north-hascii/crm-planning/models"
	"github.com/siruspen/logrus"
	"net/http"
	"strconv"
)

func (h *Handler) getOperationById(c *gin.Context) {
	operationId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	operation, err := h.services.Operation.GetOperationById(operationId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, operation)
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

	id, err := h.services.CreateOperation(&input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) updateOperation(c *gin.Context) {
	//operationId, err := getOperationId(c)
	//if err != nil {
	//	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	//	return
	//}
	id, ok := c.Get("operation_name")
	if !ok {
		logrus.Printf("Level: repos; func getOperationId(): models.Operation=%v", id)
		newErrorResponse(c, http.StatusInternalServerError, "operation id not found")
		return
	}
	logrus.Printf("NEXT=%v", id)

	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "operation id is of invalid type")
		return
	}

	var operation models.Operation
	if err := c.BindJSON(&operation); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Operation.UpdateOperation(&operation); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"request": "update operation" + strconv.Itoa(idInt),
	})
}

func (h *Handler) deleteOperation(c *gin.Context) {
	//userId, err := getUserId(c)
	//if err != nil {
	//	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	//	return
	//}
	//
	//cartItemId, err := strconv.Atoi(c.Param("id"))
	//if err != nil {
	//	newErrorResponse(c, http.StatusBadRequest, "invalid id param")
	//	return
	//}
	//
	//err = h.services.Cart.DeleteCartItemsById(userId, cartItemId)
	//if err != nil {
	//	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	//	return
	//}
	//
	//c.JSON(http.StatusOK, map[string]interface{}{
	//	"request": "delete cart item",
	//})
}
