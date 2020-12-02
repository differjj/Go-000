package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	xerrors "github.com/pkg/errors"
	"log"
)

func main() {
	var id int64 = 2
	name, err := GetUserInfo(id)

	if err != nil {
		log.Printf("main: param:id=%v\nerr:%+v\n",id, err)
		if errors.Is(err,sql.ErrNoRows){
			fmt.Println("not found user info")
		} else {
			fmt.Println("failed")
		}
	} else {
		fmt.Printf("username is: %v",name)
	}

}


//service
func GetUserInfo(id int64) (string, error) {
	return UserInfo(id)
}


//dao
var db *sql.DB

func UserInfo(id int64) (string, error) {
	var name string
	err := db.QueryRow("SELECT name FROM user WHERE id=?", id).Scan(&name)

	if err != nil {
		return "", xerrors.Wrap(err, "query failed")
	}
	return name, nil
}
