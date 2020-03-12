package shop

import (
	"errors"
	"time"
)

type Decor struct {
	timeout time.Duration
	s S
}

func (decor *Decor) AddProduct(product Product) error {
	ch := make(chan error, 1)
	go func() {
		ch <- decor.s.ProductMutex.AddProduct(product)
	}()
	select {
	case err := <-ch:
		return err
	case <-time.After(decor.s.ProductMutex.Duration):
		return errors.New("timeOut")
	}
}

func (decor *Decor) ModifyProduct(product Product) error {
	ch := make(chan error, 1)
	go func() {
		ch <- decor.s.ProductMutex.ModifyProduct(product)
	}()
	select {
	case err := <-ch:
		return err
	case <-time.After(decor.s.ProductMutex.Duration):
		return errors.New("timeOut")
	}
}

func (decor *Decor) RemoveProduct(name string) error {
	ch := make(chan error, 1)
	go func() {
		ch <- decor.s.ProductMutex.RemoveProduct(name)
	}()
	select {
	case err := <-ch:
		return err
	case <-time.After(decor.s.ProductMutex.Duration):
		return errors.New("timeOut")
	}
}

func (decor *Decor) Register(name string) error {
	ch := make(chan error, 1)
	go func() {
		ch <- decor.s.AccountMutex.Register(name)
	}()
	select {
	case err := <-ch:
		return err
	case <-time.After(decor.s.AccountMutex.Duration):
		return errors.New("timeOut")
	}
}

func (decor *Decor) AddBalance(name string, sum float32) error {
	ch := make(chan error, 1)
	go func() {
		ch <- decor.s.AccountMutex.AddBalance(name, sum)
	}()
	select {
	case err := <-ch:
		return err
	case <-time.After(decor.s.AccountMutex.Duration):
		return errors.New("timeOut")
	}
}

func (decor *Decor) Balance(name string) (float32, error) {
	type str struct {
		balance float32
		error   error
	}

	ch := make(chan str, 1)
	go func() {
		balance, err := decor.s.AccountMutex.Balance(name)
		ch <- str{
			balance: balance,
			error:   err,
		}
	}()
	select {
	case res := <-ch:
		return res.balance, res.error
	case <-time.After(decor.s.AccountMutex.Duration):
		return 0, errors.New("timeout")
	}
}

func (decor Decor) CalculateOrder(name string, order Order) (float32, error) {
	type str struct {
		sum   float32
		error error
	}

	ch := make(chan str, 1)
	go func() {
		sum, err := decor.s.OrdersMutex.CalculateOrder(name, order)
		ch <- str{
			sum:   sum,
			error: err,
		}
	}()
	select {
	case res := <-ch:
		return res.sum, res.error
	case <-time.After(decor.s.OrdersMutex.Duration):
		return 0, errors.New("timeout")
	}
}

func (decor *Decor) PlaceOrder(name string, order Order) error {
	ch := make(chan error, 1)
	go func() {
		ch <- decor.s.OrdersMutex.PlaceOrder(name, order)
	}()
	select {
	case err := <-ch:
		return err
	case <-time.After(decor.s.OrdersMutex.Duration):
		return errors.New("timeout")
	}
}
func (decor *Decor) AddBundle(name string, main Product, discount float32, additional ...Product) error {
	ch := make(chan error, 1)
	go func() {
		ch <- decor.s.BundleMutex.AddBundle(name, main, discount, additional...)
	}()
	select {
	case err := <-ch:
		return err
	case <-time.After(decor.s.BundleMutex.Duration):
		return errors.New("timeout")
	}
}

func (decor *Decor) ChangeDiscount(name string, discount float32) error {
	ch := make(chan error, 1)
	go func() {
		ch <- decor.s.BundleMutex.ChangeDiscount(name, discount)
	}()
	select {
	case err := <-ch:
		return err
	case <-time.After(decor.s.BundleMutex.Duration):
		return errors.New("timeout")
	}
}

func (decor *Decor) RemoveBundle(name string) error {
	ch := make(chan error, 1)
	go func() {
		ch <- decor.s.BundleMutex.RemoveBundle(name)
	}()
	select {
	case err := <-ch:
		return err
	case <-time.After(decor.s.BundleMutex.Duration):
		return errors.New("timeout")
	}
}

