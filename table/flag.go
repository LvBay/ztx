package table

import (
	"flag"
	"fmt"
	"os"
)

type GenCommand struct {
	Flag flag.FlagSet
}

type GenConfig struct {
	Conn       string // 数据库连接地址
	Table      string // 表名
	DescTag    bool   // 是否添加description标签
	Annotation bool   // 是否添加注释
}

var (
	genCmd = GenCommand{}
	Config *GenConfig // 输出代码的配置
)
var connMap = map[string]string{
// 这里可以添加一些常用的数据库连接
}

var usageTmpl = `ztx can generate some code from mysql

USAGE: 

	ztx [OPTIONS]

OPTIONS:

	-c 数据库地址 
	-t 表名称
	-d 添加orm-descrition标签 默认false
	-a 添加注释 Annotation 默认true
`

func Usage() {
	fmt.Println(usageTmpl)
	os.Exit(2) // 如果用户输入参数不合法,打印之后直接退出.
	return
}

func init() {
	Config = &GenConfig{}
	genCmd.Flag.Usage = Usage                                                      // 覆盖默认的Usage,打印后直接退出程序
	genCmd.Flag.StringVar(&Config.Conn, "c", "", "数据库地址 connect")                  // 数据库地址 connect
	genCmd.Flag.StringVar(&Config.Table, "t", "", "表名称 table")                     // 表名称 table
	genCmd.Flag.BoolVar(&Config.DescTag, "d", false, "添加orm-descrition标签 默认false") // 添加orm-descrition标签
	genCmd.Flag.BoolVar(&Config.Annotation, "a", true, "添加注释 Annotation 默认true")   // 添加注释 Annotation
	genCmd.Flag.Parse(os.Args[1:])
	if v, ok := connMap[Config.Conn]; ok {
		Config.Conn = v
	}
}
