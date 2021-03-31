package outcomes

/*
 * @Author: 夏华楼
 * @Date: 2021/3/25 9:48
 * @Email: Variou.xia@aishu.cn
 */
import (
	"github.com/XIAHUALOU/andis/iterator"
)

type StringSliceResult struct {
	Result []string
	Err    error
}

func NewStringSliceResult(result []string, err error) *StringSliceResult {
	return &StringSliceResult{Result: result, Err: err}
}

func (self *StringSliceResult) UnWrap() []string {
	if self.Err != nil {
		panic(self.Err)
	}
	return self.Result
}

func (self *StringSliceResult) UnWrap_Or(res []string) []string {
	if self.Err != nil {
		return res
	}
	return self.Result
}

func (self *StringSliceResult) Unwrap_Or_Else(f func() []string) []string {
	if self.Err != nil {
		return f()
	}
	return self.Result
}

func (self *StringSliceResult) Iter() *iterator.Iterator {
	return iterator.NewIterator(self.Result)
}

func (self *StringSliceResult) Index(i int) interface{} {
	if i > len(self.Result) {
		return nil
	} else {
		return self.Result[i]
	}
}

func (self *StringSliceResult) Exist(item interface{}) bool {
	for _, val := range self.Result {
		if val == item {
			return true
		}
	}
	return false
}
