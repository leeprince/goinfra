package pkg

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/3/10 23:38
 * @Desc:
 */

func Max(l, r int) int {
	if l > r {
		return l
	}
	return r
}

func ReverseSlice(v []int) {
	l := len(v)
	for i := 0; i < l/2; i++ {
		v[i], v[l-1-i] = v[l-1-i], v[i]
	}
}

func SwapInt(v []int, l, r int) {
	v[l], v[r] = v[r], v[l]
}
