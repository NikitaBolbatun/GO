package shop

import (
	"testing"
)

func Test_AddProductSuccess(t *testing.T) {
	shop := NewShop()
	err := shop.AddProduct(NewProduct("Чай", 10, ProductNormal))

	if err != nil {
		t.Fatalf("Not success add product = %v", err)
	}
}

func Test_AddProductFailedEmptyName(t *testing.T) {
	shop := NewShop()
	err := shop.AddProduct(NewProduct("", 10, ProductNormal))

	if err == nil {
		t.Fatalf("Name product empty = %v", err)
	}
}

func Test_AddProductFailedNameSpace(t *testing.T) {
	shop := NewShop()
	err := shop.AddProduct(NewProduct(" ", 10, ProductNormal))

	if err == nil {
		t.Fatalf("Name product space = %v", err)
	}
}

func Test_RemoveProduct(t *testing.T) {
	shop := NewShop()
	shop.AddProduct(NewProduct("Чай", 10, ProductNormal))
	err := shop.RemoveProduct("Чай")
	if err != nil {
		t.Fatalf("Not delete product = %v", err)
	}
}

func Test_RemoveProductWhoseNot(t *testing.T) {
	shop := NewShop()
	shop.AddProduct(NewProduct("Чай", 10, ProductNormal))
	err := shop.RemoveProduct("Кофе")
	if err == nil {
		t.Fatalf("delete  product whose not = %v", err)
	}
}
