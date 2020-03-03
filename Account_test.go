package shop

import (
	"testing"
)

func Test_Register(t *testing.T) {
	shop := NewShop()

	nameRight := "Nikita"

	if err := shop.Accounts[nameRight]; err != nil {
		t.Fatalf("Register() error = %v", err)
	}
}

func Test_RegisterTwoRegsName(t *testing.T) {
	shop := NewShop()

	name := "Nikita"
	nameTwo := "Nikita"

	shop.Register(name)
	if _, ok := shop.Accounts[name]; !ok {
		t.Fatal("Register error")
	}
	if err := shop.Register(nameTwo); err == nil {
		t.Fatalf("Register two = %v", err)
	}
}

func Test_RegisterNameEmpty(t *testing.T) {
	shop := NewShop()
	nameRight := ""
	err := shop.Register(nameRight)

	if _, ok := shop.Accounts[nameRight]; ok {
		t.Fatalf("Register error = %v", err)
	}
}

func Test_RegisterNameSpace(t *testing.T) {
	shop := NewShop()
	nameRight := " "
	shop.Register(nameRight)

	if err := shop.Accounts[nameRight]; err != nil {
		t.Fatalf("Register error = %v", err)
	}
}

func Test_AddBalance(t *testing.T) {
	shop := NewShop()
	shop.Register("Nikita")

	if err := shop.AddBalance("Nikita", 1000); err != nil {
		t.Fatalf("Not correct add = %v", err)
	}
	if shop.Accounts["Nikita"].Balance != 1000 {
		t.Fatal("Not correct add")
	}
}

func Test_AddBalanceNegative(t *testing.T) {
	shop := NewShop()
	shop.Register("Nikita")

	if err := shop.AddBalance("Nikita", -1000); err == nil {
		t.Fatalf("Not correct add = %v", err)
	}
}
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
	if err == nil {
		t.Fatalf("Type account = %v", err)
	}
}

//-----GetAccounts-----
