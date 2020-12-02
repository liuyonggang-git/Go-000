package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// dao层应该返回实际结果和错误
func GetUsernameById(id string) (*string, error) {
	var username string
	db, err := sql.Open("mysql", "root:12345678@tcp(localhost:3306)/go-000")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	err = db.QueryRow("select username from user where user_id = ?", id).Scan(&username)
	return &username, err
}
