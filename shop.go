package shop

import (
	"encoding/json"
	"errors"
)

func NewShop() *S {
	return &S{
		ProductMutex: ProductMutex{Product: make(map[string]Product)},
		AccountMutex: AccountMutex{Account: make(map[string]Account)},
		BundleMutex:  BundleMutex{Bundle: make(map[string]Bundle)},
	}
}
func CheckNameAccount(name string) error {
	if len(name) == 0 {
		return errors.New("username not correct")
	}

	if name == " " {
		return errors.New("username not correct")
	}

	return nil
}

func CheckProduct(product Product) error {
	if len(product.Name) == 0 {
		return errors.New("username not correct")
	}

	if product.Name == " " {
		return errors.New("username not correct")
	}
	if product.Type == ProductSample {
		if product.Price != 0 {
			return errors.New("sample not 0")
		}
	} else {
		if product.Price <= 0 {
			return errors.New("negative price")
		}
	}
	if _, ok := ProductTypeStruct[product.Type]; !ok {
		return errors.New("no type product")
	}

	return nil
}

func (s S) Import(data []byte) error {
	return json.Unmarshal(data, s)
}

func (s S) Export() ([]byte, error) {
	return json.MarshalIndent(s, "", "")
}
