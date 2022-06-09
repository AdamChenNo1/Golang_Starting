/*
 * File: \rpc\hellorpc\hellorpc.go                                             *
 * Project: test                                                               *
 * Created At: Wednesday, 2022/06/8 , 22:48:15                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Wednesday, 2022/06/8 , 22:49:10                              *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package hellorpc

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {

	*reply = "hello:" + request

	return nil

}
