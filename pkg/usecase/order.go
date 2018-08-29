package usecase

import (
	"fmt"

	"chatserver/pkg/domain"
	"chatserver/pkg/lib/config"
)

type Order struct {
	UserID    int `json:"user_id"`
	Quantity  int `json:"quantity"`
	ProductID int `json:"product_id"`
}

type OrderUsecase struct {
	cfg config.Config
	od  domain.OrderDomain
	ud  domain.UserDomain
}

func InitOrderUsecase(cfg config.Config, order domain.OrderDomain, user domain.UserDomain) *OrderUsecase {
	return &OrderUsecase{
		cfg: cfg,
		od:  order,
		ud:  user,
	}
}

func (o *OrderUsecase) PutNewOrder(order Order) (string, error) {

	// using user domain to validate user
	if !o.ud.IsValidUser(order.UserID) {
		return "", fmt.Errorf("invalid user")
	}

	// create new Order entity of order domain
	newOrder := domain.Order{
		Quantity:  order.Quantity,
		ProductID: order.ProductID,
	}

	// using order domain to validate stock
	if !o.od.IsValidStock(newOrder) {
		return "", fmt.Errorf("Stock is not enough")
	}

	err := o.od.CreateOrder(&newOrder)
	if err != nil {
		return "", err
	}

	return newOrder.Invoice, nil
}
