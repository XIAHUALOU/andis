package outcomes

/*
 * @Author: 夏华楼
 * @Date: 2020/12/31 10:12
 * @Email: Variou.xia@aishu.cn
 */
import "time"

type TimeResult struct {
	Result time.Duration
	Err    error
}

func NewTimeResult(result time.Duration, err error) *TimeResult {
	return &TimeResult{Result: result, Err: err}
}
func (self *TimeResult) Unwrap() time.Duration {
	if self.Err != nil {
		panic(self.Err)
	}
	return self.Result
}
func (self *TimeResult) Unwrap_Or(v time.Duration) time.Duration {
	if self.Err != nil {
		return v
	}
	return self.Result
}

func (self *TimeResult) Unwrap_Or_Else(f func() time.Duration) time.Duration {
	if self.Err != nil {
		return f()
	}
	return self.Result
}
