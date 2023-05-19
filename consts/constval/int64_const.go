package constval

type Int64Const interface {
	Key() int64
	Value() string
}

type int64ConstImpl struct {
	key   int64
	value string
}

func NewInt64(key int64, value string) Int64Const {
	return &int64ConstImpl{key: key, value: value}
}

func (c *int64ConstImpl) Value() string {
	return c.value
}

func (c *int64ConstImpl) Key() int64 {
	return c.key
}
