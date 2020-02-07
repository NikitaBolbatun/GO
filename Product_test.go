package shop

import (
	"github.com/stretchr/testify/require"
	"testing"
)

//-----AddProduct testing----

func Test_AddProductWithTheSameName(t *testing.T) {
	shop := NewShop()

	productTests := []Product{{
		Name:  "Apple",
		Price: 0,
		Type:  ProductNormal,
	}, {
		Name:  "Apple",
		Price: 10,
		Type:  ProductPremium,
	},
	}

	// скорее всего мы не должны изменять продукт с помощью тестируемого метода
	for _, test := range productTests {
		err := shop.AddProduct(test)
		require.NotZero(t, err)
	}
}

func Test_AddProductSampleWithNotNullPrice(t *testing.T) {
	shop := NewShop()

	product := Product{
		Name:  "Apple",
		Price: 10,
		Type:  ProductSample,
	}

	// пробник может быть только бесплатным
	err := shop.AddProduct(product)
	require.NotZero(t, err)
}

func Test_AddProductWithIncorrectType(t *testing.T) {
	shop := NewShop()

	product := Product{
		Name:  "Apple",
		Price: 10,
		Type:  101,
	}

	// непредусмотренный тип
	err := shop.AddProduct(product)
	require.NotZero(t, err)
}

func Test_AddProductWithNegativePrice(t *testing.T) {
	shop := NewShop()

	product := Product{
		Name:  "Apple",
		Price: -10,
		Type:  ProductSample,
	}

	// в условиях магазина не было, но продукт с отрицательной ценой бывает только в мышеловке
	err := shop.AddProduct(product)
	require.NotZero(t, err)
}

//-----ModifyProduct testing----

//ошибки такие же, что и выше, но есть проверка на существование продукта
func Test_ModifyProductSampleWithNotNullPrice(t *testing.T) {
	shop := NewShop()

	shop.Products["Apple"] = &Product{
		Name:  "Apple",
		Price: 10,
		Type:  ProductNormal,
	}
	product := Product{
		Name:  "Apple",
		Price: -10,
		Type:  ProductSample,
	}

	// пробник может быть только бесплатным
	err := shop.AddProduct(product)
	require.NotZero(t, err)
}

func Test_ModifyProductWithIncorrectType(t *testing.T) {
	shop := NewShop()

	shop.Products["Apple"] = &Product{
		Name:  "Apple",
		Price: 10,
		Type:  ProductNormal,
	}
	product := Product{
		Name:  "Apple",
		Price: 10,
		Type:  101,
	}

	// непредусмотренный тип
	err := shop.ModifyProduct(product)
	require.NotZero(t, err)
}

func Test_ModifyProductWithNegativePrice(t *testing.T) {
	shop := NewShop()

	shop.Products["Apple"] = &Product{
		Name:  "Apple",
		Price: 10,
		Type:  ProductNormal,
	}
	product := Product{
		Name:  "Apple",
		Price: -10,
		Type:  ProductNormal,
	}

	err := shop.ModifyProduct(product)
	require.NotZero(t, err)
}

//-----RemoveProduct testing----
