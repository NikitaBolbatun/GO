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

func (bundleMutex BundleMutex) getBundle(name string) (Bundle, error) {
	bundleMutex.RLock()
	defer bundleMutex.RUnlock()

	if bundle, ok := bundleMutex.Bundle[name]; ok {
		return bundle, nil
	} else {
		return Bundle{}, errors.New("bundle no exist")
	}
}

func (bundleMutex BundleMutex) setBundle(name string, bundle Bundle) {
	bundleMutex.Lock()

	bundleMutex.Bundle[name] = bundle

	bundleMutex.Unlock()
}
func NewBundle(main Product, bundleType BundleType, discount float32, additional ...Product) Bundle {
	return Bundle{
		Products: append(additional, main),
		Type:     bundleType,
		Discount: discount,
	}
}

func (bundleMutex BundleMutex) AddBundle(name string, product Product, bundleType BundleType, discount float32, additional ...Product) error {

	if _, ok := bundleMutex.getBundle(name); ok == nil {
		return errors.New("bundle exists")
	}

	if _, ok := bundleTypeStruct[bundleType]; !ok {
		return errors.New("no type bundle")
	}
	if discount < 1 || discount > 99 {
		return errors.New("negative discont")
	}
	if product.Type == ProductSample {
		return errors.New("additional product ")
	}
	bundleMutex.setBundle(name, NewBundle(product, bundleType, discount, additional...))
	return nil
}

func (bundleMutex BundleMutex) ChangeDiscount(name string, discount float32) error {

	if discount < 1 || discount > 99 {
		return errors.New("not discount")
	}
	bundle, err := bundleMutex.getBundle(name)
	if err != nil {
		return err
	}
	bundle.Discount = discount

	bundleMutex.setBundle(name, bundle)
	return nil
}

func (bundleMutex BundleMutex) RemoveBundle(name string) error {

	if _, err := bundleMutex.getBundle(name); err != nil {
		return err
	}

	bundleMutex.Lock()

	delete(bundleMutex.Bundle, name)

	bundleMutex.Unlock()
	return nil
}
