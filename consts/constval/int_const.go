/*
 * @Date: 2020-12-15 16:34:06
 * @LastEditors: aiden.deng (Zhenpeng Deng)
 * @LastEditTime: 2020-12-15 16:44:18
 */
package constdef

type IntConst interface {
	Name() string
	Value() int
}

type intConstImpl struct {
	value int
	name  string
}

func NewInt(value int, name string) IntConst {
	return &intConstImpl{value: value, name: name}
}

func (c *intConstImpl) Name() string {
	return c.name
}

func (c *intConstImpl) Value() int {
	return c.value
}
