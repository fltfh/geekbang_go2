package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"log"
)

type User struct {
	id   int
	name string
}

// DaoQueryRow dao层查询数据，ErrNoRows没有被Wrap到errors中的思考
// panic 属于栈溢出等系统严重性错误
// error 属于应用逻辑存在异常，或不可预见性问题，并且频率较低的情况
// ErrNoRows 是属于应用逻辑中可能情况的分支之一。并且可能出现频率较高，所以放到data中返回，是属于正常逻辑判断的一部分。

func DaoQueryRow(db *sql.DB) (*User, error) {
	row := db.QueryRow("select * from user WHERE id = ?", 1)
	user := &User{}
	err := row.Scan(&user.id, &user.name)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, errors.Wrap(err, "data not found")
	}
	return user, nil
}

func main() {

	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/cmdb")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Errorf("connect database error: %w", err)
		return
	}

	data, err := DaoQueryRow(db)
	if err != nil {
		fmt.Errorf("database query error: %w", err)
		return
	}

	if data == nil {
		fmt.Printf("database query data not found")
	} else {
		fmt.Printf("database query data is : %v", data)
	}
}

