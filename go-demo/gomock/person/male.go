/*
 * File: /gomock/person/male.go                                                *
 * Project: go-demo                                                            *
 * Created At: Sunday, 2022/06/26 , 12:29:59                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Sunday, 2022/06/26 , 12:30:27                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package person

type Male interface {
	Get(id int64) error
}
