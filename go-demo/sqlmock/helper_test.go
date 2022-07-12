/*
 * File: /sqlmock/helper_test.go                                               *
 * Project: go-demo                                                            *
 * Created At: Monday, 2022/06/27 , 15:42:12                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/06/27 , 16:09:36                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import (
	"database/sql"
	"os"
	"testing"

	"gorm.io/driver/mysql"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/gorm"
)

var (
	mock sqlmock.Sqlmock
	DB   *gorm.DB
)

// TestMain是在当前package下，最先运行的一个函数，常用于初始化
func TestMain(m *testing.M) {
	var (
		db  *sql.DB
		err error
	)
	//把匹配器设置成相等匹配器，不设置默认使用正则匹配
	db, mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		panic(err)
	}
	DB, err = gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	// m.Run 是调用包下面各个Test函数的入口
	os.Exit(m.Run())
}
