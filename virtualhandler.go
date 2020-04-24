package gohandler

type VirtualHandler interface{
	Error() bool
	Init(err *error)
	Set(err error)
}


