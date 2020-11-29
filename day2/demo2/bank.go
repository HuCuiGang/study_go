package main

import (
	"errors"
	"fmt"
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

type Bank struct {
	users map[string]*User
	currentLoginUser string

	mu sync.Mutex
}
var B *Bank // 全局函数 未初始化

func init()  {
	B = NewBanks() //初始化银行
}

//创建银行
func NewBanks() *Bank {
	return &Bank{
		users: map[string]*User{},
	}
}

//注册用户
func (b *Bank) AddBankUser(username string,password string) error{
	b.mu.Lock()
	defer b.mu.Unlock()
	_,ex := b.users[username]
	if ex {
		return errors.New("用户已存在")
	}
	user := CreateUser(username,password)
	b.users[user.username] = user
	return nil
}
//登录
func (b *Bank) LoginBank(username string,password string) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	v,ex := b.users[username]
	if !ex {
		return errors.New("用户名不存在")
	}
	if v.password !=password{
		return errors.New("密码错误")
	}
	b.currentLoginUser = username
	fmt.Println("登录成功")
	return nil
}
//查询余额
func (b *Bank) FindBankBalance() (float32,error){
	b.mu.Lock()
	defer b.mu.Unlock()
	balance := b.users[b.currentLoginUser].balance
	return balance,nil
}
//取钱
func (b *Bank) CutBankMoney(money float32) error  {
	if b.currentLoginUser == "" {
		return errors.New("用户未登录")
	}
	b.mu.Lock()
	defer b.mu.Unlock()
	ex := b.users[b.currentLoginUser].CutMoney(money)
	if ex!=nil {
		return ex
	}
	return nil
}
//存钱
func (b *Bank)SaveBankMoney(money float32) error {
	if b.currentLoginUser == ""{
		return errors.New("用户未登录")
	}
	b.mu.Lock()
	defer b.mu.Unlock()
	ex :=b.users[b.currentLoginUser].SaveMoney(money)
	if ex!=nil {
		return ex
	}
	return nil
}
