package test

import (
	"fmt"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/9/26 10:45
 * @Desc:
 */

// 定义一个接口类型
type itf interface {
	str(s string)
}

// 定义一个结构体类型
type imp1 struct {
}

// 实现接口方法
func (receiver imp1) str(s string) {
	fmt.Println("imp1 s", s)
}

func retITF(i itf) {
	i.str("retITF")
}

func TestInterface(t *testing.T) {
	i1 := &imp1{}
	i1.str("TestInterface 1")

	retITF(i1)
}

// 定义一个接口类型
type Shape interface {
	Area() float64
}

// 定义一个结构体类型
type Rectangle struct {
	Width  float64
	Height float64
}

// 实现接口方法
func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func TestShapeInterface(t *testing.T) {
	// 创建一个指针类型的结构体对象
	rect := &Rectangle{
		Width:  10,
		Height: 5,
	}
	fmt.Println("rect Area:", rect.Area())

	// 将指针类型的结构体对象传递给接口类型
	/*
		var shape Shape 是否可以改成  var shape *Shape?
			不可以将 var shape Shape 改成 var shape *Shape。在 Golang 中，接口类型是一个抽象类型，它本身不是指针类型。因此，我们不能声明一个指向接口类型的指针。
			接口类型的变量可以持有任何实现了该接口的类型的值，包括指针类型和非指针类型。因此，我们可以直接声明一个接口类型的变量，而不需要使用指针
	*/
	var shape Shape = rect

	// 调用接口方法
	fmt.Println("shape Area:", shape.Area())
}
