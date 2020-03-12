package shop

import (
	"errors"
	"sort"
	"sync"
	"time"
)

var AccountTypeStruct = map[AccountType]struct{}{
	AccountNormal:  {},
	AccountPremium: {},
}

type AccountMutex struct {
	Account map[string]Account
	time.Duration
	sync.RWMutex
}

func NewAccount(Name string) Account {
	return Account{
		Name:    Name,
		Balance: 0,
		Type:    AccountNormal,
	}
}
func (accountMutex AccountMutex) getAccount(name string) (Account, error) {
	accountMutex.Lock()
	defer accountMutex.Unlock()
	account, err := accountMutex.Account[name]
	if !err {
		return Account{}, errors.New("username no exists")
	}
	return account, nil
}

func (accountMutex *AccountMutex) setAccount(account Account) {
	accountMutex.Lock()

	accountMutex.Account[account.Name] = account

	accountMutex.Unlock()
}

func (accountMutex *AccountMutex) Register(name string) error {
	err := CheckNameAccount(name)
	if err != nil {
		return err
	}
	if _, err := accountMutex.getAccount(name); err == nil {
		return errors.New("users exist")
	}
	accountMutex.setAccount(NewAccount(name))
	return nil
}
func (accountMutex *AccountMutex) ModifyAccountType(name string, accountType AccountType) error {
	account, err := accountMutex.getAccount(name)
	if err != nil {
		return err
	}
	if _, ok := AccountTypeStruct[accountType]; !ok {
		return errors.New("no type account")
	}
	account.Type = accountType
	accountMutex.setAccount(account)
	return nil
}

func (accountMutex *AccountMutex) AddBalance(name string, cash float32) error {
	account, err := accountMutex.getAccount(name)
	if err != nil {
		return err
	}
	if cash < 0 {
		return errors.New("negative cash")
	}
	account.Balance += cash
	accountMutex.setAccount(account)
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
