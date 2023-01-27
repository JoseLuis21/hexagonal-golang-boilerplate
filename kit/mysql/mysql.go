package kit

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)



func MysqlCreateConnection(user, pass, host, port, dbName string) (*sql.DB, error){
	fmt.Println(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true", user, pass, host, port, dbName))
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true", user, pass, host, port, dbName))
	if err != nil {
		return nil, err
	}

	return db, nil
}
