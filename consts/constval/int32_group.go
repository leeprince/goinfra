package constval

import "fmt"

type Int32Group struct {
	value2Const map[int32]Int32Const
}

func NewInt32Group(consts ...Int32Const) *Int32Group {
	group := &Int32Group{
		value2Const: make(map[int32]Int32Const),
	}
	for _, c := range consts {
		group.value2Const[c.Key()] = c
	}
	return group
}

func (g *Int32Group) Get(v int32) (Int32Const, bool) {
	c, ok := g.value2Const[v]
	if !ok {
		return nil, false
	}
	return c, true
}

func (g *Int32Group) MustGet(v int32) Int32Const {
	c, ok := g.value2Const[v]
	if !ok {
		panic(fmt.Errorf("const %v does not exist", v))
	}
	return c
}

func (g *Int32Group) IsValid(v int32) bool {
	_, ok := g.value2Const[v]
	if !ok {
		return false
	}
	return true
}

func (g *Int32Group) Consts() []Int32Const {
	consts := make([]Int32Const, 0)
	for _, c := range g.value2Const {
		consts = append(consts, c)
	}
	return consts
}
