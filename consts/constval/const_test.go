package constdef_test

import (
	"github.com/leeprince/goinfra/consts/constdef"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestConst(t *testing.T) {
	Convey("TestConst", t, func() {
		Convey("ConstWithInt", func() {
			c := constdef.NewInt(1, "管理后台")
			So(c.Name(), ShouldEqual, "管理后台")
			So(c.Value(), ShouldEqual, int(1))
		})

		Convey("ConstWithString", func() {
			c := constdef.NewString("1001", "管理后台")
			So(c.Name(), ShouldEqual, "管理后台")
			So(c.Value(), ShouldEqual, "1001")

		})

		Convey("Group", func() {
			group := constdef.NewIntGroup(
				constdef.NewInt(1001, "xxx1管理"),
				constdef.NewInt(1002, "xxx2管理"),
				constdef.NewInt(1003, "xxx3管理"),
			)

			c, ok := group.Get(1001)
			So(ok, ShouldBeTrue)
			So(c.Value(), ShouldEqual, 1001)
			So(c.Name(), ShouldEqual, "xxx管理")

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
