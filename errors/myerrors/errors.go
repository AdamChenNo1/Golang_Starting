/*
 * File: /myerrors/errors.go                                                   *
 * Project: errors                                                             *
 * Created At: Monday, 2022/06/20 , 12:36:16                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/06/20 , 13:53:23                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package myerrors

import (
	"github.com/marmotedu/errors"
	code "github.com/marmotedu/sample-code"
)

type MyError struct {
	// Code defines the business error code.
	Code int `json:"code"`

	// Message contains the detail of this message.
	// This message is suitable to be exposed to external
	Message string `json:"message"`

	Details any `json:"details,omitempty"`
}

func ParseError(e error) MyError {
	myerror := MyError{
	}
	if e, ok := err.(*withCode); ok {
		myerror.Code = 
	}

}
