/*
 * File: /gomock/user/user_test.go                                             *
 * Project: go-demo                                                            *
 * Created At: Sunday, 2022/06/26 , 12:36:33                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/06/27 , 09:12:11                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package user

import (
	"testing"

	"go_start/go-demo/gomock/mock"

	"github.com/golang/mock/gomock"
)

func TestUserGetInfo(t *testing.T) {
	ctl := gomock.NewController(t) // 它代表 mock 生态系统中的顶级控件。定义了 mock 对象的范围、生命周期和期待值。另外它在多个 goroutine 中是安全的

	var id int64 = 1

	mockMale := mock.NewMockMale(ctl) // 创建一个新的 mock 实例
	gomock.InOrder(                   // 声明给定的调用应按顺序进行（是对 gomock.After 的二次封装）
		mockMale.EXPECT().Get(id).Return(nil),
	)

	user := NewUser(mockMale) // 创建 User 实例，注入 mock 对象

	err := user.GetUserInfo(id)
	if err != nil {
		t.Error("user.GetUserInfo error: ", err)
	}

}
