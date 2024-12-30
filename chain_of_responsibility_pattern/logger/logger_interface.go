package logger

import "solid_design/chain_of_responsibility/enum"

type Logger interface {
	Log(logLevel enum.Logger, msg string)
	SetNext(l Logger)
}


func GetLogger()  Logger {
	err := &Error{}
	debug := &Debug{}
	info := &Info{}
	info.SetNext(debug)
	debug.SetNext(err)
	return info
}