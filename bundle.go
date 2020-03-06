package shop

import (
	"errors"
	"sync"
)

var bundleTypeStruct = map[BundleType]struct{}{
	BundleNormal: {},
	BundleSample: {},
}

type BundleMutex struct {
	Bundle map[string]Bundle
	sync.RWMutex
}

func NewBundle(main Product, bundleType BundleType, discount float32, additional ...Product) Bundle {
	return Bundle{
		Products: append(additional, main),
		Type:     bundleType,
		Discount: discount,
	}
}

//редачить
func (bundleMutex BundleMutex) AddBundle(name string, product Product, bundleType BundleType, discount float32, additional ...Product) error {

	if _, ok := bundleMutex.Bundle[name]; ok {
		return errors.New("bundle exists")
	}

	if _, ok := bundleTypeStruct[bundleType]; !ok {
		return errors.New("no type bundle")
	}
	if discount < 1 || discount > 99 {
		return errors.New("negative discont")
	}
	bundleMutex.Lock()
	defer bundleMutex.Unlock()

	if product.Type == ProductSample {
		return errors.New("additional product ")
	}

	b := NewBundle(product, bundleType, discount, additional...)
	bundleMutex.Bundle[name] = b
	return nil
}

func (bundleMutex BundleMutex) ChangeDiscount(name string, discount float32) error {

	if discount < 1 || discount > 99 {
		return errors.New("not discount")
	}

	bundleMutex.Lock()
	defer bundleMutex.Unlock()

	if _, ok := bundleMutex.Bundle[name]; !ok {
		return errors.New("not bundle")
	}
	//bundleMutex.Bundle[name].Discount = discount
	return nil
}

func (bundleMutex BundleMutex) RemoveBundle(name string) error {

	bundleMutex.Lock()
	defer bundleMutex.Unlock()

	if _, ok := bundleMutex.Bundle[name]; !ok {
		return errors.New("not bundle")
	}

	delete(bundleMutex.Bundle, name)
	return nil
}
