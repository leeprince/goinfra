/*
 * @Date: 2020-12-15 16:34:06
 * @LastEditors: aiden.deng (Zhenpeng Deng)
 * @LastEditTime: 2020-12-15 16:44:18
 */
package constval

type IntConst interface {
	Key() int
	Value() string
}

type intConstImpl struct {
	key   int
	value string
}

func NewInt(key int, value string) IntConst {
	return &intConstImpl{key: key, value: value}
}

func (c *intConstImpl) Value() string {
	return c.value
}

func (c *intConstImpl) Key() int {
	return c.key
}
