package main

import "fmt"

type Account struct {
	AccountNo        string
	AccountStatement []AccountStatement
	Statement        string
}

type AccountStatement struct {
	date    string
	amount  float32
	balance float32
}

func Deposit(account Account, amount int) float32 {
	account.deposit(amount)
	size := len(account.AccountStatement) - 1
	balance := account.AccountStatement[size].balance
	return float32(balance)
}

func Withdraw(account Account, amount int) float32 {
	account.withdraw(amount)
	last := len(account.AccountStatement) - 1
	return float32(account.AccountStatement[last].balance)
}

func PrintStatement(account Account) string {
	account.printStatement()
	return account.Statement
}

func (b *Account) deposit(depositAmount int) {
	balance := b.AccountStatement[0].balance
	for i := range b.AccountStatement {
		b.AccountStatement[i].amount = float32(depositAmount)
		balance += float32(depositAmount)
		b.AccountStatement[i].balance += balance
	}
}

func (b *Account) withdraw(amount int) {
	balance := b.AccountStatement[0].balance
	for i := range b.AccountStatement {
		b.AccountStatement[i].amount = float32(-amount)
		balance -= float32(amount)
		b.AccountStatement[i].balance = balance
	}
}

func (b *Account) printStatement() {
	statement := "DATE-------|AMOUNT-----|BALANCE----"

	for i := len(b.AccountStatement) - 1; i >= 0; i-- {
		acc := b.AccountStatement[i]
		amount := fmt.Sprintf("%.2f", acc.amount)
		balance := fmt.Sprintf("%.2f", acc.balance)
		statement += fmt.Sprintf("\n%v-|%v----|%v----", acc.date, amount, balance)
	}
	b.Statement = statement
}
