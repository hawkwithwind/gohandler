package gohandler

type GoHandler struct {
	JsonHandler
	ValueHandler
	StrconvHandler
}

func (o *GoHandler) Error() bool {
	return o.JsonHandler.Error()
}

func (o *GoHandler) Init(errp *error) {
	o.JsonHandler.Init(errp)
	o.ValueHandler.Init(errp)
	o.StrconvHandler.Init(errp)
}

func (o *GoHandler) Set(err error) {
	o.JsonHandler.Set(err)
	o.ValueHandler.Set(err)
	o.StrconvHandler.Set(err)
}


