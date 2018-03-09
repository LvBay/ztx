package table

import (
	"database/sql"

	beeLogger "github.com/beego/bee/logger"
)

func ConnDb(conn string) *sql.DB {
	db, err := sql.Open("mysql", conn)
	// 即使conn连接失败,err还是为空
	if err != nil {
		beeLogger.Log.Fatalf("Could not connect to mysql database using '%s': %s", conn, err)
	}
	return db
}
