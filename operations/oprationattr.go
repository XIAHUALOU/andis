package operations

/*
 * @Author: 夏华楼
 * @Date: 2020/12/31 10:12
 * @Email: Variou.xia@aishu.cn
 */
import (
	"fmt"
	"github.com/XIAHUALOU/andis/outcomes"
	"time"
)

const (
	ATTR_EXPR = "expr"
	ATTR_NX   = "nx"
	ATTR_XX   = "xx"
)

type empty struct{}

type OperationAttr struct {
	Name  string
	Value interface{}
}

func (self OperationAttrs) Find(name string) *outcomes.InterfaceResult {
	for _, attr := range self {
		if attr.Name == name {
			return outcomes.NewInterfaceResult(attr.Value, nil)
		}
	}
	return outcomes.NewInterfaceResult(nil, fmt.Errorf("OperationAttrs found error %s", name))
}

type OperationAttrs []*OperationAttr

func WithExpire(t time.Duration) *OperationAttr { //过期时间
	return &OperationAttr{Name: ATTR_EXPR, Value: t}
}

func WithNX() *OperationAttr { //仅在键不存在时设置键
	return &OperationAttr{
		Name:  ATTR_NX,
		Value: empty{},
	}
}
func WithXX() *OperationAttr { //只有在键已存在时才设置
	return &OperationAttr{Name: ATTR_XX, Value: empty{}}
}
