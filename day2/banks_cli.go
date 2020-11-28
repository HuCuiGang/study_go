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
			_,err = fmt.Scan(&password)
			if err!=nil {
				log.Println("输入错误，重新输入：")
				continue
			}

			err = Bank.Login(username,password)
			if err!=nil {
				log.Println("登录失败，重新登录")
				continue
			}


		}



	}

}

func printMenu() {
	fmt.Println()
	fmt.Println("=====================")
	fmt.Println("登入系统请输入：1")
	fmt.Println("注册请输入：2")
	fmt.Println("退出:EXT")
	fmt.Println("=====================")
}
