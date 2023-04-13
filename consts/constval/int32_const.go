package constdef

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
