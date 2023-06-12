package item_repository

import (
	"github.com/Iqbalhasanu/Msib-Hacktiv8-Golang-Assignment2.git/entity"
	"github.com/Iqbalhasanu/Msib-Hacktiv8-Golang-Assignment2.git/pkg/errs"
)

type ItemRepository interface {
	FindItemsByItemCode([]string, int) ([]*entity.Items, errs.MessageErr)
}
