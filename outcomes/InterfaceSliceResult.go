package outcomes

/*
 * @Author: 夏华楼
 * @Date: 2020/12/31 10:12
 * @Email: Variou.xia@aishu.cn
 */
import (
	"github.com/XIAHUALOU/andis/iterator"
)

type SliceResult struct {
	Result []interface{}
	Err    error
}

func NewSliceResult(result []interface{}, err error) *SliceResult {
	return &SliceResult{Result: result, Err: err}
}

func (self *SliceResult) UnWrap() []interface{} {
	if self.Err != nil {
		panic(self.Err)
	}
	return self.Result
}

func (self *SliceResult) UnWrap_Or(res []interface{}) []interface{} {
	if self.Err != nil {
		return res
	}
	return self.Result
}

func (self *SliceResult) Unwrap_Or_Else(f func() []interface{}) []interface{} {
	if self.Err != nil {
		return f()
	}
	return self.Result
}

func (self *SliceResult) Iter() *iterator.Iterator {
	return iterator.NewIterator(self.Result)
}

func (self *SliceResult) Index(i int) interface{} {
	if i > len(self.Result) {
		return nil
	} else {
		return self.Result[i]
	}
}

func (self *SliceResult) Exist(item interface{}) bool {
	for _, val := range self.Result {
		if val == item {
			return true
		}
	}
	return false
}
