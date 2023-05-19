package constval

import "fmt"

/*
 * @Date: 2020-11-03 10:51:19
 * @LastEditors: aiden.deng (Zhenpeng Deng)
 * @LastEditTime: 2020-12-15 16:46:30
 */

type Int64Group struct {
	value2Const map[int64]Int64Const
}

func NewInt64Group(consts ...Int64Const) *Int64Group {
	group := &Int64Group{
		value2Const: make(map[int64]Int64Const),
	}
	for _, c := range consts {
		group.value2Const[c.Key()] = c
	}
	return group
}

func (g *Int64Group) Get(v int64) (Int64Const, bool) {
	c, ok := g.value2Const[v]
	if !ok {
		return nil, false
	}
	return c, true
}

func (g *Int64Group) MustGet(v int64) Int64Const {
	c, ok := g.value2Const[v]
	if !ok {
		panic(fmt.Errorf("const %v does not exist", v))
	}
	return c
}

func (g *Int64Group) IsValid(v int64) bool {
	_, ok := g.value2Const[v]
	if !ok {
		return false
	}
	return true
}

func (g *Int64Group) Consts() []Int64Const {
	consts := make([]Int64Const, 0)
	for _, c := range g.value2Const {
		consts = append(consts, c)
	}
	return consts
}
