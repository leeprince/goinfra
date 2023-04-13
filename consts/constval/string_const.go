/*
 * @Date: 2020-12-15 16:34:06
 * @LastEditors: aiden.deng (Zhenpeng Deng)
 * @LastEditTime: 2020-12-15 16:41:35
 */
package constdef

type StringConst interface {
	Name() string
	Value() string
}

type stringConstImpl struct {
	value string
	name  string
}

func NewString(value string, name string) StringConst {
	return &stringConstImpl{value: value, name: name}
}

func (c *stringConstImpl) Name() string {
	return c.name
}

func (c *stringConstImpl) Value() string {
	return c.value
}
