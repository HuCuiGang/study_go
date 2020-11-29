package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("欢迎来到小刚银行")
	log.SetFlags(log.Llongfile | log.LstdFlags) // 设置log输入  log.Llongfile 完整的文件路径   log.LstdFlags 显示行号
loop:  //标识符
	for  {
		Cli1()
		fmt.Print("请输入:")
		var i string
		_,err :=fmt.Scan(&i)
		if err!=nil {
			log.Println("输入格式不正确,重新输入:")
			continue
		}
		switch i {
		case "1":
			fmt.Println("请输入用户名")
			var username string
			var password string
			_,err := fmt.Scan(&username)
			if err!=nil {
				log.Println("输入格式不正确,重新输入:")
				continue
			}
			_,err = fmt.Scan(&password)
			if err!=nil {
				log.Println("输入格式不正确,重新输入:")
				continue
			}
			err = B.LoginBank(username,password)
			if err!=nil {
				log.Println(err)
				log.Println("请重新登录!")
				continue
			}
			enter(username)

		case "2":
			fmt.Println("请输入用户名:")
			var username string
			_,err := fmt.Scan(&username)
			if err!=nil {
				log.Println("输入格式不正确,重新输入:")
				continue
			}
			fmt.Println("请输入密码:")
			var password string
			_,err =fmt.Scan(&password)
			if err!=nil {
				fmt.Println("输入格式不正确,请重新输入:")
				continue
			}
			err = B.AddBankUser(username,password)
			if err!=nil {
				log.Println(err)
				return
			}
			fmt.Println("注册成功!")

		case "EXT":
			break loop //退出至for循环
		default:
			fmt.Println("输入格式不正确,请重新输入:")
		}


		log.Println("=======================")
	}
}

func enter(useranme string )  {
	for {
		Cli2()
		fmt.Print("请输入:")
		var i string
		_,err :=fmt.Scan(&i)
		if err!=nil {
			log.Println("输入格式不正确,重新输入:")
			continue
		}
		switch i {
		case "1":
			balance,err := B.FindBankBalance()
			if err!=nil {
				log.Println(err)
				return
			}
			fmt.Println("用户名",useranme + "余额:",balance)

		case "2":
			fmt.Println("存多少?请输入:")
			var money float32
			_,err := fmt.Scan(&money)
			if err!=nil {
				log.Println("输入格式不正确,重新输入:")
				continue
			}
			err = B.SaveBankMoney(money)
			if err!=nil {
				log.Println(err)
				return
			}
			fmt.Println("存款成功!")

		case "3":
			fmt.Println("您要取多少钱?请输入:")
			var money float32
			_,err := fmt.Scan(&money)
			if err!=nil {
				log.Println("输入格式不正确,重新输入:")
				continue
			}
			err = B.CutBankMoney(money)
			if err!=nil {
				log.Println(err)
				return
			}
			fmt.Println("取款完成")

		case "EXT":
			return

		default:
			fmt.Println("输入格式不正确,重新输入:")
			continue
		}

	}
}

func Cli1(){
	fmt.Println()
	fmt.Println("============================")
	fmt.Println("登入系统请输入1")
	fmt.Println("注册请输入2")
	fmt.Println("退出输入EXT")
	fmt.Println("==============================")
}
func Cli2()  {
	fmt.Println()
	fmt.Println("==============================")
	fmt.Println("查询余额请输入1")
	fmt.Println("存钱请输入:1")
	fmt.Println("取前请输入2")
	fmt.Println("退出输入EXT")
	fmt.Println("==============================")
}
