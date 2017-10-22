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

- -conn 数据库地址.可以直输入-conn="233",相当于-conn="root:123456@tcp(127.0.0.1:3306)/zcm_crm?charset=utf8&loc=Asia%2FShanghai". 也.
- -dtag 是否在字段后面添加description标签,默认为fasle
- -comt 是否在字段后面添加注释,默认为fasle
- -table 表名.暂不支持多表
- 该命令会在当前目录下的models目录下创建代码文件
- 如果想使用自定义模板,在项目目录下创建model.tpl
- 如何写自定义模板?看代码~~~

### eg:
```
ztx -conn="root:123456@tcp(127.0.0.1:3306)/zcm_crm?charset=utf8&loc=Asia%2FShanghai" -table=activity -comt=true
```

## todo
- 自动生成单元测试代码
