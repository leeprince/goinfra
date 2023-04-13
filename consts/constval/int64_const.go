/*
 * @Date: 2020-12-15 16:34:06
 * @LastEditors: aiden.deng (Zhenpeng Deng)
 * @LastEditTime: 2020-12-15 16:44:27
 */
package constdef

type Int64Const interface {
	Name() string
	Value() int64
}

type int64ConstImpl struct {
	value int64
	name  string
}

func NewInt64(value int64, name string) Int64Const {
	return &int64ConstImpl{value: value, name: name}
}

func (c *int64ConstImpl) Name() string {
	return c.name
}

func (c *int64ConstImpl) Value() int64 {
	return c.value
}
