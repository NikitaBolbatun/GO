package shop

import (
	"errors"
	"sync"
)

var AccountTypeStruct = map[AccountType]struct{}{
	AccountNormal:  {},
	AccountPremium: {},
}

func NewAccount(name string) Account {
	return Account{
		Name:    name,
		Balance: 0,
		Type:    AccountNormal,
	}
}

type AccountMutex struct {
	Account map[string]Account
	sync.RWMutex
}

func (accountMutex *AccountMutex) getAccount(name string) (Account, error) {
	account, err := accountMutex.Account[name]
	if !err {
		return Account{}, errors.New("username exists")
	}
	return account, nil
}

func (accountMutex *AccountMutex) setAccount(name string) error {
	if err := CheckNameAccount(name); err != nil {
		return err
	}
	var accounts = NewAccount(name)
	accountMutex.Account[name] = accounts
	return nil
}
func (accountMutex *AccountMutex) Register(name string) error {
	AccountMutex.Lock()
	defer AccountMutex.Unlock()
	err := CheckNameAccount(name)
	if err != nil {
		return err
	}
	err = accountMutex.setAccount(name)
	if err != nil {
		return err
	}
	return nil
}

//Доделать и тесты
func (s S) ModifyAccountType(name string, accountType AccountType) error {

	if _, ok := s.Accounts[name]; !ok {
		return errors.New("no register")
	}
	if _, ok := AccountTypeStruct[accountType]; !ok {
		return errors.New("no type account")
	}

	s.accountMutex.Lock()
	defer s.accountMutex.Unlock()

	account := s.Accounts[name]
	account.Type = accountType
	return nil
}

func (s S) AddBalance(name string, cash float32) error {

	if cash < 0 {
		return errors.New("negative cash")
	}

	if _, ok := s.Accounts[name]; !ok {
		return errors.New("username not exists")
	}

	account := s.Accounts[name]
	account.Balance += cash

	s.Accounts[name] = account
	return nil
}

func (s S) Balance(name string) (float32, error) {

	if _, ok := s.Accounts[name]; !ok {
		return 0, errors.New("no register")
	}

	s.accountMutex.Lock()
	defer s.accountMutex.Unlock()

	return s.Accounts[name].Balance, nil
}

func (s S) GetAccount(name string) (Account, error) {
	account, ok := s.Accounts[name]

	s.accountMutex.Lock()
	defer s.accountMutex.Unlock()

	if !ok {
		return Account{}, errors.New("no register")
	}

	return *account, nil
}

/*
func (s S) GetAccounts(sor AccountSortType) Account {
	accounts := make([]Account, len(s.Accounts))

	s.accountMutex.RLock()
	for _, account := range s.Accounts {
		accounts = append(accounts, account)
	}
	s.accountMutex.RUnlock()
	switch sor {
	case SortByName:
		sort.Slice(accounts, func(i, j int) bool { return accounts[i].Name < accounts[j].Name })
	case SortByNameReverse:
		sort.Slice(accounts, func(i, j int) bool { return accounts[i].Name > accounts[j].Name })
	case SortByBalance:
		sort.Slice(accounts, func(i, j int) bool { return accounts[i].Balance < accounts[j].Balance })
	}

	return accounts

}

*/
