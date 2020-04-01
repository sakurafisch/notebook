# postgresql

```go
package main

import (
	"database/sql"
	"fmt"

	"github.com/astaxie/beego/logs"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	initDb()
}

func initDb() error {
	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "test", "test", "postgres")

	var err error
	db, err = sql.Open("postgres", dbInfo)

	if err != nil {
		logs.Error("打开数据库失败", err)
		return err
	}

	if err := db.Ping(); err != nil {
		logs.Error("连接数据库失败", err)
		return err
	}
	fmt.Println("ok")
	return nil
}
```
