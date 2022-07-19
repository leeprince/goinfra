# 测试工具
---


## 概述

Go 的测试支持在包内优先执行一个 TestMain(m *testing.M) 函数，可以在这里为 package 下所有测试做一些初始化的工作。
```
func TestMain(m *testing.M) {

    ...

    // m.Run 是调用包下面各个Test函数的入口
    os.Exit(m.Run())
}
```