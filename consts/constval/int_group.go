package constdef

import "fmt"

/*
 * @Date: 2020-11-03 10:51:19
 * @LastEditors: aiden.deng (Zhenpeng Deng)
 * @LastEditTime: 2020-12-15 16:45:03
 */

type IntGroup struct {
	value2Const map[int]IntConst
}

func NewIntGroup(consts ...IntConst) *IntGroup {
	group := &IntGroup{
		value2Const: make(map[int]IntConst),
	}
	for _, c := range consts {
		group.value2Const[c.Value()] = c
	}
	return group
}

func (g *IntGroup) Get(v int) (IntConst, bool) {
	c, ok := g.value2Const[v]
	if !ok {
		return nil, false
	}
	return c, true
}

func (g *IntGroup) MustGet(v int) IntConst {
	c, ok := g.value2Const[v]
	if !ok {
		panic(fmt.Errorf("const %v does not exist", v))
	}
	return c
}

func (g *IntGroup) IsValid(v int) bool {
	_, ok := g.value2Const[v]
	if !ok {
		return false
	}
	return true
}

func (g *IntGroup) Consts() []IntConst {
	consts := make([]IntConst, 0)
	for _, c := range g.value2Const {
		consts = append(consts, c)
	}
	return consts
}
