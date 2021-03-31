package outcomes

/*
 * @Author: 夏华楼
 * @Date: 2020/12/31 10:12
 * @Email: Variou.xia@aishu.cn
 */
type StringResult struct {
	Result string
	Err    error
}

func NewStringResult(result string, err error) *StringResult {
	return &StringResult{Result: result, Err: err}
}

func (self *StringResult) Unwrap() string {
	if self.Err != nil {
		panic(self.Err)
	}
	return self.Result
}
func (self *StringResult) Unwrap_Or(str string) string {
	if self.Err != nil {
		return str
	}
	return self.Result
}

func (self *StringResult) Unwrap_Or_Else(f func() string) string {
	if self.Err != nil {
		return f()
	}
	return self.Result
}

func (self *StringResult) Unwrap_Exist() bool { //只判断key存在与否，不需要值时使用
	if self.Err != nil {
		return false
	}
	return true
}
