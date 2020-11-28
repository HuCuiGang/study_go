package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)
import "github.com/HuCuiGang/study_go/day1/utils"

type user struct{
	Name string `json:"name12"`
	Age int
	Sex bool
}
func main(){
	i :=utils.Add(1,1)
	fmt.Println("1+1=",i)

	body,err := ioutil.ReadFile("./day1/config.json")
	if err!=nil {
		panic(err)
	}
	user := user{}
	err = json.Unmarshal(body,&user)
	if err!=nil{
		panic(err)
	}

	fmt.Println(user)
}



