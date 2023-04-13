# 状态码管理
---

# BizErr

# 通用状态码及描述
状态码格式： 业务线(2位)+微服务(2位)+功能模块(2位)+业务逻辑标识(3位)
    - 特殊的状态码：正确响应的状态码等于0(非0都表示为处理失败)、未授权、禁止访问
    - 该通用状态码的格式为：业务线(业务线通用10)+微服务(通用服务:00)+功能模块(通用功能：00)+业务逻辑标识(xxx)

# 关于 `"github.com/pkg/errors"` 与 `fmt.Errorf` 对包裹错误(wrap)的处理
> 详见：bizerr_def_test.go 中 `TestFmtErrorfWrap()` 与 `TestErrorWrap()` 的测试结果

结论：推荐使用 `fmt.Errorf`。原因：`errors` 包错误的解包裹(Unwrap)支持性更好！