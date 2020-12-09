package user

import (
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

func TestParseMessageAction(t *testing.T) {
	parsedActModel := ParseMessageAction("rename]")

	Convey("rename", t, func() {
		So(parsedActModel.Action, ShouldEqual, "rename")
		So(parsedActModel.Message, ShouldEqual, "")

		parsedActModel2 := ParseMessageAction("rename")

		So(parsedActModel2.Message, ShouldEqual, "rename")
	})

}

func TestTrim_strings(t *testing.T) {
	Convey("trim", t, func() {
		InputContent := "@blue 私聊一下"

		So(strings.HasPrefix(InputContent, "@"), ShouldBeTrue)

		s := strings.TrimLeft(InputContent, "@")
		So(s, ShouldEqual, "blue 私聊一下")

		index := strings.Index(s, " ")

		name := s[0:index]

		So(name, ShouldEqual, "blue")

		msg := s[index+1:]
		So(msg, ShouldEqual, "私聊一下")
	})
}
