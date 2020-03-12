package shop

import (
	"errors"
	"sync"
	"time"
)

type OrdersMutex struct {
	AccountMutex
	sync.RWMutex
	orders map[string]float32
	time.Duration
}

var DiscountType = map[ProductType]map[AccountType]float32{
	ProductPremium: {AccountPremium: 20, AccountNormal: 5},
	ProductNormal:  {AccountPremium: -50, AccountNormal: 0},
}

func (ordersMutex OrdersMutex) CalculateOrder(name string, order Order) (float32, error) {
	if order.Products == nil &&
		order.Bundles == nil {
		return 0, errors.New("order items not init")
	}
	if len(order.Products) == 0 &&
		len(order.Bundles) == 0 {
		return 0, errors.New("not purchases")
	}
	account, ok := ordersMutex.AccountMutex.getAccount(name)
	if ok != nil {
		return 0, ok
	}
	// discount
	allDiscount := float32(0)
	for _, product := range order.Products {
		var discount = DiscountType[product.Type][account.Type]
		allDiscount += product.Price * (1 - discount*0.01)
	}

	// bundles
	bundlesPrice := float32(0)
	for _, bundle := range order.Bundles {

		if bundle.Discount < 1 || bundle.Discount > 99 {
			return 0, errors.New("size discount error")
		}

		price := float32(0)
		for _, product := range bundle.Products {
			price += product.Price
		}
		bundlesPrice += price * (1 - bundle.Discount*0.01)
	}

	allPrice := allDiscount + bundlesPrice

	return allPrice, nil
}

func (ordersMutex OrdersMutex) PlaceOrder(name string, order Order) error {
	price, err := ordersMutex.CalculateOrder(name, order)
	if err != nil {
		return errors.New("not calculate order")
	}

	account, ok := ordersMutex.AccountMutex.Account[name]
	if !ok {
		return errors.New("not name register")
	}

	if account.Balance < price {
		return errors.New("insufficient funds")
	}
	account.Balance -= price

	return nil
}
