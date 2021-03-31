package outcomes

/*
 * @Author: 夏华楼
 * @Date: 2021/3/25 11:32
 * @Email: Variou.xia@aishu.cn
 */
type FloatResult struct {
	Result float64
	Err    error
}

func NewFloatResult(result float64, err error) *FloatResult {
	return &FloatResult{Result: result, Err: err}
}
func (self *FloatResult) Unwrap() float64 {
	if self.Err != nil {
		panic(self.Err)
	}
	return self.Result
}
func (self *FloatResult) Unwrap_Or(v float64) float64 {
	if self.Err != nil {
		return v
	}
	return self.Result
}

func (self *FloatResult) Unwrap_Or_Else(f func() float64) float64 {
	if self.Err != nil {
		return f()
	}
	return self.Result
}
