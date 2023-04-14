package constval_test

import (
	"github.com/leeprince/goinfra/consts/constval"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestConst(t *testing.T) {
	Convey("TestConst", t, func() {
		Convey("ConstWithInt", func() {
			c := constval.NewInt(1, "管理后台")
			So(c.Name(), ShouldEqual, "管理后台")
			So(c.Value(), ShouldEqual, int(1))
		})

		Convey("ConstWithString", func() {
			c := constval.NewString("1001", "管理后台")
			So(c.Name(), ShouldEqual, "管理后台")
			So(c.Value(), ShouldEqual, "1001")

		})

		Convey("Group", func() {
			group := constval.NewIntGroup(
				constval.NewInt(1001, "xxx1管理"),
				constval.NewInt(1002, "xxx2管理"),
				constval.NewInt(1003, "xxx3管理"),
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
