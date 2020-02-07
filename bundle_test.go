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
		Price: 0,
		Type:  ProductNormal,
	}, {
		Name:  "Orange",
		Price: 10,
		Type:  ProductNormal,
	},
	}

	err := shop.AddBundle("New", product, -12, productsAdditional...)
	// прокинули некорректную скидку, а получили ошибку про бандл
	require.EqualError(t, err, "not discount")
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

	err := shop.AddBundle("New", product, 11, productsAdditional...)
	if err != nil {
		t.Error(err)
	}

	// метод должен отвечать за добавление нового бандла, а не изменения старого
	err = shop.AddBundle("New", product, 25, productsAdditional...)
	require.NotZero(t, err)
}

func Test_AddBundleWithOneProduct(t *testing.T) {
	shop := NewShop()
	product := Product{
		Name:  "Coffee",
		Price: 10,
		Type:  ProductNormal,
	}

	productsAdditional := []Product{{
		Name:  "Coffee",
		Price: 10,
		Type:  ProductNormal,
	}, {
		Name:  "Coffee",
		Price: 10,
		Type:  ProductNormal,
	},
	}

	// хорошо бы возвращать ошибку, ведь мы добавляем в бандл один и тот же продукт
	err := shop.AddBundle("New", product, 25, productsAdditional...)
	require.NotZero(t, err)
}

func Test_AddBundleWithSampleAndManyOther(t *testing.T) {
	shop := NewShop()
	product := Product{
		Name:  "Orange",
		Price: 10,
		Type:  ProductNormal,
	}

	productsAdditional := []Product{{
		Name:  "Apple",
		Price: 10,
		Type:  ProductNormal,
	}, {
		Name:  "Coffee",
		Price: 0,
		Type:  ProductSample,
	},
	}

	// в бандле есть пробник, но эллементов больше чем два, должно вернуть ошибку
	err := shop.AddBundle("New", product, 25, productsAdditional...)
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
	err := shop.AddBundle("New", product, 25, productAdditional)
	require.NotZero(t, err)
}

func Test_AddBundleSample(t *testing.T) {
	shop := NewShop()
	product := Product{
		Name:  "Orange",
		Price: 10,
		Type:  ProductNormal,
	}

	productAdditional := Product{
		Name:  "Coffee",
		Price: 0,
		Type:  ProductSample,
	}

	// пробник в бандле не бесплатный
	err := shop.AddBundle("New", product, 25, productAdditional)
	if err != nil {
		t.Error(err)
	}

	// мы добавили бандл соответствующий типу BundleSample, но тип остался прежним
	require.Equal(t, BundleSample, shop.Bundles["New"].Type)
}
