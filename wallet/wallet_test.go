package wallet

import (
	"testing"
)

// 保证安全性：使用方法获取内部属性
func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		// 储值
		wallet.Deposit(Bitcoin(10))

		// 获取余额
		//got := wallet.Balance()
		want := Bitcoin(10)

		// 通过 &val 找到内存中的值
		// 运行发现两个值的内存地址并不相同，在 Deposit 中处理的是来自测试的 *副本*
		//fmt.Println("testing got is", &wallet.balance)
		assertBalance(t, wallet, want)
	})

	// 与上面是两套测试环境？
	// 使用 errcheck 检查出，未处理的错误，增加 assertNoError
	t.Run("Withdraw with funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}

		// 取钱
		err := wallet.Withdraw(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(10)
		// 初始化
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		// 即 null
		// 如果访问 nil，将会触发运行时 panic
		//if err == nil {
		//	t.Error("wanted an error but didn't get one")
		//}
		assertError(t, err, InsufficientFundsError)
	})
}

func assertError(t *testing.T, got error, want error) {
	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	if got != nil {
		t.Fatal("got an error but didnt want one")
	}
}

func assertBalance(t *testing.T, got Wallet, want Bitcoin) {
	balance := got.Balance()
	if balance != want {
		t.Errorf("got %s want %s", balance, want)
	}
}