package service

import (
	"net/http"

	"github.com/Iqbalhasanu/Msib-Hacktiv8-Golang-Assignment2.git/dto"
	"github.com/Iqbalhasanu/Msib-Hacktiv8-Golang-Assignment2.git/entity"
	"github.com/Iqbalhasanu/Msib-Hacktiv8-Golang-Assignment2.git/pkg/errs"
	"github.com/Iqbalhasanu/Msib-Hacktiv8-Golang-Assignment2.git/repository/order_repository"
)

type orderService struct {
	orderRepo   order_repository.OrderRepository
	itemService ItemService
}

type OrderService interface {
	CreateOrder(payload dto.NewOrderRequest) (*dto.NewOrderResponse, errs.MessageErr)
	GetAllOrder() (*dto.AllOrderResponse, errs.MessageErr)
	GetOrderById(int) (*dto.NewOrderResponse, errs.MessageErr)
	UpdateOrder(orderId int, payload dto.NewOrderRequest) (*dto.NewOrderResponse, errs.MessageErr)
	DeleteOrder(orderId int) (*dto.DeleteOrderResponse, errs.MessageErr)
}

func NewOrderService(orderRepo order_repository.OrderRepository, itemService ItemService) OrderService {
	return &orderService{
		orderRepo:   orderRepo,
		itemService: itemService,
	}
}

func (o *orderService) CreateOrder(payload dto.NewOrderRequest) (*dto.NewOrderResponse, errs.MessageErr) {
	orderPayload := entity.Orders{
		OrderedAt:    payload.OrderedAt,
		CustomerName: payload.CustomerName,
	}

	itemsPayload := []entity.Items{}

	for _, eachItem := range payload.Items {
		item := entity.Items{
			ItemCode:    eachItem.ItemCode,
			Quantity:    eachItem.Quantity,
			Description: eachItem.Description,
		}

		itemsPayload = append(itemsPayload, item)
	}

	newOrder, err := o.orderRepo.CreateOrder(orderPayload, itemsPayload)
	if err != nil {
		return nil, err
	}

	var newItem []dto.ItemResponse

	for _, eachItem := range newOrder.Items {
		item := dto.ItemResponse{
			ItemId:      eachItem.ItemId,
			ItemCode:    eachItem.ItemCode,
			Description: eachItem.Description,
			Quantity:    eachItem.Quantity,
			OrderId:     eachItem.OrderId,
		}

		newItem = append(newItem, item)
	}

	response := &dto.NewOrderResponse{
		Message: "order created successfully",
		Data: dto.OrderResponse{
			OrderId:      newOrder.OrderId,
			CreatedAt:    newOrder.CreatedAt,
			UpdatedAt:    newOrder.UpdatedAt,
			OrderedAt:    newOrder.OrderedAt,
			CustomerName: newOrder.CustomerName,
			Items:        newItem,
		},
		StatusCode: http.StatusCreated,
	}

	return response, nil
}

func (o *orderService) GetAllOrder() (*dto.AllOrderResponse, errs.MessageErr) {
	allOrder, err := o.orderRepo.GetAllOrder()
	if err != nil {
		return nil, err
	}

	var orders []dto.OrderResponse
	for _, order := range allOrder {
		var items []dto.ItemResponse
		for _, item := range order.Items {
			i := dto.ItemResponse{
				ItemId:      item.ItemId,
				ItemCode:    item.ItemCode,
				Description: item.Description,
				Quantity:    item.Quantity,
				OrderId:     item.OrderId,
				CreatedAt:   item.CreatedAt,
				UpdatedAt:   item.UpdatedAt,
			}

			items = append(items, i)
		}

		o := dto.OrderResponse{
			OrderId:      order.OrderId,
			CreatedAt:    order.CreatedAt,
			UpdatedAt:    order.CreatedAt,
			CustomerName: order.CustomerName,
			Items:        items,
		}

		orders = append(orders, o)
	}

	response := &dto.AllOrderResponse{
		Message:    "all order found",
		Data:       orders,
		StatusCode: http.StatusOK,
	}

	return response, nil
}

func (o *orderService) GetOrderById(param int) (*dto.NewOrderResponse, errs.MessageErr) {
	order, err := o.orderRepo.GetOrderById(param)
	if err != nil {
		return nil, err
	}

	var items []dto.ItemResponse
	for _, eachItem := range order.Items {
		item := dto.ItemResponse{
			ItemId:      eachItem.ItemId,
			ItemCode:    eachItem.ItemCode,
			Description: eachItem.Description,
			Quantity:    eachItem.Quantity,
			OrderId:     eachItem.OrderId,
		}

		items = append(items, item)
	}

	orderResponse := dto.OrderResponse{
		OrderId:      order.OrderId,
		CreatedAt:    order.CreatedAt,
		UpdatedAt:    order.UpdatedAt,
		OrderedAt:    order.OrderedAt,
		CustomerName: order.CustomerName,
		Items:        items,
	}

	response := &dto.NewOrderResponse{
		Message:    "order found",
		Data:       orderResponse,
		StatusCode: http.StatusOK,
	}
	return response, nil
}

func (o *orderService) UpdateOrder(orderId int, payload dto.NewOrderRequest) (*dto.NewOrderResponse, errs.MessageErr) {
	itemCodes := payload.ItemsToItemCode()
	itemsPayload := []entity.Items{}
	if itemCodes != nil {
		_, err := o.itemService.FindItemsByItemCode(itemCodes, orderId)
		if err != nil {
			return nil, err
		}

		for _, eachItem := range payload.Items {
			item := entity.Items{
				ItemCode:    eachItem.ItemCode,
				Description: eachItem.Description,
				Quantity:    eachItem.Quantity,
			}

			itemsPayload = append(itemsPayload, item)
		}
	}

	orderPayload := entity.Orders{
		OrderId:      orderId,
		OrderedAt:    payload.OrderedAt,
		CustomerName: payload.CustomerName,
	}

	orderItem, err := o.orderRepo.UpdateOrder(orderPayload, itemsPayload)
	if err != nil {
		return nil, err
	}

	itemsResponse := []dto.ItemResponse{}
	if len(orderItem.Items) != 0 {
		for _, eachItem := range orderItem.Items {
			itemResponse := eachItem.ItemToItemResponse()

			itemsResponse = append(itemsResponse, itemResponse)
		}
	}

	result := &dto.NewOrderResponse{
		Message: "order updated successfully",
		Data: dto.OrderResponse{
			OrderId:      orderItem.Order.OrderId,
			OrderedAt:    orderItem.Order.OrderedAt,
			CustomerName: orderItem.Order.CustomerName,
			CreatedAt:    orderItem.Order.CreatedAt,
			UpdatedAt:    orderItem.Order.UpdatedAt,
			Items:        itemsResponse,
		},
		StatusCode: http.StatusOK,
	}

	return result, nil
}

func (o *orderService) DeleteOrder(orderId int) (*dto.DeleteOrderResponse, errs.MessageErr) {
	_, err := o.orderRepo.GetOrderById(orderId)
	if err != nil {
		return nil, err
	}

	err = o.orderRepo.DeleteOrder(orderId)
	if err != nil {
		return nil, err
	}

	response := &dto.DeleteOrderResponse{
		Message:    "order deleted successfully",
		StatusCode: http.StatusOK,
	}

	return response, nil
}
