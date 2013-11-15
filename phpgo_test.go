package phpgo

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

// string.go tests

func TestEcho(t *testing.T) {

	Echo("Hello")
}

// filesystem.go tests
func TestFileGetContents(t *testing.T) {
	urlTest := "http://golang.org/"
	fsTest := "README.md"
	failTest := "filethatdoesnotexist.txt.md.exe"

	Convey("HTTP should return a string", t, func() {
		aTest, _ := FileGetContents(urlTest)

		So(aTest, ShouldHaveSameTypeAs, "string")
	})

	Convey("Filesystem should return a string", t, func() {
		bTest, _ := FileGetContents(fsTest)

		So(bTest, ShouldHaveSameTypeAs, "string")
	})

	Convey("Filesystem should err", t, func() {
		_, err := FileGetContents(failTest)

		So(err.Error(), ShouldContainSubstring, "The system cannot find the file specified.")
	})
}
