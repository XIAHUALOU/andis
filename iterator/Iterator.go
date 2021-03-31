package iterator

/*
 * @Author: 夏华楼
 * @Date: 2020/12/31 10:12
 * @Email: Variou.xia@aishu.cn
 */
type Iterator struct {
	data  interface{}
	index int
}

func NewIterator(data interface{}) *Iterator {
	return &Iterator{data: data}
}

func (self *Iterator) HasNext() bool {
	if self.data == nil {
		return false
	} else {
		switch self.data.(type) {
		case []string:
			if len(self.data.([]string)) > 0 && self.index < len(self.data.([]string)) {
				return true
			} else {
				return false
			}
		case []interface{}:
			if len(self.data.([]interface{})) > 0 && self.index < len(self.data.([]interface{})) {
				return true
			} else {
				return false
			}
		default:
			return false
		}
	}
}

func (self *Iterator) Next() (ret interface{}) {
	switch self.data.(type) {
	case []string:
		ret = self.data.([]string)[self.index]
		self.index = self.index + 1
	case []interface{}:
		ret = self.data.([]interface{})[self.index]
		self.index = self.index + 1
	}
	return ret
}
