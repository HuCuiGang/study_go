package main

import (
	"errors"
	"sync"
)

// 1. 设计一个CLI的
// 银行系统 可以多个用户
// 基础操作
//登录
//存钱
//取钱
//获取余额

// 1.设计  2.写

// 1. 用户数据的存储   [user1,user2,user3]  , map[user_name]User
// 2. CLI与用户的交互
type User struct {
	username string
	password string
	balance float32
	mu sync.Mutex
}

func CreateUser(username string,password string) *User {
	// 有可能你会初始化其他的东西
	return &User{
		username: username,
		password: password,
	}
}

//校验用户名密码
func (u *User) LoginUser(username string ,password string) error  {
	u.mu.Lock()
	defer u.mu.Unlock()
	if u.username!=username  {
		return errors.New("用户名错误")
	}
	if u.password != password{
		return errors.New("密码错误")
	}
	return nil
}
//查询余额
func (u *User) FindBalance() ( float32,error){
	u.mu.Lock()
	defer u.mu.Unlock()
	return u.balance,nil
}
//存钱
func (u *User) SaveMoney(money float32) error {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.balance += money
	return nil
}
//取钱
func (u *User) CutMoney(money float32) error {
	u.mu.Lock()
	defer u.mu.Unlock()
	if money > u.balance {
		return errors.New("余额不足")
	}
	u.balance -= money
	return nil
}
