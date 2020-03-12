package shop

import (
	"errors"
	"sync"
	"time"
)

var ProductTypeStruct = map[ProductType]struct{}{
	ProductNormal:  {},
	ProductPremium: {},
	ProductSample:  {},
}

type ProductMutex struct {
	Product map[string]Product
	sync.RWMutex
	time.Duration
}

func NewProduct(name string, price float32, productType ProductType) Product {
	return Product{
		Name:  name,
		Price: price,
		Type:  productType,
	}
}
func (productMutex ProductMutex) getProduct(name string) (Product, error) {
	productMutex.RLock()
	defer productMutex.RUnlock()

	if product, ok := productMutex.Product[name]; ok {
		return product, nil
	} else {
		return product, errors.New("Product not exist")
	}
}

func (productMutex ProductMutex) setProduct(product Product) {
	productMutex.Lock()

	productMutex.Product[product.Name] = product

	productMutex.Unlock()
}

func (productMutex ProductMutex) AddProduct(product Product) error {
	err := CheckProduct(product)
	if err != nil {
		return err
	}
	if _, err := productMutex.getProduct(product.Name); err == nil {
		return err
	}

	productMutex.setProduct(product)
	return nil
}

func (productMutex ProductMutex) ModifyProduct(product Product) error {
	if _, ok := productMutex.getProduct(product.Name); ok != nil {
		return ok
	}
	err := CheckProduct(product)
	if err != nil {
		return err
	}
	productMutex.setProduct(product)
	return nil
}

func (productMutex ProductMutex) RemoveProduct(name string) error {
	if _, err := productMutex.getProduct(name); err != nil {
		return err
	}
	productMutex.Lock()
	delete(productMutex.Product, name)
	productMutex.Unlock()
	return nil
}
