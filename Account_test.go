package shop

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

//-----ModifyAccount-----

func Test_ModifyAccountType(t *testing.T) {
	shop := NewShop()

	account := Account{
		Name:    "Petr",
		Balance: 0,
		Type:    AccountNormal,
	}

	shop.Accounts["Petr"] = &account

	// не предусмотрен такой тип аккаунта
	err := shop.ModifyAccountType("Petr", 101)
	assert.NotZero(t, err)
}

//-----GetAccounts-----

func Test_GetAccounts(t *testing.T) {
	// вот бы хотя бы оставить эти элементы, не то что отсортировать
	t.Run("SortByBalance", GetAccountsSortByBalance)
	t.Run("SortByName", GetAccountsSortByName)
	t.Run("SortByNameReverse", GetAccountsSortByNameReverse)
}

func GetAccountsSortByBalance(t *testing.T) {
	shop := NewShop()

	acc1 := Account{
		Name:    "Petr",
		Balance: 0,
		Type:    AccountNormal,
	}
	acc2 := Account{
		Name:    "Adams",
		Balance: 10,
		Type:    AccountNormal,
	}
	shop.Accounts["Petr"] = &acc1
	shop.Accounts["Adams"] = &acc2

	// элементы исчезли
	result := shop.GetAccounts(SortByBalance)
	require.Equal(t, []Account{acc1, acc2}, result)
}

func GetAccountsSortByName(t *testing.T) {
	shop := NewShop()

	acc1 := Account{
		Name:    "Petr",
		Balance: 0,
		Type:    AccountNormal,
	}
	acc2 := Account{
		Name:    "Adams",
		Balance: 10,
		Type:    AccountNormal,
	}
	shop.Accounts["Petr"] = &acc1
	shop.Accounts["Adams"] = &acc2

	result := shop.GetAccounts(SortByName)
	require.Equal(t, []Account{acc2, acc1}, result)
}

func GetAccountsSortByNameReverse(t *testing.T) {
	shop := NewShop()

	acc1 := Account{
		Name:    "Petr",
		Balance: 0,
		Type:    AccountNormal,
	}
	acc2 := Account{
		Name:    "Adams",
		Balance: 10,
		Type:    AccountNormal,
	}
	shop.Accounts["Petr"] = &acc1
	shop.Accounts["Adams"] = &acc2

	result := shop.GetAccounts(SortByNameReverse)
	require.Equal(t, []Account{acc1, acc2}, result)
}

func Test_GetAccountsSortByUnknowingType(t *testing.T) {
	shop := NewShop()

	acc1 := Account{
		Name:    "Petr",
		Balance: 0,
		Type:    AccountNormal,
	}
	acc2 := Account{
		Name:    "Adams",
		Balance: 10,
		Type:    AccountNormal,
	}
	shop.Accounts["Petr"] = &acc1
	shop.Accounts["Adams"] = &acc2

	// не знаю ошибка ли, подал непонятно что в метод, а он вернул какие-то значения
	result := shop.GetAccounts(10)
	assert.Empty(t, result)
}
