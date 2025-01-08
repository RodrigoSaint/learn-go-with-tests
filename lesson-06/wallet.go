package wallet

import (
	"errors"
	"fmt"
)

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

func (wallet *Wallet) Withdraw(amount Bitcoin) error {

	if amount > wallet.balance {
		return errors.New("Insuficient funds")
	}
	wallet.balance -= amount
	return nil
}

func (wallet *Wallet) Balance() Bitcoin {
	return wallet.balance
}
