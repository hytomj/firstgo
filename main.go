package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//建立连接
	db, err := sql.Open("mysql", "root:1q2w@tcp(192.168.1.189:3306)/testdb?charset=utf8")
	checkErr(err)

	//插入数据
	stmt, err := db.Prepare("INSERT worktime set start_time=?,end_time=?")
	checkErr(err)

	res, err := stmt.Exec("00:00:00", "23:59:59")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	//查询数据
	rows, err := db.Query("select * from worktime")
	checkErr(err)
	for rows.Next() {
		var uid int
		var start_time string
		var end_time string
		err = rows.Scan(&uid, &start_time, &end_time)
		checkErr(err)
		fmt.Println(uid, start_time, end_time)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
