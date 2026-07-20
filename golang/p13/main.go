package main

import "fmt"

type BankAccount struct {
	Owner   string
	Balance float64
}

func (b *BankAccount) Deposit(amount float64) {
	b.Balance += amount
}
func (b *BankAccount) Withdraw(amount float64) {
	b.Balance -= amount
}
func (b *BankAccount) PrintBalance() float64 {
	fmt.Println(b.Balance)
	return b.Balance
}

func main() {
	account := BankAccount{Owner: "John", Balance: 1000.0}
	account.PrintBalance()
	account.Deposit(500.0)
	account.PrintBalance()
	account.Withdraw(200.0)
	account.PrintBalance()
}
