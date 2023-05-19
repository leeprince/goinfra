/*
 * @Date: 2020-12-15 16:34:06
 * @LastEditors: aiden.deng (Zhenpeng Deng)
 * @LastEditTime: 2020-12-15 16:41:35
 */
package constval

type StringConst interface {
	Key() string
	Value() string
}

type stringConstImpl struct {
	key   string
	value string
}

func NewString(key string, value string) StringConst {
	return &stringConstImpl{key: key, value: value}
}

func (c *stringConstImpl) Value() string {
	return c.value
}

func (c *stringConstImpl) Key() string {
	return c.key
}
