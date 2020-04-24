package gohandler

import "strconv"

type StrconvHandler struct {
	errp *error
}

func (o *StrconvHandler) Error() bool {
	return (*o.errp) != nil
}

func (o *StrconvHandler) Init(errp *error) {
	o.errp = errp
}

func (o *StrconvHandler) Set(err error) {
	*o.errp = err
}

func (o *StrconvHandler) ParseInt(s string, base int) int64 {
	if o.Error() {
		return 0
	}

	i64, err := strconv.ParseInt(s, base, 64)
	o.Set(err)
	
	return i64
}

func (o *StrconvHandler) ParseUint(s string, base int) uint64 {
	if o.Error() {
		return 0
	}

	u64, err := strconv.ParseUint(s, base, 64)
	o.Set(err)
	
	return u64
}

