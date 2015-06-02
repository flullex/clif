package cli

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestName(t *testing.T) {
	Convey("Formatting output", t, func() {
		Convey("Formatting with stripped directives", func() {
			f := NewDefaultFormatter(nil)
			s := f.Format("Foo <headline>bar<reset> baz")
			So(s, ShouldEqual, "Foo bar baz")
		})
		Convey("Formatting with formatted directives", func() {
			f := NewDefaultFormatter(map[string]string{
				"headline": "H!",
				"reset":    "R!",
			})
			s := f.Format("Foo <headline>bar<reset> baz")
			So(s, ShouldEqual, "Foo H!barR! baz")
		})
		Convey("Formatting does not replace not registered default tokens", func() {
			f := NewDefaultFormatter(nil)
			s := f.Format("Foo <headline>bar<reset> <baz> boing")
			So(s, ShouldEqual, "Foo bar <baz> boing")
		})
		Convey("Formatting does not replace not registered custom tokens", func() {
			f := NewDefaultFormatter(map[string]string{
				"baz": "BAZ",
			})
			s := f.Format("Foo <headline>bar<reset> <baz> boing")
			So(s, ShouldEqual, "Foo <headline>bar<reset> BAZ boing")
		})
	})
}
