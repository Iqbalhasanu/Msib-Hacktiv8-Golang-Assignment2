package service

import (
	"fmt"

	"github.com/Iqbalhasanu/Msib-Hacktiv8-Golang-Assignment2.git/entity"
	"github.com/Iqbalhasanu/Msib-Hacktiv8-Golang-Assignment2.git/pkg/errs"
	"github.com/Iqbalhasanu/Msib-Hacktiv8-Golang-Assignment2.git/repository/item_repository"
)

type itemService struct {
	itemRepo item_repository.ItemRepository
}

type ItemService interface {
	FindItemsByItemCode([]string, int) ([]*entity.Items, errs.MessageErr)
}

func NewItemService(itemRepo item_repository.ItemRepository) ItemService {
	return &itemService{
		itemRepo: itemRepo,
	}
}

func (i *itemService) FindItemsByItemCode(itemCodes []string, orderId int) ([]*entity.Items, errs.MessageErr) {
	items, err := i.itemRepo.FindItemsByItemCode(itemCodes, orderId)
	if err != nil {
		return nil, err
	}

	for _, eachItemCode := range itemCodes {
		isFound := false

		for _, eachItem := range items {
			if eachItemCode == eachItem.ItemCode {
				isFound = true
			}
		}

		if !isFound {
			return nil, errs.NotFound(fmt.Sprintf("item with code %s doesn't exists", eachItemCode))
		}
	}

	return items, nil
}
