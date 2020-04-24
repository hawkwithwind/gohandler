package gohandler

import "fmt"

type ValueHandler struct {
	errp *error
}

func (o *ValueHandler) Error() bool {
	return (*o.errp) != nil
}

func (o *ValueHandler) Init(errp *error) {
	o.errp = errp
}

func (o *ValueHandler) Set(err error) {
	*o.errp = err
}

func (o *ValueHandler) ListValue(value interface{}) []interface{} {
	if o.Error() {
		return nil
	}

	var err error

	switch value := value.(type) {
	case []interface{}:
		return value
	case nil:
		err = fmt.Errorf("expect listvalue, but nil")
	default:
		err = fmt.Errorf("expect listvalue but %T", value)
	}

	o.Set(err)
	return nil
}

func (o *ValueHandler) NilAs(value interface{}, defaultValue interface{}) interface{} {
	// this function won't produce error, for semantic consistency, this function didn't short circuit o.Err != nil
	
	if value == nil {
		return defaultValue
	} else {
		return value
	}
}



