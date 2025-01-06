package wallet

import "fmt"

type Bitcoin float64

type Stringer interface {
	String() string
}

func (btc Bitcoin) String() string {
	return fmt.Sprintf("%f BTC", btc)
}

type Wallet struct {
	balance Bitcoin
}

func (wallet *Wallet) Deposit(amount Bitcoin) {
	wallet.balance += amount
}

func (wallet *Wallet) Withdraw(amount Bitcoin) {
	wallet.balance -= amount
}

func (wallet *Wallet) Balance() Bitcoin {
	return wallet.balance
}
