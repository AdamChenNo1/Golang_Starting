/*
 * File: /gomock/user/user.go                                                  *
 * Project: go-demo                                                            *
 * Created At: Sunday, 2022/06/26 , 12:30:45                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Sunday, 2022/06/26 , 12:36:13                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package user

import (
	"go_start/go-demo/gomock/person"
)

type User struct {
	Person person.Male
}

func NewUser(p person.Male)*User {
	return &User{p}
}

func (u *User)GetUserInfo(id int64)error {
	return u.Person.Get(id)
}
