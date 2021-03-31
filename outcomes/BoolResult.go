package outcomes

/*
 * @Author: 夏华楼
 * @Date: 2020/12/31 10:12
 * @Email: Variou.xia@aishu.cn
 */
type BooleanResult struct {
	Result bool
	Err    error
}

func NewBooleanResult(result bool, err error) *BooleanResult {
	return &BooleanResult{Result: result, Err: err}
}
func (self *BooleanResult) Unwrap() bool {
	if self.Err != nil {
		panic(self.Err)
	}
	return self.Result
}
func (self *BooleanResult) Unwrap_Or(v bool) bool {
	if self.Err != nil {
		return v
	}
	return self.Result
}

func (self *BooleanResult) Unwrap_Or_Else(f func() bool) bool {
	if self.Err != nil {
		return f()
	}
	return self.Result
}
