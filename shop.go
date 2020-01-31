package shop

import (
	"encoding/json"
)

func NewShop() *S {
	return &S{
		Products: make(map[string]Product),
		Bundles:  make(map[string]Bundle),
		Accounts: make(map[string]Account),
	}
}

func (s S) Import(data []byte) error {
	return json.Unmarshal(data, s)
}

func (s S) Export() ([]byte, error) {
	return json.MarshalIndent(s, "", "")
}
