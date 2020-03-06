package shop

import (
	"errors"
	"sync"
)

var ProductTypeStruct = map[ProductType]struct{}{
	ProductNormal:  {},
	ProductPremium: {},
	ProductSample:  {},
}

type ProductMutex struct {
	Product map[string]Product
	sync.RWMutex
}

func NewProduct(name string, price float32, productType ProductType) Product {
	return Product{
		Name:  name,
		Price: price,
		Type:  productType,
	}
}

func (productMutex ProductMutex) AddProduct(product Product) error {
	err := CheckProduct(product)
	if err != nil {
		return err
	}
	productMutex.Lock()
	defer productMutex.Unlock()
	productMutex.Product[product.Name] = product

	return nil
}

func (productMutex ProductMutex) ModifyProduct(product Product) error {
	if _, ok := productMutex.Product[product.Name]; !ok {
		return errors.New("product not found")
	}
	err := CheckProduct(product)
	if err != nil {
		return err
	}
	productMutex.Lock()
	defer productMutex.Unlock()

	productMutex.Product[product.Name] = product

	return nil
}

func (productMutex ProductMutex) RemoveProduct(name string) error {

	productMutex.Lock()
	defer productMutex.Unlock()

	if _, ok := productMutex.Product[name]; !ok {
		return errors.New("cannot delete nil product")
	}

	delete(productMutex.Product, name)
	return nil
}
