package constdef

import "fmt"

/*
 * @Date: 2020-11-03 10:51:19
 * @LastEditors: aiden.deng (Zhenpeng Deng)
 * @LastEditTime: 2020-12-15 16:46:51
 */

type StringGroup struct {
	value2Const map[string]StringConst
}

func NewStringGroup(consts ...StringConst) *StringGroup {
	group := &StringGroup{
		value2Const: make(map[string]StringConst),
	}
	for _, c := range consts {
		group.value2Const[c.Value()] = c
	}
	return group
}

func (g *StringGroup) Get(v string) (StringConst, bool) {
	c, ok := g.value2Const[v]
	if !ok {
		return nil, false
	}
	return c, true
}

func (g *StringGroup) MustGet(v string) StringConst {
	c, ok := g.value2Const[v]
	if !ok {
		panic(fmt.Errorf("const %v does not exist", v))
	}
	return c
}

func (g *StringGroup) IsValid(v string) bool {
	_, ok := g.value2Const[v]
	if !ok {
		return false
	}
	return true
}

func (g *StringGroup) Consts() []StringConst {
	consts := make([]StringConst, 0)
	for _, c := range g.value2Const {
		consts = append(consts, c)
	}
	return consts
}
