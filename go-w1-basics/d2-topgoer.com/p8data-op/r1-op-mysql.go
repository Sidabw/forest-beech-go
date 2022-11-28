package data_op

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// import "github.com/go-sql-driver/mysql"
type User struct {
	Id    int    `db:"id"`
	Name  string `db:"name"`
	Age   int    `db:"age"`
	Email string `db:"email"`
}

var Db *sqlx.DB

func init() {
	// https://www.topgoer.com/数据库操作/go操作mysql/mysql使用.html
	database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/mysql8_first_db")
	if err != nil {
		fmt.Println("open mysql fail", err)
		return
	}
	// 不明白：Db, err := sqlx.Open,,, 为啥不行
	Db = database
}

func R1OpMysql() {
	defer Db.Close()
	insertLearn()
	// selectLearn()
	// updateLearn()
	// deleteLearn()
}

func insertLearn() {
	row, err := Db.Exec("insert into user(name, age, email) values (?, ?, ?)", "ha", 1111, "a@k.c")
	if err != nil {
		fmt.Println("insert fail", err)
		return
	}

	id, err := row.LastInsertId()
	if err != nil {
		fmt.Println("row.LastInsertId() invoke fail", err)
		return
	}
	fmt.Println("insert successful, id:", id)

}

func selectLearn() {
	var users []User
	err := Db.Select(&users, "select * from user where name = ?", "ha")
	if err != nil {
		fmt.Println("select fail", err)
		return
	}
	for i := range users {
		fmt.Println("user: ", users[i])
	}
}

func updateLearn() {
	exec, err := Db.Exec("update user set email = ? where name = ?", "a.a@b.b.ci", "ha")
	if err != nil {
		fmt.Println("update fail", err)
		return
	}
	affected, _ := exec.RowsAffected()
	// 第一次3，第二次0，后面。。。0
	fmt.Println("rows affected: ", affected)
}

func deleteLearn() {
	exec, _ := Db.Exec("delete from user where id = ?", 11)
	affected, _ := exec.RowsAffected()
	fmt.Println("deleted rows: ", affected)
}
