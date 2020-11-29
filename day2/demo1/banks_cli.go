package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("欢迎来到小刚银行")

	log.SetFlags(log.Llongfile | log.LstdFlags)
loop:
	for {
		printMenu()
		fmt.Print("请输入：")
		var i string
		_, err := fmt.Scan(&i)
		if err != nil {
			log.Println("输入错误，请重新输入：")
			fmt.Println("输入错误，请重新输入：")
			continue
		}
		switch i {
		case "EXT":
			break loop //退出至for循环

		case "1" :
			fmt.Println("请输入用户名:")
			var username string
			var password string
			_,err :=fmt.Scan(&username)
			if err!=nil{
				log.Println("输入错误，重新输入：")
				continue
			}
			fmt.Println("请输入密码:")
			_,err = fmt.Scan(&password)
			if err!=nil {
				log.Println("输入错误，重新输入：")
				continue
			}

			err = Bank.Login(username,password)
			if err!=nil {
				log.Println("登录失败，重新登录")
				log.Println(err)
				continue
			}
			bank(username)

		case "2":
			fmt.Println("请输入用户名：")
			var username string
			var password string
			_,err = fmt.Scan(&username)
			if err!=nil{
				log.Println("输入错误，重新输入：")
				continue
			}
			fmt.Println("请输入密码")
			_,err = fmt.Scan(&password)
			if err!=nil{
				log.Println("输入错误，重新输入：")
				continue
			}
			err = Bank.AddUser(username,password)
			if err!=nil {
				log.Println(err)
				return
			}
			fmt.Println("注册成功")
		default :
			fmt.Println("输入的啥玩意？")
		}
	}
	fmt.Println("祝您生活愉快！")
}
func bank(username string)  {
	for  {
		printMenu2()
		fmt.Println("请输入：")
		var i string
		_,err :=fmt.Scan(&i)
		if err!=nil {
			log.Println("输入错误，重新输入：")
			continue
		}
		switch i {
		case "1":
			balance,err := Bank.ObtainingBalance()
			if err!=nil{
				log.Println(err)
				return
			}
			fmt.Printf("用户: %s 余额： %f /n",username,balance)
		case "2":
			fmt.Printf("存多少？请输入：")
			var count float32
			_,err = fmt.Scan(&count)
			if err!=nil {
				log.Printf("输入错误,重新输入")
				continue
			}
			err = Bank.Deposit(count)
			if err!=nil{
				log.Println(err)
				return
			}
			fmt.Printf("存款成功")
		case "3":
			fmt.Printf("要取多少钱？请输入：")
			var count float32
			_,err = fmt.Scan(&count)
			if err!=nil {
				log.Printf("输入错误，重新输入：")
				continue
			}
			err = Bank.Withdrawals(count)
			if err!=nil{
				log.Println(err)
				continue
			}
			fmt.Println("取款完成！")
		case "EXT":
			return
		default :
			fmt.Println("输入的啥玩意？")
		}
	}
}

//打印目录
func printMenu() {
	fmt.Println()
	fmt.Println("=====================")
	fmt.Println("登入系统请输入：1")
	fmt.Println("注册请输入：2")
	fmt.Println("退出:EXT")
	fmt.Println("=====================")
}
func printMenu2()  {
	fmt.Println()
	fmt.Println("====================")
	fmt.Println("查询余额输入：1")
	fmt.Println("存钱输入：2")
	fmt.Println("取钱输入：3")
	fmt.Println("退出输入：EXT")
	fmt.Println("-===================")
}
