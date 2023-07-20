/*
 * @Date: 2020-12-15 16:34:06
 * @LastEditors: aiden.deng (Zhenpeng Deng)
 * @LastEditTime: 2020-12-15 16:44:18
 */
package constval

type StringUint16Const interface {
	Key() string
	Value() uint16
}

type Uint16ConstImpl struct {
	key   string
	value uint16
}

func NewStringUint16(key string, value uint16) StringUint16Const {
	return &Uint16ConstImpl{key: key, value: value}
}

func (c *Uint16ConstImpl) Value() uint16 {
	return c.value
}

func (c *Uint16ConstImpl) Key() string {
	return c.key
}
