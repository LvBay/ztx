## 工具包,目的:节省时间.

## 安装
```
go get github.com/beego/bee
cd GOPATH
git clone git@github.com:LvBay/ztx.git
cd ztx
go install
```

## 目前支持的功能:

### 根据表结构自动生成相关代码,支持自定义模板.主要是在bee工具的基础上稍作修改.相关用法:

- -c 数据库地址.eg: ```-c="root:123456@tcp(127.0.0.1:3306)/test?charset=utf8&loc=Asia%2FShanghai"``` 默认127.0.0.1:3306
- -d 添加orm-descrition标签,默认为fasle
- -a 添加注释,默认为true
- -t 表名.暂不支持多表
- 该命令会在当前目录下的models目录下创建代码文件
- 如果想使用自定义模板,在项目目录下创建model.tpl

### eg:
```
ztx -t=test
```

## todo
- 自动生成单元测试代码
