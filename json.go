package gohandler

import "encoding/json"

type JsonHandler struct {
	errp *error
}

func (o *JsonHandler) Error() bool {
	return (*o.errp) != nil
}

func (o *JsonHandler) Init(errp *error) {
	o.errp = errp
}

func (o *JsonHandler) Set(err error) {
	*o.errp = err
}

func (o *JsonHandler) ToJson(v interface{}) string {
	if o.Error() {
		return ""
	}

	jsonstr, err := json.Marshal(v)
	o.Set(err)
	
	if o.Error() {
		return ""
	}

	return string(jsonstr)
}


func (o *JsonHandler) FromJson(jsonstr string) map[string]interface{} {
	if o.Error() {
		return nil
	}

	ret := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonstr), &ret)
	o.Set(err)
	
	return ret
}
