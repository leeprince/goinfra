package osinfo

import (
	"fmt"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/12 00:52
 * @Desc:
 */

func Test_GetOsUniqueId(t *testing.T) {
	mark, err := GetOsUniqueId()
	fmt.Println(err)
	fmt.Println(mark)
}

func Test_GetOsUniqueIdBase64(t *testing.T) {
	mark, err := GetOsUniqueIdBase64()
	fmt.Println(err)
	fmt.Println(mark)
}

func Test_CheckOsUniqueIdBase64(t *testing.T) {
	in := "RkRGQzcyN0ItRUIwQi01QTBFLTkzNjYtREI1ODM3RDQ2MTlCOzIyLjIuMDtsZWVwcmluY2VtYWNib29rcHJvLmxvY2FsOw=="
	b := CheckOsUniqueIdBase64(in)
	fmt.Println(b)
}

func Test_CheckOsUniqueIdBase64InArr(t *testing.T) {
	in := []string{
		"RkRGQzcyN0ItRUIwQi01QTBFLTkzNjYtREI1ODM3RDQ2MTlCOzIyLjIuMDtsZWVwcmluY2VtYWNib29rcHJvLmxvY2FsOw==",
	}
	b := CheckOsUniqueIdBase64InArr(in)
	fmt.Println(b)
}
