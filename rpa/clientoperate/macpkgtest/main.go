package main

import (
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
	"time"
)

// 库有升级，可对照修改
func main() {
	// 查找目标窗口
	title := "目标窗口标题"
	window := cocoa.NSWindow_fromPointer(
		objc.GetClass("NSWindow").Call("windowWithTitle:", core.String(title)).Pointer(),
	)
	if window == nil {
		panic("未找到目标窗口")
	}
	// 将目标窗口置于前台
	window.MakeKeyAndOrderFront(nil)
	// 模拟点击操作
	point := core.CGPoint{X: 300, Y: 300}
	event := cocoa.NSEvent_MouseEventWithType(
		cocoa.NSLeftMouseDown,
		point,
		cocoa.NSLeftMouse,
	)
	event.Post(cocoa.NSEventTypeSystemDefined)
	event = cocoa.NSEvent_MouseEventWithType(
		cocoa.NSLeftMouseUp,
		point,
		cocoa.NSLeftMouse,
	)
	event.Post(cocoa.NSEventTypeSystemDefined)
	// 模拟输入操作
	event = cocoa.NSEvent_KeyEvent(
		cocoa.NSKeyDown,
		core.NSString_FromString("a"),
		true,
	)
	event.Post(cocoa.NSEventTypeSystemDefined)
	event = cocoa.NSEvent_KeyEvent(
		cocoa.NSKeyDown,
		core.NSString_FromString("b"),
		true,
	)
	event.Post(cocoa.NSEventTypeSystemDefined)
	// 模拟快捷键操作
	event = cocoa.NSEvent_KeyEvent(
		cocoa.NSKeyDown,
		core.NSString_FromString("c"),
		true,
	)
	event.SetFlags(cocoa.NSEventModifierFlagCommand)
	event.Post(cocoa.NSEventTypeSystemDefined)
	time.Sleep(time.Millisecond * 100)
	event = cocoa.NSEvent_KeyEvent(
		cocoa.NSKeyDown,
		core.NSString_FromString("\r"),
		true,
	)
	event.SetFlags(cocoa.NSEventModifierFlagCommand)
	event.Post(cocoa.NSEventTypeSystemDefined)
}
