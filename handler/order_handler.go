package handler

import (
	"net/http"
	"strconv"

	"github.com/Iqbalhasanu/Msib-Hacktiv8-Golang-Assignment2/dto"
	"github.com/Iqbalhasanu/Msib-Hacktiv8-Golang-Assignment2/pkg/errs"
	"github.com/Iqbalhasanu/Msib-Hacktiv8-Golang-Assignment2/service"
	"github.com/gin-gonic/gin"
)

type orderHandler struct {
	orderService service.OrderService
}

func NewOrderHandler(orderService service.OrderService) orderHandler {
	return orderHandler{
		orderService: orderService,
	}
}

func (o *orderHandler) CreateOrder(ctx *gin.Context) {
	var orderRequestBody dto.NewOrderRequest

	if err := ctx.ShouldBindJSON(&orderRequestBody); err != nil {
		messageErr := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(messageErr.Status(), messageErr)
		return
	}

	newOrder, err := o.orderService.CreateOrder(orderRequestBody)

	if err != nil {
		ctx.JSON(err.Status(), err)
	}

	ctx.JSON(newOrder.StatusCode, newOrder)
}

func (o *orderHandler) GetAllOrder(ctx *gin.Context) {
	orders, err := o.orderService.GetAllOrder()

	if err != nil {
		ctx.JSON(err.Status(), err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Data":    orders,
		"message": "Success",
		"status":  http.StatusOK,
	})
}

func (o *orderHandler) DeleteOrder(ctx *gin.Context) {

	orderId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	if err := o.orderService.DeleteOrder(orderId); err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "order deleted"})
}

func (o *orderHandler) UpdateOrder(ctx *gin.Context) {
	var orderRequest dto.NewOrderRequest

	if err := ctx.ShouldBindJSON(&orderRequest); err != nil {
		ctx.JSON(422, gin.H{
			"errMessage": "unprocessible entity",
		})
		return
	}

	orderId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(400, gin.H{
			"errMessage": err.Error(),
		})
		return
	}

	updatedOrder, err := o.orderService.UpdateOrder(orderId, orderRequest)

	if err != nil {
		ctx.JSON(400, err)
		return
	}

	ctx.JSON(updatedOrder.Code, updatedOrder)
}
