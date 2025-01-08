package wallet

import "testing"

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("Want %f bug got %f", want, got)
	}
}

func TestWalletDeposit(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(10.0))

	assertBalance(t, wallet, Bitcoin(10.0))
}

func TestWalletWithdraw(t *testing.T) {

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10.0))
		wallet.Withdraw(Bitcoin(5.0))

		assertBalance(t, wallet, Bitcoin(5.0))
	})

	t.Run("withdraw with not enough money", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10.0))
		err := wallet.Withdraw(Bitcoin(15.0))

		assertBalance(t, wallet, Bitcoin(10.0))

		if err == nil {
			t.Fatal("You shouldn't be able to withdraw more than you have")
		}
	})
}

func TestBitcoinToString(t *testing.T) {

	btc := Bitcoin(10.0)

	got := btc.String()
	want := "10.000000 BTC"
	if got != want {
		t.Errorf("Want %s bug got %s", want, got)
	}

}
