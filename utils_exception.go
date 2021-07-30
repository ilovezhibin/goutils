package main

import "strconv"

type UtilsException struct {
	Code int
	Msg string
}

/*
	创建一个错误
 */
func ThrowException(msg string) *UtilsException {
	return &UtilsException{-1,msg}
}

func ThrowExceptionByErr(err error) *UtilsException {
	return &UtilsException{-1,err.Error()}
}

func (e *UtilsException) Error() string {
	return "code:"+strconv.Itoa(e.Code)+",msg:"+e.Msg
}
