/*
 * File: \main.go                                                              *
 * Project: go-mysql                                                           *
 * Created At: Sunday, 2022/05/29 , 20:52:16                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Sunday, 2022/05/29 , 20:54:27                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:toor@tcp(db-elchn.cloud.okteto.net:3306)/")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}
