package table

import (
	"database/sql"
	"errors"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

// func TestConnDb(t *testing.T) {
// 	connStr := "rootttt:zcmlc20162016@tcp(192.168.1.233:3306)/zcm_crm?charset=utf8&loc=Asia%2FShanghai"
// 	db := ConnDb(connStr)
// 	defer db.Close()
// }

func TestGetTableObject(t *testing.T) {

	connStr := "root:zcmlc20162016@tcp(192.168.1.233:3306)/zcm_crm?charset=utf8&loc=Asia%2FShanghai"
	conn := ConnDb(connStr)
	defer conn.Close()

	type args struct {
		db     *sql.DB
		tbname string
	}
	tests := []struct {
		name string
		args args
		want *Table
	}{
		// TODO: Add test cases.
		{args: args{db: conn, tbname: "activity"}, want: &Table{Name: "activity", Pk: "id"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTableObject(tt.args.db, tt.args.tbname); !compare1(got, tt.want) {
				t.Error(got.Name == tt.want.Name, got.Pk == tt.want.Pk)
				t.Errorf("GetTableObject() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func compare1(tbx, tby *Table) bool {
	if tbx.Name != tby.Name || tbx.Pk != tby.Pk {
		return false
	}
	return true
}

func compare2(tbx *Table) error {
	if len(tbx.Columns) != 9 {
		return errors.New("长度不一致")
	}
	// for _, v := range tbx.Columns {
	// 	fmt.Printf("%+v\n", v)
	// }
	return nil
}

func TestGetColumns(t *testing.T) {

	// db
	connStr := "root:zcmlc20162016@tcp(192.168.1.233:3306)/zcm_crm?charset=utf8&loc=Asia%2FShanghai"
	conn := ConnDb(connStr)
	defer conn.Close()

	// table
	table := GetTableObject(conn, "activity")

	type args struct {
		db    *sql.DB
		table *Table
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{args: args{db: conn, table: table}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetColumns(tt.args.db, tt.args.table)
			if err := compare2(tt.args.table); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestWriteModel(t *testing.T) {
	// db
	connStr := "root:zcmlc20162016@tcp(192.168.1.233:3306)/zcm_crm?charset=utf8&loc=Asia%2FShanghai"
	conn := ConnDb(connStr)
	defer conn.Close()

	// table
	table := GetTableObject(conn, "activity")

	GetColumns(conn, table)

	tables := []*Table{}
	tables = append(tables, table)
	WriteModelFiles(tables, ".", nil)
}
