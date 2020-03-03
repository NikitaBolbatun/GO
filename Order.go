package shop

import (
	"errors"
)

var DiscountType = map[ProductType]map[AccountType]float32{
	ProductPremium: {AccountPremium: 20, AccountNormal: 5},
	ProductNormal:  {AccountPremium: -50, AccountNormal: 0},
}

func (s S) CalculateOrder(name string, order Order) (float32, error) {

	s.accountMutex.RLock()
	account, err := s.GetAccount(name)
	s.accountMutex.RUnlock()
	if err != nil {
		return 0, errors.New("not name register")
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

func (s S) PlaceOrder(name string, order Order) (error, error) {
	s.orderMutex.Lock()
	price, err := s.CalculateOrder(name, order)
	s.orderMutex.Unlock()
	if err != nil {
		return errors.New("not calculate order"), nil
	}

	account, ok := s.Accounts[name]
	if !ok {
		return errors.New("not name register"), nil
	}

	if account.Balance < price {
		return errors.New("insufficient funds"), nil
	}
	account.Balance -= price
	return nil, nil
}
