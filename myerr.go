package goutils

/**
eg:
func MyList() *goutils.MyError {
	myErr := goutils.MyErr()

	//return myErr.SetErr("Fail")
	//return myErr.SetErr("Fail").SetCode(1)
	//return myErr.SetErr("Fail").SetCode(1).SetData(nil)
}
*/

// errCode 0-成功 非0-失败
func MyErr(params ...interface{}) *MyError {
	err := &MyError{
		ErrCode: 0,
		ErrMsg:"",
		ExtData: "{}",
	}
	switch len(params) {
	case 1:
		err.ErrCode = params[0].(int)
	case 2:
		err.ErrCode = params[0].(int)
		err.ErrMsg = params[1].(string)
	case 3:
		err.ErrCode = params[0].(int)
		err.ErrMsg = params[1].(string)
		err.ExtData = params[2]
	default:
	}
	return err
}

type MyError struct {
	ErrCode int         //错误代号
	ErrMsg  string      //错误信息
	ExtData interface{} // 错误扩展数据
}

func (e *MyError) SetMsg(msg string) *MyError {
	e.ErrMsg = msg
	return e
}

func (e *MyError) SetCode(code int) *MyError {
	e.ErrCode = code
	return e
}

func (e *MyError) SetData(extData interface{}) *MyError {
	e.ExtData = extData
	return e
}
