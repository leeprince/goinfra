package main

import (
	"fmt"
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/5/24 18:18
 * @Desc:
 */

func main() {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	// 指定目标应用程序的进程ID
	targetProcessID := int32(1234)

	// 获取Shell对象
	shell, err := oleutil.CreateObject("Shell.Application")
	if err != nil {
		fmt.Println("创建Shell对象失败:", err)
		return
	}
	defer shell.Release()

	// 获取Windows对象
	windows, err := oleutil.CallMethod(shell, "Windows")
	if err != nil {
		fmt.Println("获取Windows对象失败:", err)
		return
	}

	// 遍历窗口
	enum, err := windows.CallMethod("_NewEnum")
	if err != nil {
		fmt.Println("获取窗口枚举失败:", err)
		return
	}

	enumerator := enum.ToIUnknown()
	defer enumerator.Release()

	for {
		item, err := oleutil.CallMethod(enumerator, "Next")
		if err != nil {
			fmt.Println("获取下一个窗口失败:", err)
			break
		}

		if item.Val == 0 {
			break
		}

		window := item.ToIDispatch()
		processID, err := oleutil.GetProperty(window, "ProcessID")
		window.Release()

		if err != nil {
			fmt.Println("获取窗口进程ID失败:", err)
			break
		}

		if processID.Value().(int32) != targetProcessID {
			continue
		}

		// 遍历窗口内的所有元素
		//traverseElements(window.ToIDispatch())
	}
}

func traverseElements(element *ole.IDispatch) {
	// 获取元素的名称
	name, _ := oleutil.GetProperty(element, "Name")
	fmt.Println("元素名称:", name.ToString())

	// 获取元素的控件类型
	controlType, _ := oleutil.GetProperty(element, "ControlType")

	// 判断控件类型并执行相应操作
	if controlType.Val == 50000 { // 按钮
		fmt.Println("元素类型: 按钮")

		// 模拟点击按钮
		invokePattern, err := oleutil.CallMethod(element, "GetCurrentPattern", -4158) // -4158 对应 InvokePattern
		if err == nil {
			invokePattern.ToIDispatch().CallMethod("Invoke")
		}
	} else if controlType.Val == 50004 { // 文本框
		fmt.Println("元素类型: 文本框")

		// 获取文本框的内容
		text, _ := oleutil.GetProperty(element, "Value")
		fmt.Println("文本框内容:", text.ToString())
	}

	// 获取当前元素的子元素
	children, _ := oleutil.CallMethod(element, "FindAll", 4) // 4 对应 TreeScope_Children

	enumerator := children.ToIDispatch()
	defer enumerator.Release()

	// 递归遍历子元素
	for {
		item, _ := oleutil.CallMethod(enumerator, "Next")
		if item.Val == 0 {
			break
		}

		childElement := item.ToIDispatch

		fmt.Println(childElement())
	}
}
