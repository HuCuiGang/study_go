package main

import (
	"github.com/HuCuiGang/study_go/day3/cli"
	"github.com/HuCuiGang/study_go/day3/conf"
)

func main()  {
	conf.InitConfig()

	c := cli.New()
	c.Run()

}