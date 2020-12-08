package user

import "errors"

type BankAccount struct {
	account string
	password string
	balance float32
}

func New(account,password string) *BankAccount {
	return &BankAccount{account: account ,password: password}
}

func (b *BankAccount) CheckPassword(password string) bool  {
	if password == b.password {
		return true
	}
	return false
}

//取钱
func (b *BankAccount) Withdrawals(money float32) error {
	if b.balance < money {
		errors.New("余额不足")
	}
	b.balance -=money
	return nil
}

//存钱
func (b *BankAccount) Deposit(money float32)  {
	b.balance += money
}

//查询余额
func (b *BankAccount) BalanceEnquiry() float32 {
	return b.balance
}
