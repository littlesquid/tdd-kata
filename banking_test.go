package main

import (
	"fmt"
	"testing"
	"time"
)

func TestDepositToAccount(t *testing.T) {
	//given
	accountStatement := AccountStatement{
		date:    time.Now().Format("02/01/2006"),
		amount:  200,
		balance: 0,
	}

	account := Account{
		AccountNo:        "2032252290",
		AccountStatement: []AccountStatement{accountStatement},
	}

	amount := 200
	expected := float32(200)

	//when
	res := Deposit(account, amount)

	//then
	if float32(expected) != res {
		t.Error(fmt.Sprintf("balance should be %v but got %v", expected, res))
	}
}

func TestDepositToAccount2(t *testing.T) {
	//given
	accountStatement1 := AccountStatement{
		date: time.Now().Format("02/01/2006"),
	}
	accountStatement2 := AccountStatement{
		date: time.Now().Add(-24 * time.Hour).Format("02/01/2006"),
	}

	account := Account{
		AccountNo:        "2032252290",
		AccountStatement: []AccountStatement{accountStatement1, accountStatement2},
	}

	amount := 350
	expected := float32(700)

	//when
	res := Deposit(account, amount)

	//then
	if float32(expected) != res {
		t.Error(fmt.Sprintf("balance should be %v but got %v", expected, res))
	}
}

func TestWithdrawAccount(t *testing.T) {
	//given
	accountStatement := AccountStatement{
		date:    time.Now().Format("02/01/2006"),
		balance: 500,
	}
	account := Account{
		AccountNo:        "2032252290",
		AccountStatement: []AccountStatement{accountStatement},
	}
	amount := 300
	expected := float32(200)
	//when
	res := Withdraw(account, amount)

	//then
	if float32(expected) != res {
		t.Error(fmt.Sprintf("balance should be %v but got %v", expected, res))
	}
}

func TestWithdrawAccount2(t *testing.T) {
	//given
	accountStatement1 := AccountStatement{
		date:    time.Now().Format("02/01/2006"),
		balance: 500,
	}
	accountStatement2 := AccountStatement{
		date: time.Now().Format("02/01/2006"),
	}
	account := Account{
		AccountNo:        "2032252290",
		AccountStatement: []AccountStatement{accountStatement1, accountStatement2},
	}
	amount := 100
	expected := float32(300)
	//when
	res := Withdraw(account, amount)

	//then
	if float32(expected) != res {
		t.Error(fmt.Sprintf("balance should be %v but got %v", expected, res))
	}
}

func TestPrintStatement(t *testing.T) {
	//given
	accountStatement1 := AccountStatement{
		date:    "01/04/2014",
		amount:  1000,
		balance: 1000,
	}
	accountStatement2 := AccountStatement{
		date:    "02/04/2014",
		amount:  -100,
		balance: 900,
	}
	accountStatement3 := AccountStatement{
		date:    "10/04/2014",
		amount:  500,
		balance: 1400,
	}

	account := Account{
		AccountNo:        "2032252290",
		AccountStatement: []AccountStatement{accountStatement1, accountStatement2, accountStatement3},
	}

	expected := "DATE-------|AMOUNT-----|BALANCE----\n10/04/2014-|500.00----|1400.00----\n02/04/2014-|-100.00----|900.00----\n01/04/2014-|1000.00----|1000.00----"

	//when
	statement := PrintStatement(account)

	//then
	if expected != statement {
		t.Error(fmt.Sprintf("balance should be %v but got %v", expected, statement))
	}
}
