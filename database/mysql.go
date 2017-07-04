package mysql

import (
	"database/sql"
	"fmt"
	_ "mysql"
)

func TestMySQL() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/exp3?charset=utf8")
	defer db.Close()
	if nil != err {
		fmt.Println(err)
	}
	// query
	rows, err := db.Query("SELECT * FROM student")
	if nil != err {
		fmt.Println(err)
	}

	for rows.Next() {
		var stuid string
		var name string
		var sex string
		var age int
		var birthday string
		var info string
		fmt.Println(rows.Columns())
		err = rows.Scan(&stuid, &name, &sex, &age, &birthday, &info)
		fmt.Println(stuid, "\t\t", name)
	}
	defer rows.Close()
}
