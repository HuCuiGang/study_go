package main

import (
	"sync"
)

type User struct {
	username string
	password string
	balance float32

	mu sync.Mutex
}
func NewUser(username string,password string) *User {
	return &User{
		username: username,
		password: password,
	}
}

//校验PassWord
func (u *User) CheckPassword(password string) bool{
	u.mu.Lock()
	defer u.mu.Unlock()
	if u.password == password {
		return true
	}
	return false
}


//存款
func ( u *User) Deposit(money float32)  {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.balance += money
}
//取款
func (u *User) Withdrawals (money float32) bool {
	u.mu.Lock()
	defer u.mu.Unlock()
	if u.balance < money{
		return false
	}
	u.balance -= money
	return true
}
//获取余额
func (u *User) ObtainingBalance() float32{
	u.mu.Lock()
	defer u.mu.Unlock()
	return u.balance
}
//------------------------------------------------------------------


