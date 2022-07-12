/*
 * File: /sqlmock/blog.go                                                      *
 * Project: go-demo                                                            *
 * Created At: Monday, 2022/06/27 , 15:47:40                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/06/27 , 15:50:38                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string
}

type Author struct {
	Name  string
	Email string
}

type Blog struct {
	ID      int
	Author  Author `gorm:"embedded"`
	Upvotes int32
}
