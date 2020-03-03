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

type ProductsMutex struct {
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

func (s S) AddProduct(product Product) error {
	err := CheckProduct(product)
	if err != nil {
		return err
	}
	s.productMutex.Lock()
	defer s.productMutex.Unlock()
	s.Products[product.Name] = &product

	return nil
}

func (s S) ModifyProduct(product Product) error {
	if _, ok := s.Products[product.Name]; !ok {
		return errors.New("product not found")
	}
	err := CheckProduct(product)
	if err != nil {
		return err
	}
	s.productMutex.Lock()
	defer s.productMutex.Unlock()

	s.Products[product.Name] = &product

	return nil
}

func (s S) RemoveProduct(name string) error {

	s.productMutex.Lock()
	defer s.productMutex.Unlock()

	if _, ok := s.Products[name]; !ok {
		return errors.New("cannot delete nil product")
	}

	delete(s.Products, name)
	return nil
}
