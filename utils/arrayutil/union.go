package arrayutil

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/1/11 14:46
 * @Desc:
 */

// Union 并集：将两个切片字符串合并为一个新的切片，并去除重复的元素。可以使用map来实现去重
func Union(s1 []string, s2 []string) []string {
	union := make(map[string]bool)
	for _, str := range s1 {
		union[str] = true
	}
	for _, str := range s2 {
		union[str] = true
	}
	result := make([]string, 0, len(union))
	for str := range union {
		result = append(result, str)
	}
	return result
}
