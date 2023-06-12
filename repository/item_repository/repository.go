package item_repository

import (
	"github.com/Iqbalhasanu/Msib-Hacktiv8-Golang-Assignment2/entity"
	"github.com/Iqbalhasanu/Msib-Hacktiv8-Golang-Assignment2/pkg/errs"
)

type ItemRepository interface {
	FindItemsByItemCodes(itemCodes []string) ([]*entity.Item, errs.MessageErr)
}
