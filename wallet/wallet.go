package wallet

import (
	"errors"
	"fmt"
)

// 使 int 具备描述性
type Bitcoin int

// 将其定义为错误类型
var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")

type Wallet struct {
	balance Bitcoin
}

// 类型别名特性：声明方法
type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

//func (w Wallet) Deposit(amount int) {
//	fmt.Println("balance in Deposit is", &w.balance)
//	w.balance += amount
//}

// 避免上述情况，用 *指针* 解决
func (w *Wallet) Deposit(amount Bitcoin) {
	//fmt.Println("balance in Deposit is", &w.balance)
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error{
	if amount > w.balance {
		return InsufficientFundsError
	}
	w.balance -= amount
	return nil
}

func (w Wallet) Balance() Bitcoin {
	return w.balance
}
