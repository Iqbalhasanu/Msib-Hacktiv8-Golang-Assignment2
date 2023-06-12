package handler

import (
	"github.com/Iqbalhasanu/Msib-Hacktiv8-Golang-Assignment2.git/dto"
	"github.com/Iqbalhasanu/Msib-Hacktiv8-Golang-Assignment2.git/pkg/errs"
	"github.com/Iqbalhasanu/Msib-Hacktiv8-Golang-Assignment2.git/pkg/helpers"
	"github.com/Iqbalhasanu/Msib-Hacktiv8-Golang-Assignment2.git/service"
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

// CreateNewOrder godoc
// @Summary Create Order
// @Description Create a new order
// @Tags order
// @ID create-new-order
// @Accept json
// @Produce json
// @Param RequestBody body dto.NewOrderRequest true "request body json"
// @Success 201 {object} dto.NewOrderResponse
// @Router /orders [post]
func (o *orderHandler) CreateOrder(ctx *gin.Context) {
	var orderRequest dto.NewOrderRequest

	if err := ctx.ShouldBindJSON(&orderRequest); err != nil {
		error := errs.UnprocessableEntity("invalid request body")
		ctx.JSON(error.Status(), error)
		return
	}

	result, err := o.orderService.CreateOrder(orderRequest)
	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(result.StatusCode, result)
}

// GetAllOrders godoc
// @Summary Get All Orders
// @Description Get details of all orders
// @Tags order
// @ID get-all-orders
// @Accept json
// @Produce json
// @Success 200 {object} dto.AllOrderResponse
// @Router /orders [get]
func (o *orderHandler) GetAllOrders(ctx *gin.Context) {
	result, err := o.orderService.GetAllOrder()
	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)
}

// GetDetailOrder godoc
// @Summary Get Details Order
// @Description Get details an order
// @Tags order
// @ID get-order
// @Accept json
// @Produce json
// @Param OrderId path int true "Id of the order"
// @Success 200 {object} dto.NewOrderResponse
// @Router /orders/{OrderId} [get]
func (o *orderHandler) GetOrder(ctx *gin.Context) {
	orderId, err := helpers.GetParam(ctx, "orderId")

	result, err := o.orderService.GetOrderById(orderId)
	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)
}

// UpdateOrder godoc
// @Summary Update Order
// @Description Update an order
// @Tags order
// @ID update-order
// @Accept json
// @Produce json
// @Param OrderId path int true "Id of the order"
// @Param RequestBody body dto.NewOrderRequest true "request body json"
// @Success 200 {object} dto.NewOrderResponse
// @Router /orders/{OrderId} [put]
func (o *orderHandler) UpdateOrder(ctx *gin.Context) {
	var orderRequest dto.NewOrderRequest

	if err := ctx.ShouldBindJSON(&orderRequest); err != nil {
		error := errs.UnprocessableEntity("invalid request body")
		ctx.JSON(error.Status(), error)
		return
	}

	orderId, err := helpers.GetParam(ctx, "orderId")
	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	result, err := o.orderService.UpdateOrder(orderId, orderRequest)

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(result.StatusCode, result)
}

// DeleteOrder godoc
// @Summary Delete Order
// @Description Delete an Order
// @Tags order
// @ID delete-order
// @Accept json
// @Produce json
// @Param OrderId path int true "Id of the order"
// @Success 200 {object} dto.DeleteOrderResponse
// @Router /orders/{OrderId} [delete]
func (o *orderHandler) DeleteOrder(ctx *gin.Context) {
	orderId, err := helpers.GetParam(ctx, "orderId")
	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	result, err := o.orderService.DeleteOrder(orderId)
	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(result.StatusCode, result)
}
