package shop

import (
	"errors"
	"sort"
)

func NewAccount(Name string) Account {
	return Account{
		Name:    Name,
		Balance: 0,
		Type:    AccountNormal,
	}
}

func (s S) Register(Name string) (int, error) {

	if _, ok := s.Accounts[Name]; ok {
		return 0, errors.New("username exists")
	}
	if len(Name) < 1 {
		return 0, errors.New("username not correct")
	}
	s.accountMutex.RLock()
	account := NewAccount(Name)
	s.accountMutex.RUnlock()
	s.Accounts[Name] = &account
	return 0, nil
}

func (s S) ModifyAccountType(Name string, accountType AccountType) error {

	if _, ok := s.Accounts[Name]; !ok {
		return errors.New("no register")
	}

	s.accountMutex.Lock()
	defer s.accountMutex.Unlock()

	account := s.Accounts[Name]
	account.Type = accountType
	return nil
}

func (s S) Balance(Name string) (float32, error) {

	if _, ok := s.Accounts[Name]; !ok {
		return 0, errors.New("no register")
	}

	s.accountMutex.Lock()
	defer s.accountMutex.Unlock()

	return s.Accounts[Name].Balance, nil
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

//noinspection GoTypesCompatibility
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
