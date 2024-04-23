package errs

type ErrContext struct {
	HttpStatus int    `json:"httpStatus"`
	ErrMsg     string `json:"errMsg"`
	ErrDesc    string `json:"errDescription"`
}

func NewErrContextWithError(errCtx ErrContext, err error) ErrContext {
	return ErrContext{
		HttpStatus: errCtx.HttpStatus,
		ErrMsg:     errCtx.ErrMsg,
		ErrDesc:    err.Error(),
	}
}

func NewErrContext(errCtx ErrContext) ErrContext {
	return ErrContext{
		HttpStatus: errCtx.HttpStatus,
		ErrMsg:     errCtx.ErrMsg,
	}
}

func (r ErrContext) Error() string {
	return r.ErrMsg
}

func (r *ErrContext) GetHttpStatus() int {
	return r.HttpStatus
}

func (r *ErrContext) GetErrDesc() string {
	return r.ErrDesc
}
