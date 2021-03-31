package outcomes

/*
 * @Author: 夏华楼
 * @Date: 2020/12/31 10:12
 * @Email: Variou.xia@aishu.cn
 */
type InterfaceResult struct {
	Result interface{}
	Err    error
}

func NewInterfaceResult(result interface{}, err error) *InterfaceResult {
	return &InterfaceResult{Result: result, Err: err}
}
func (self *InterfaceResult) Unwrap() interface{} {
	if self.Err != nil {
		panic(self.Err)
	}
	return self.Result
}
func (self *InterfaceResult) Unwrap_Or(v interface{}) interface{} {
	if self.Err != nil {
		return v
	}
	return self.Result
}

func (self *InterfaceResult) Unwrap_Or_Else(f func() interface{}) interface{} {
	if self.Err != nil {
		return f()
	}
	return self.Result
}
