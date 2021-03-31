package outcomes

/*
 * @Author: 夏华楼
 * @Date: 2020/12/31 10:12
 * @Email: Variou.xia@aishu.cn
 */
type IntResult struct {
	Result int64
	Err    error
}

func NewIntResult(result int64, err error) *IntResult {
	return &IntResult{Result: result, Err: err}
}
func (self *IntResult) Unwrap() int64 {
	if self.Err != nil {
		panic(self.Err)
	}
	return self.Result
}
func (self *IntResult) Unwrap_Or(v int64) int64 {
	if self.Err != nil {
		return v
	}
	return self.Result
}

func (self *IntResult) Unwrap_Or_Else(f func() int64) int64 {
	if self.Err != nil {
		return f()
	}
	return self.Result
}
