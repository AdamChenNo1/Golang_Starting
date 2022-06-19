/*
 * File: \codec\codec.go                                                       *
 * Project: simple-rpc                                                         *
 * Created At: Tuesday, 2022/06/14 , 13:11:23                                  *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Tuesday, 2022/06/14 , 13:30:49                               *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package codec

import "io"

type Header struct {
	ServiceMethod string
	Seq           uint64 //sequence number,chosen by client to identify a request
	Error         string
}

type Codec interface {
	io.Closer
	ReadHeader(*Header) error
	ReadBody(any) error
	Write(*Header, any) error
}

type Type string

const (
	GobType  Type = "application/gob"
	JSONType Type = "application/json"
)

type NewCodecFunc func(io.ReadWriteCloser) Codec

var NewCodecFuncMap map[Type]NewCodecFunc

func init() {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
}
