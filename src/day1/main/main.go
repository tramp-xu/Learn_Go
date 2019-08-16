package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func checkErr(err error)  {
	if err != nil {
		panic(err)
	}
}

func main () {
	db, err := sql.Open("mysql", "root:xuuu@911@tcp(localhost:3306)/tb_test")
	checkErr(err)

	// 查询数据
	rows, err := db.Query("select * from emp")
	checkErr(err)

	for rows.Next() {
		var id int
		var name string
		var deptId int
		var salary float32
		err = rows.Scan(&id, &name, &deptId, &salary)
		checkErr(err)
		fmt.Println(id)
		fmt.Println(name)
		fmt.Println(deptId)
		fmt.Println(salary)
	}
}