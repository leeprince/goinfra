/*
 * @Date: 2020-12-15 16:34:06
 * @LastEditors: aiden.deng (Zhenpeng Deng)
 * @LastEditTime: 2020-12-15 16:40:27
 */
package constval2

type Int32Const interface {
	Name() string
	Value() int32
}

type int32ConstImpl struct {
	value int32
	name  string
}

func NewInt32(value int32, name string) Int32Const {
	return &int32ConstImpl{value: value, name: name}
}

func (c *int32ConstImpl) Name() string {
	return c.name
}

func (c *int32ConstImpl) Value() int32 {
	return c.value
}
