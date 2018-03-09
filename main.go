package main

import (
	"fmt"

	"github.com/LvBay/ztx/table"

	"github.com/beego/bee/logger/colors"
)

func main() {

	ok := table.GenCode()
	if ok {
		fmt.Println(colors.Green(FLAG))
	} else {
		fmt.Println(colors.Black(FLAG))
	}
}

const FLAG string = `
		＿＿＿
	　　 ／　　　▲
	／￣　 ヽ　■■
	●　　　　　■■
	ヽ＿＿＿　　■■
	　　　　）＝｜
	　　　／　｜｜
	　∩∩＿＿とﾉ
	　しし———
`
