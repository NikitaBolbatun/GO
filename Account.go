package shop

import (
	"errors"
	"sort"
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
func (accountMutex *AccountMutex) ModifyAccountType(name string, accountType AccountType) error {
	accountMutex.Lock()
	defer accountMutex.Unlock()
	if _, ok := accountMutex.Account[name]; !ok {
		return errors.New("no register")
	}
	if _, ok := AccountTypeStruct[accountType]; !ok {
		return errors.New("no type account")
	}

	account := accountMutex.Account[name]
	account.Type = accountType
	return nil
}

func (accountMutex *AccountMutex) AddBalance(name string, cash float32) error {

	if cash < 0 {
		return errors.New("negative cash")
	}

	if _, ok := accountMutex.Account[name]; !ok {
		return errors.New("username not exists")
	}

	account := accountMutex.Account[name]
	account.Balance += cash

	accountMutex.Account[name] = account
	return nil
}

func (accountMutex *AccountMutex) Balance(name string) (float32, error) {
	accountMutex.Lock()
	defer accountMutex.Unlock()
	if _, ok := accountMutex.Account[name]; !ok {
		return 0, errors.New("no register")
	}

	return accountMutex.Account[name].Balance, nil
}

func (accountMutex *AccountMutex) GetAccount(name string) (Account, error) {
	account, ok := accountMutex.Account[name]

	accountMutex.Lock()
	defer accountMutex.Unlock()

	if !ok {
		return Account{}, errors.New("no register")
	}

	return account, nil
}

func (accountMutex *AccountMutex) GetAccounts(sor AccountSortType) []Account {
	accounts := make([]Account, len(accountMutex.Account))

	accountMutex.RLock()
	for _, account := range accountMutex.Account {
		accounts = append(accounts, account)
	}
	accountMutex.RUnlock()
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
