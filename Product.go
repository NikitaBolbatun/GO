package shop

import "errors"

func (s S) AddProduct(product Product) error {
	if product.Name == "" {
		return errors.New("product without name")
	}

	s.Products[product.Name] = &product

	return nil
}

func (s S) ModifyProduct(product Product) error {
	if _, ok := s.Products[product.Name]; !ok {
		return errors.New("product not found")
	}
	s.Products[product.Name] = &product

	return nil
}

func (s S) RemoveProduct(name string) error {

	if _, ok := s.Products[name]; !ok {
		return errors.New("cannot delete nil product")
	}

	delete(s.Products, name)
	return nil
}
