package storage

import (
	"errors"
	"github.com/HuCuiGang/study_go/day3/user"
	"sync"
)

type Simple struct {
	mu sync.Locker
	db map[string]*user.BankAccount
}

//实现TodoListStorage
func NewSimple() TodoListStorage {
	return &Simple{
		db: map[string]*user.BankAccount{},
	}
}

//创建账户
func (s *Simple) OpenAccount(account, password string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.db[account] = user.New(account,password)
	return nil
}

//校验用户
func (s *Simple) CheckAccount(account, password string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	bankAccount,ex := s.db[account]
	if !ex {
		return errors.New("没有该账户")
	}
	if bankAccount.CheckPassword(password) {
		return nil
	}
	return errors.New("密码错误")
}

//查询余额
func (s Simple) BalanceEnquiry(account string) (balance float32, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	bankAccount,ex := s.db[account]
	if !ex {
		return 0, errors.New("没有该账户")
	}
	return bankAccount.BalanceEnquiry(),nil
}

//存钱
func (s Simple) Withdrawals(account string, money float32) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	bankAccount,ex := s.db[account]
	if !ex {
		return  errors.New("没有该账户")
	}
	return bankAccount.Withdrawals(money)
}

//取款

func (s Simple) Deposit(account string, money float32) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	bankAccount,ex := s.db[account]
	if !ex {
		return  errors.New("没有该账户")
	}
	bankAccount.Deposit(money)
	return nil
}
