package main

import (
	"errors"
	"fmt"
	"sync"
)

type Banks struct {
	users map[string]*User // 用户必须采用指针 这样才能保证我们修改和访问的数据都是同一块内存
	logUser string //保存当前登录的用户名
	mu sync.Mutex
}
var Bank *Banks // 全局函数 未初始化

func init()  {
	Bank = NewBanks() //初始化银行
}

//创建银行
func NewBanks() *Banks {
	return &Banks{
		users: map[string]*User{},
	}
}

//添加用户
func (b *Banks) AddUser(username string,password string) error{
	if b.logUser == ""{
		return errors.New("未登录系统")
	}
	b.mu.Lock()
	defer b.mu.Unlock()

	_,ex := b.users[username] // 如果这样去取MAP data,ex := 会有两个返回 data为数据,ex为改数据是否存在
	if ex {
		return errors.New("该用户已存在")
	}
	//用户不存在就创建
	user := NewUser(username,password)
	b.users[username] = user
	return nil //没有错误返回nil
}
//登录
func (b *Banks)Login(username string,password string) error{
	b.mu.Lock()
	defer b.mu.Unlock()
	user,ex := b.users[username]
	if !ex {
		errors.New("该用户不存在")
	}
	if !user.CheckPassword(password) {
		errors.New("密码不正确")
	}
	b.logUser=username
	return nil
}

//获取余额
func (b *Banks) ObtainingBalance() (float32,error) {
	if b.logUser=="" {
		return 0,errors.New("未登录系统")
	}
	b.mu.Lock()
	defer b.mu.Unlock()
	user := b.users[b.logUser]
	return user.ObtainingBalance(),nil
}

//存款
func (b *Banks) Deposit(money float32) error {
	if b.logUser=="" {
		errors.New("未登录系统")
	}
	b.mu.Lock()
	defer b.mu.Unlock()
	b.users[b.logUser].Deposit(money)
	return nil
}

func (b *Banks)Withdrawals(money float32) error  {
	if b.logUser==""  {
		return fmt.Errorf("%s","未登录系统")
	}
	b.mu.Lock()
	defer b.mu.Unlock()
	if !b.users[b.logUser].Withdrawals(money) {
		fmt.Errorf("%s","余额不足")
	}
	return nil
}

//// i>10
//func add(i int) (int, error) {
//	if i < 10 {
//		return 0, fmt.Errorf("i > 10")
//		//panic("a > 10")
//	}
//}