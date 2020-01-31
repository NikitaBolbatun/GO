package shop

import (
	"errors"
)

var DiscountType = map[ProductType]map[AccountType]float32{
	ProductPremium: {AccountPremium: 20, AccountNormal: 5},
	ProductNormal:  {AccountPremium: -50, AccountNormal: 0},
}

func (s S) CalculateOrder(name string, order Order) (float32, error) {

	account, err := s.GetAccount(name)
	if err != nil {
		return 0, errors.New("not name register")
	}

	// products
	ProductsMoney := float32(0)
	for _, product := range order.Products {
		var discount = DiscountType[product.Type][account.Type]
		ProductsMoney += product.Price * (1 - discount*0.01)
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

	allPrice := ProductsMoney + bundlesPrice

	return allPrice, nil
}

func (s S) PlaceOrder(name string, order Order) (int, error) {

	price, err := s.CalculateOrder(name, order)
	if err != nil {
		return 0, errors.New("not calculate order")
	}

	account, ok := s.Accounts[name]
	if !ok {
		return 0, errors.New("not name register")
	}

	if account.Balance < price {
		return 0, errors.New("insufficient funds")
	}
	account.Balance -= price
	return 0, nil
}
