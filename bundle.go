package shop

import "errors"

func NewBundle(main Product, discount float32, additional ...Product) Bundle {
	return Bundle{
		Products: append(additional, main),
		Type:     BundleNormal,
		Discount: discount,
	}
}

func (s S) AddBundle(name string, product Product, discount float32, additional ...Product) error {

	if discount < 1 || discount > 99 {
		return errors.New("bundle not found")
	}

	b := NewBundle(product, discount, additional...)
	s.Bundles[name] = &b
	return nil
}

func (s S) ChangeDiscount(name string, discount float32) error {

	if discount < 1 || discount > 99 {
		return errors.New("not discount")
	}

	if _, ok := s.Bundles[name]; !ok {
		return errors.New("not bundle")
	}

	s.Bundles[name].Discount = discount
	return nil
}

func (s S) RemoveBundle(name string) error {

	if _, ok := s.Bundles[name]; !ok {
		return errors.New("not bundle")
	}

	delete(s.Bundles, name)
	return nil
}
