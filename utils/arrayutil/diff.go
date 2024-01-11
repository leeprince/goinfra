package arrayutil

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/1/11 14:47
 * @Desc:
 */

// Diff 差集：遍历第一个切片字符串，将不在第二个切片字符串中的元素添加到结果切片中
func Diff(s1 []string, s2 []string) []string {
	difference := make([]string, 0)
	set := make(map[string]bool)
	for _, str := range s2 {
		set[str] = true
	}
	for _, str := range s1 {
		if !set[str] {
			difference = append(difference, str)
		}
	}
	return difference
}
