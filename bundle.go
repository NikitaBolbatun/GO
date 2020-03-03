package shop

import "errors"

var bundleTypeStruct = map[BundleType]struct{}{
	BundleNormal: {},
	BundleSample: {},
}

func NewBundle(main Product, bundleType BundleType, discount float32, additional ...Product) Bundle {
	return Bundle{
		Products: append(additional, main),
		Type:     bundleType,
		Discount: discount,
	}
}

//редачить
func (s S) AddBundle(name string, product Product, bundleType BundleType, discount float32, additional ...Product) error {

	if _, ok := s.Bundles[name]; ok {
		return errors.New("bundle exists")
	}

	if _, ok := bundleTypeStruct[bundleType]; !ok {
		return errors.New("no type bundle")
	}
	if discount < 1 || discount > 99 {
		return errors.New("negative discont")
	}
	s.bundleMutex.Lock()
	defer s.bundleMutex.Unlock()

	if product.Type == ProductSample {
		return errors.New("additional product ")
	}

	b := NewBundle(product, bundleType, discount, additional...)
	s.Bundles[name] = &b
	return nil
}

func (s S) ChangeDiscount(name string, discount float32) error {

	if discount < 1 || discount > 99 {
		return errors.New("not discount")
	}

	s.bundleMutex.Lock()
	defer s.bundleMutex.Unlock()

	if _, ok := s.Bundles[name]; !ok {
		return errors.New("not bundle")
	}

	s.Bundles[name].Discount = discount
	return nil
}

func (s S) RemoveBundle(name string) error {

	s.bundleMutex.Lock()
	defer s.bundleMutex.Unlock()

	if _, ok := s.Bundles[name]; !ok {
		return errors.New("not bundle")
	}

	delete(s.Bundles, name)
	return nil
}
