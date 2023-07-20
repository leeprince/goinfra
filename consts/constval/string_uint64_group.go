package constval

import "fmt"

/*
 * @Date: 2020-11-03 10:51:19
 * @LastEditors: aiden.deng (Zhenpeng Deng)
 * @LastEditTime: 2020-12-15 16:46:30
 */

type StringUint16Group struct {
	value2Const map[string]StringUint16Const
}

func NewStringUint16Group(consts ...StringUint16Const) *StringUint16Group {
	group := &StringUint16Group{
		value2Const: make(map[string]StringUint16Const),
	}
	for _, c := range consts {
		group.value2Const[c.Key()] = c
	}
	return group
}

func (g *StringUint16Group) Get(v string) (StringUint16Const, bool) {
	c, ok := g.value2Const[v]
	if !ok {
		return nil, false
	}
	return c, true
}

func (g *StringUint16Group) MustGet(v string) StringUint16Const {
	c, ok := g.value2Const[v]
	if !ok {
		panic(fmt.Errorf("const %v does not exist", v))
	}
	return c
}

func (g *StringUint16Group) IsValid(v string) bool {
	_, ok := g.value2Const[v]
	if !ok {
		return false
	}
	return true
}

func (g *StringUint16Group) Consts() []StringUint16Const {
	consts := make([]StringUint16Const, 0)
	for _, c := range g.value2Const {
		consts = append(consts, c)
	}
	return consts
}
