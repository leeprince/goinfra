/*
 * @Date: 2020-11-03 10:56:01
 * @LastEditors: aiden.deng (Zhenpeng Deng)
 * @LastEditTime: 2020-12-15 16:47:56
 */
package constval2_test

import (
	"github.com/leeprince/goinfra/pkg/constval2"
	"testing"
	
	. "github.com/smartystreets/goconvey/convey"
)

func TestConst(t *testing.T) {
	
	Convey("TestConst", t, func() {
		Convey("ConstWithInt", func() {
			c := constval2.NewInt(1, "管理后台")
			So(c.Name(), ShouldEqual, "管理后台")
			So(c.Value(), ShouldEqual, int(1))
		})
		
		Convey("ConstWithString", func() {
			c := constval2.NewString("1001", "管理后台")
			So(c.Name(), ShouldEqual, "管理后台")
			So(c.Value(), ShouldEqual, "1001")
			
		})
		
		Convey("Group", func() {
			group := constval2.NewIntGroup(
				constval2.NewInt(1001, "销项管理"),
				constval2.NewInt(1002, "进项管理"),
				constval2.NewInt(1003, "电子档案"),
			)
			
			c, ok := group.Get(1001)
			So(ok, ShouldBeTrue)
			So(c.Value(), ShouldEqual, 1001)
			So(c.Name(), ShouldEqual, "销项管理")
			
			c, ok = group.Get(2000)
			So(ok, ShouldBeFalse)
			So(c, ShouldBeNil)
			
			ok = group.IsValid(1002)
			So(ok, ShouldBeTrue)
			
			ok = group.IsValid(1004)
			So(ok, ShouldBeFalse)
			
		})
	})
	
}
