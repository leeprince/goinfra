# Microsoft UI Automation

https://learn.microsoft.com/zh-cn/dotnet/framework/ui-automation/ui-automation-fundamentals
---



# openAI
## microsoft ui automation
golang 如何通过microsoft ui automation 操作Windows 客户端,需要操作的内容包含：
```
通过标题找到控件，并开始输入字符串，然后再回车，最后再按F2
通过AutomationId找到控件，并开始输入字符串，然后再回车，最后再按F2
```

```
package main

import (
    "fmt"
    ".com/go-ole/go-ole"
    "github.com/go-ole/go-ole/oleutil"
    "github.com/moutend/go-wca/pkg/uia"
)

func main() {
    // 初始化COM库
    ole.CoInitialize(0)

    // 获取桌面窗口
    desktop, _ := uia.GetRootElement()

    // 通过标题找到控件
    title := "窗口标题"
    condition := uia.NewAndCondition(
        uia.NewPropertyCondition(uia.UIA_NamePropertyId, title),
        uia.NewPropertyCondition(uia.UIA_ControlTypePropertyId, uia.UIA_EditControlTypeId),
    )
    edit, _ := desktop.FindFirst(uia.TreeScopeDescendants, condition)
    if edit != nil {
        // 输入字符串
        oleutil.PutProperty(edit, "Value", "要输入的字符串")
        // 回车
        oleutil.CallMethod(edit, "Press", 13)
        // 按F2
        oleutil.CallMethod(edit, "Press", 113)
    }

    // 通过AutomationId找到控件
    automationId := "控件的AutomationId"
    condition = uia.NewAndCondition(
        uia.NewPropertyCondition(uia.UIA_AutomationIdPropertyId, automationId),
        uia.NewPropertyCondition(uia.UIA_ControlTypePropertyId, uia.UIA_EditControlTypeId),
    )
    edit, _ = desktop.FindFirst(uia.TreeScopeDescendants, condition)
    if edit != nil {
        // 输入字符串
        oleutil.PutProperty(edit, "Value", "要输入的字符串")
        // 回车
        oleutil.CallMethod(edit, "Press", 13)
        // 按F2
        oleutil.CallMethod(edit, "Press", 113)
    }

    // 释放COM库
    ole.CoUninitialize()
}
```


