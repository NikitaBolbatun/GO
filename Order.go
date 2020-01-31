package shop

import (
	"errors"
)

var DiscountType = map[ProductType]map[AccountType]float32{
	ProductPremium: {AccountPremium: 20, AccountNormal: 5},
	ProductNormal:  {AccountPremium: -50, AccountNormal: 0},
}

func (s S) CalculateOrder(name string, order Order) float32 {

	account, err := s.GetAccount(name)
	if err != nil {
		return 0
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
			return 0
		}

		price := float32(0)
		for _, product := range bundle.Products {
			price += product.Price
		}
		bundlesPrice += price * (1 - bundle.Discount*0.01)
	}

	allprice := ProductsMoney + bundlesPrice
	return allprice
}

func (s S) PlaceOrder(name string, order Order) error {

	var price = s.CalculateOrder(name, order)
	acc, ok := s.Accounts[name]
	if !ok {
		return errors.New("not name register")
	}

	acc.Balance -= price
	return nil
}
