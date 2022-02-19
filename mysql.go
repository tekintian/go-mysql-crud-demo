// mysql.go
/**
go mysql 数据库增删改查 示例!
*/
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "test:test888@tcp(127.0.0.1:3306)/test?charset=utf8mb4")
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT t_user  SET pwd=?,username=?,create_time=?")

	res, err := stmt.Exec("zhangsan", "技术部", "2016-12-09")
	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println(id)

	//select
	rs, err := db.Query("select id,pwd,username,create_time from t_user where id>0")
	if err != nil {
		panic(err)
	}
	//循环输出查询结果
	for rs.Next() {
		var id int
		var pwd string
		var username string
		var create_time string
		err = rs.Scan(&id, &pwd, &username, &create_time)
		fmt.Println(id, pwd, username, create_time)
	}

	//update
	upsmt, err := db.Prepare("update t_user set pwd=?, username=? where id=?")
	af, err := upsmt.Exec("888888", "Tekin", id)
	if err != nil {
		panic(err)
	} else {
		n, err := af.RowsAffected()
		if err != nil {
			fmt.Println("更新失败!", err)
		}
		fmt.Printf("用户ID为 %d 的数据修改成功! \n 影响行数 %d", id, n)
	}

	//delete
	//stmt, err = db.Prepare("delete from t_user where id=?")
	//res, err = stmt.Exec(id)
	//if err != nil {
	//	panic(err)
	//} else {
	//	fmt.Println("数据ID", id, "删除成功!")
	//}

}
