package constval

type Int32Const interface {
	Key() int32
	Value() string
}

type int32ConstImpl struct {
	key   int32
	value string
}

func NewInt32(key int32, value string) Int32Const {
	return &int32ConstImpl{key: key, value: value}
}

func (c *int32ConstImpl) Value() string {
	return c.value
}

func (c *int32ConstImpl) Key() int32 {
	return c.key
}
