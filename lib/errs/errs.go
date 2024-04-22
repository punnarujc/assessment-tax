package errs

type ErrContext struct {
	HttpStatus int    `json:"httpStatus"`
	ErrMsg     string `json:"errMsg"`
	ErrDesc    string `json:"errDescription"`
}

func NewErrContext(errCtx ErrContext, err error) ErrContext {
	return ErrContext{
		HttpStatus: errCtx.HttpStatus,
		ErrMsg:     errCtx.ErrMsg,
		ErrDesc:    err.Error(),
	}
}

func (r *ErrContext) Error() string {
	return r.ErrMsg
}
