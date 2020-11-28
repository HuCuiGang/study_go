package main

import (
	"fmt"
	"io/ioutil"
)
import "gopkg.in/yaml.v3"

type user struct {
	Name string
	Age  int
	Sex  bool
}

func main() {
	body, err := ioutil.ReadFile("./day1/demo1/config.yaml")
	if err != nil {
		panic(err)
	}
	user :=user{}
	err = yaml.Unmarshal(body,&user)
	if err!=nil{
		panic(err)
	}
	fmt.Println(user)
	user.Age=20
	body,err = yaml.Marshal(&user)
	if err!=nil {
		panic(err)
	}
	err = ioutil.WriteFile("./day1/demo1/config.yaml",body,0600)
	if err!= nil{
		panic(err)
	}
	fmt.Println(string(body))
}
