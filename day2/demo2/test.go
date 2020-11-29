package main

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

// code:
type user struct {
	Name     string
	Password string
	Balance  float32
}

//func main() {
//	data := []user{
//		{
//			Name:     "dollarkiller",
//			Password: "123456",
//			Balance:  180,
//		},
//		{
//			Name:     "gangge",
//			Password: "123456",
//			Balance:  180,
//		},
//	}
//
//	// 下面用户登录 并取款100
//	name := "dollarkiller"
//	passwd := "123456"
//
//	// s1,s2 := range data s1,s2不是data中数据 是缓存
//	for idx, value := range data {
//		if value.Name == name && value.Password == passwd {
//			fmt.Println("登录成功")
//			value.Balance -= value.Balance - 100 //取款100
//			data[idx] = value
//		}
//	}
//
//	//
//
//	for _, value := range data {
//		fmt.Println("用户名：", value.Name)
//		fmt.Println("余额：", value.Balance)
//	}
//
//	//var users map[user]string
//
//	users := map[string]user{
//		"dollarkiller": {
//			Name:     "dollarkiller",
//			Password: "123456",
//			Balance:  180,
//		},
//		"gangge": {
//			Name:     "gangge",
//			Password: "123456",
//			Balance:  0,
//		},
//	}
//
//	v,ex := users[name]
//	if !ex {
//		fmt.Println("用户不存在")
//		return
//	}
//	if v.Name == name && v.Password == passwd{
//		fmt.Println("登录成功")
//		v.Balance = v.Balance - 100
//		users[name] = v
//	}
//
//	fmt.Println(users)
//}

