package service

import (
	"net/http"

	"github.com/Iqbalhasanu/Msib-Hacktiv8-Golang-Assignment2/dto"
	"github.com/Iqbalhasanu/Msib-Hacktiv8-Golang-Assignment2/entity"
	"github.com/Iqbalhasanu/Msib-Hacktiv8-Golang-Assignment2/pkg/errs"
	"github.com/Iqbalhasanu/Msib-Hacktiv8-Golang-Assignment2/repository/order_repository"
)

type orderService struct {
	orderRepo   order_repository.OrderRepository
	itemService ItemService
}

type OrderService interface {
	CreateOrder(payload dto.NewOrderRequest) (*dto.NewOrderResponse, errs.MessageErr)
	GetAllOrder() (*[]order_repository.OrderItem, errs.MessageErr)
	UpdateOrder(orderId int, payload dto.NewOrderRequest) (*dto.GetOrderResponse, errs.MessageErr)
	DeleteOrder(id int) errs.MessageErr
}

func NewOrderService(orderRepo order_repository.OrderRepository, itemService ItemService) OrderService {
	return &orderService{
		orderRepo:   orderRepo,
		itemService: itemService,
	}
}

func (o *orderService) CreateOrder(payload dto.NewOrderRequest) (*dto.NewOrderResponse, errs.MessageErr) {
	orderPayload := entity.Order{
		OrderedAt:    payload.OrderedAt,
		CustomerName: payload.CustomerName,
	}

	itemsPayload := []entity.Item{}
	itemsResponse := []dto.ItemRequest{}

	for _, eachItem := range payload.Items {
		item := entity.Item{
			ItemCode:    eachItem.ItemCode,
			Quantity:    eachItem.Quantity,
			Description: eachItem.Description,
		}

		items := dto.ItemRequest{
			ItemCode:    eachItem.ItemCode,
			Quantity:    eachItem.Quantity,
			Description: eachItem.Description,
		}
		itemsPayload = append(itemsPayload, item)
		itemsResponse = append(itemsResponse, items)
	}

	newOrder, err := o.orderRepo.CreateOrder(orderPayload, itemsPayload)

	if err != nil {
		return nil, err
	}

	response := &dto.NewOrderResponse{
		Message: "Success",
		Data: dto.NewOrderRequest{
			OrderedAt:    newOrder.OrderedAt,
			CustomerName: newOrder.CustomerName,
			Items:        itemsResponse,
		},
		StatusCode: http.StatusCreated,
	}

	return response, nil
}

func (o *orderService) GetAllOrder() (*[]order_repository.OrderItem, errs.MessageErr) {

	orders, err := o.orderRepo.GetAllOrder()

	if err != nil {
		return nil, err
	}

	return orders, nil

}

func (o *orderService) UpdateOrder(orderId int, payload dto.NewOrderRequest) (*dto.GetOrderResponse, errs.MessageErr) {
	itemCodes := payload.ItemsToItemCode()

	_, err := o.itemService.FindItemsByItemCodes(itemCodes)

	if err != nil {
		return nil, err
	}

	orderPayload := entity.Order{
		OrderId:      orderId,
		OrderedAt:    payload.OrderedAt,
		CustomerName: payload.CustomerName,
	}

	itemsPayload := []entity.Item{}

	for _, eachItem := range payload.Items {
		item := entity.Item{
			ItemCode:    eachItem.ItemCode,
			Quantity:    eachItem.Quantity,
			Description: eachItem.Description,
		}

		itemsPayload = append(itemsPayload, item)
	}

	orderItem, err := o.orderRepo.UpdateOrder(orderPayload, itemsPayload)

	if err != nil {
		return nil, err
	}

	itemsResponse := []dto.ItemResponse{}

	for _, eachItem := range orderItem.Items {
		itemResponse := eachItem.ItemToItemResponse()

		itemsResponse = append(itemsResponse, itemResponse)
	}

	result := dto.GetOrderResponse{
		Code: http.StatusOK,
		Data: dto.OrderResponse{
			Id:           orderItem.OrderId,
			CreatedAt:    orderItem.CreatedAt,
			UpdatedAt:    orderItem.UpdatedAt,
			CustomerName: orderItem.CustomerName,
			Items:        itemsResponse,
		},
	}

	return &result, nil
}

func (o *orderService) DeleteOrder(id int) errs.MessageErr {

	err := o.orderRepo.DeleteOrder(id)

	if err != nil {
		return err
	}

	return nil
}
