package shop

import (
	"github.com/stretchr/testify/require"
	"testing"
)

//-----CalculateOrder-----

func Test_CalculateOrderBundleDiscount(t *testing.T) {
	shop := NewShop()

	shop.Accounts["Dimas"] = &Account{
		Name:    "Dimas",
		Balance: 100,
		Type:    AccountPremium,
	}

	products := []Product{{
		Name:  "banana",
		Price: 40,
		Type:  ProductNormal,
	}, {
		Name:  "sapre",
		Price: 30,
		Type:  ProductNormal,
	}, {
		Name:  "apple",
		Price: 30,
		Type:  ProductPremium,
	},
	}

	bundles := []Bundle{
		{Products: products, Type: BundleNormal, Discount: 32},
	}

	// при подсчете суммы бандла не учитываются скидки на товары составляющие его
	order := Order{
		Products: products,
		Bundles:  bundles,
	}

	price, err := shop.CalculateOrder("Dimas", order)
	if err != nil {
		t.Error(err)
	}
	require.Equal(t, float32(216.72), price)
}

func Test_CalculateOrderFloatError(t *testing.T) {
	shop := NewShop()

	shop.Accounts["Dimas"] = &Account{
		Name:    "Dimas",
		Balance: 100,
		Type:    AccountNormal,
	}

	products := make([]Product, 1000)
	for i := 0; i < 1000; i++ {
		products[i] = Product{
			Name:  "product # " + string(i),
			Price: 123.5,
			Type:  ProductPremium,
		}
	}

	order := Order{
		Products: products,
		Bundles:  nil,
	}

	price, err := shop.CalculateOrder("Dimas", order)
	if err != nil {
		t.Error(err)
	}
	// погрешность...
	require.Equal(t, float32(117325), price)
}

func Test_CalculateOrderTypesErrors(t *testing.T) {
	// тесты на отсуствие проверок
	t.Run("Test Calculate Order with not null price product", CalculateOrderProductSampleNotNullPrice)
	t.Run("Test Calculate Order with unknowing type of product", CalculateOrderProductWithUnknowingType)
	t.Run("Test Calculate Order with incorrect bundle", CalculateOrderWithIncorrectBundle)
}

func CalculateOrderProductSampleNotNullPrice(t *testing.T) {
	shop := NewShop()

	shop.Accounts["Dimas"] = &Account{
		Name:    "Dimas",
		Balance: 100,
		Type:    AccountNormal,
	}

	products := []Product{{
		Name:  "banana",
		Price: 40,
		Type:  ProductNormal,
	}, {
		Name:  "sapre",
		Price: 30,
		Type:  ProductNormal,
	}, {
		Name:  "apple",
		Price: 30,
		Type:  ProductSample,
	},
	}

	order := Order{
		Products: products,
		Bundles:  nil,
	}

	//в заказе есть пробник с нулевой ценой
	_, err := shop.CalculateOrder("Dimas", order)
	require.NotZero(t, err)
}

func CalculateOrderProductWithUnknowingType(t *testing.T) {
	shop := NewShop()

	shop.Accounts["Dimas"] = &Account{
		Name:    "Dimas",
		Balance: 100,
		Type:    AccountNormal,
	}

	products := []Product{{
		Name:  "banana",
		Price: 40,
		Type:  ProductNormal,
	}, {
		Name:  "sapre",
		Price: 30,
		Type:  ProductNormal,
	}, {
		Name:  "apple",
		Price: 30,
		Type:  12,
	},
	}

	order := Order{
		Products: products,
		Bundles:  nil,
	}

	//в заказе есть продукт с неизвестным типом, такая же ошибка будет и с бандлом
	_, err := shop.CalculateOrder("Dimas", order)
	require.NotZero(t, err)
}

func CalculateOrderWithIncorrectBundle(t *testing.T) {
	shop := NewShop()

	shop.Accounts["Dimas"] = &Account{
		Name:    "Dimas",
		Balance: 100,
		Type:    AccountNormal,
	}

	products := []Product{{
		Name:  "banana",
		Price: 40,
		Type:  ProductNormal,
	}, {
		Name:  "sapre",
		Price: 30,
		Type:  ProductNormal,
	}, {
		Name:  "apple",
		Price: 0,
		Type:  ProductSample,
	},
	}

	bundle := Bundle{
		Products: products,
		Type:     BundleSample,
		Discount: 12,
	}

	order := Order{
		Products: nil,
		Bundles:  []Bundle{bundle},
	}

	//в заказе есть некорректно сформированный бандл
	_, err := shop.CalculateOrder("Dimas", order)
	require.NotZero(t, err)
}

//-----PlaceOrder-----

func TestS_PlaceOrderFloatError(t *testing.T) {
	shop := NewShop()

	shop.Accounts["Dimas"] = &Account{
		Name:    "Dimas",
		Balance: 10000,
		Type:    AccountNormal,
	}

	products := make([]Product, 10000)
	for i := 0; i < 10000; i++ {
		products[i] = Product{
			Name:  "product # " + string(i),
			Price: 0.1,
			Type:  ProductPremium,
		}
	}

	order := Order{
		Products: products,
		Bundles:  nil,
	}

	_, err := shop.PlaceOrder("Dimas", order)
	if err != nil {
		t.Error(err)
	}
	// погрешность...
	require.Equal(t, float32(9050), shop.Accounts["Dimas"].Balance)
}
