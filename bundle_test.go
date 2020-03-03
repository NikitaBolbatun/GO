package shop

import (
	"github.com/stretchr/testify/require"
	"testing"
)

//-----AddBundle-----

func Test_AddBundleIncorrectDiscount(t *testing.T) {
	shop := NewShop()

	product := Product{
		Name:  "Coffee",
		Price: 10,
		Type:  ProductNormal,
	}

	productsAdditional := []Product{{
		Name:  "Apple",
		Price: 5,
		Type:  ProductNormal,
	}, {
		Name:  "Orange",
		Price: 10,
		Type:  ProductNormal,
	},
	}

	err := shop.AddBundle("New", product, BundleNormal, -12, productsAdditional...)
	if err == nil {
		t.Fatalf("%v", err)
	}
}

func Test_AddBundleWithTheSameName(t *testing.T) {
	shop := NewShop()

	product := Product{
		Name:  "Coffee",
		Price: 10,
		Type:  ProductNormal,
	}

	productsAdditional := []Product{{
		Name:  "Apple",
		Price: 0,
		Type:  ProductNormal,
	}, {
		Name:  "Orange",
		Price: 10,
		Type:  ProductNormal,
	},
	}
	err := shop.AddBundle("New", product, BundleNormal, 30, productsAdditional...)
	if err != nil {
		t.Error(err)
	}
	// метод должен отвечать за добавление нового бандла, а не изменения старого
	err = shop.AddBundle("New", product, BundleNormal, 25, productsAdditional...)

	require.NotZero(t, err)
}

func Test_AddBundleSampleWithNotNullPrice(t *testing.T) {
	shop := NewShop()
	product := Product{
		Name:  "Orange",
		Price: 10,
		Type:  ProductNormal,
	}

	productAdditional := Product{
		Name:  "Coffee",
		Price: 12,
		Type:  ProductSample,
	}

	// пробник в бандле не бесплатный
	err := shop.AddBundle("New", product, 25, 30, productAdditional)
	require.NotZero(t, err)
}
