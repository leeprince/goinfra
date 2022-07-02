package security

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/6/28 下午10:30
 * @Desc:
 */

type Option struct {
    IsToHex bool // 是否转为十六进制
}

type OptionFunc func(opt *Option)

// 初始化可选项
func initOption(opts ...OptionFunc) *Option {
    opt := &Option{
            IsToHex: false,
        }
    for _, optFunc := range opts {
        optFunc(opt)
    }
    return opt
}

func WithIsToHex(v bool) OptionFunc {
    return func(opt *Option) {
        opt.IsToHex = v
    }
}