package model

import (
	"database/sql"
	"fmt"
)

// type Model struct {
// 	ID         int `gorm:"primary_key" json:"id"`
// 	CreatedOn  int `json:"created_on"`
// 	ModifiedOn int `json:"modified_on"`
// 	DeletedOn  int `json:"deleted_on"`
// }

func init1() {

	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/database_name")
	if err != nil {
		fmt.Println("连接数据库失败:", err)
		return
	}
	defer db.Close()

	// 测试连接是否成功
	err = db.Ping()
	if err != nil {
		fmt.Println("数据库连接失败:", err)
		return
	}

	// 执行 SQL 查询
	rows, err := db.Query("SELECT * FROM table_name")
	if err != nil {
		fmt.Println("查询数据失败:", err)
		return
	}
	defer rows.Close()

	// 遍历结果集
	for rows.Next() {
		var col1 int
		var col2 string
		err := rows.Scan(&col1, &col2)
		if err != nil {
			fmt.Println("读取数据失败:", err)
			return
		}
		fmt.Println("col1:", col1)
		fmt.Println("col2:", col2)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("遍历结果集失败:", err)
	}
}
