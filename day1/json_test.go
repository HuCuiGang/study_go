package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

type user2 struct{
	Name string `json:"name12" yaml:"name12" xml:"name12"`
	Age int	`json:"age"`
	Sex bool
}

func TestJson(t *testing.T)  {
	body,err :=ioutil.ReadFile("./config.json")
	if err!=nil {
		panic(err)
	}
	//fmt.Println(string(body))
	user := user2{}
	err = json.Unmarshal(body,&user)
	if err!= nil {
		panic(err)
	}
	fmt.Println(user)
}