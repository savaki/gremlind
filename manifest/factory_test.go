package manifest

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRead(t *testing.T) {
	var contents string
	var err error
	var m *Manifest

	Convey("Given manifest contents", t, func() {
		contents = `
		name  = "the name"
		repo  = "the repo"
		notes = "the notes"

		program blue {
			cmd = ["blue-command", "b1", "b2"]
		}

		program red {
			cmd = ["red-command", "r1", "r2"]
		}

		service acme {
			port = 1234
		}

		check website {
			script   = "the script"
			url      = "the url"
			interval = "the interval"
		}

		tags {
			hello = "world"
			argle = "bargle"
		}
		`

		Convey("When I decode the manifest", func() {
			m, err = Read(contents)

			Convey("Then I expect no errors", func() {
				So(err, ShouldBeNil)
			})

			Convey("And I expect the name to be assigned", func() {
				So(m.Name, ShouldEqual, "the name")
			})

			Convey("And I expect the repo to be assigned", func() {
				So(m.Repo, ShouldEqual, "the repo")
			})

			Convey("And I expect the notes to be assigned", func() {
				So(m.Notes, ShouldEqual, "the notes")
			})

			Convey("And I expect the program to be assigned", func() {
				So(len(m.Program), ShouldEqual, 2)

				So(m.Program["red"], ShouldNotBeNil)
				So(m.Program["red"].Cmd, ShouldResemble, []string{"red-command", "r1", "r2"})

				So(m.Program["blue"], ShouldNotBeNil)
				So(m.Program["blue"].Cmd, ShouldResemble, []string{"blue-command", "b1", "b2"})
			})

			Convey("And I expect the service to be assigned", func() {
				So(len(m.Service), ShouldEqual, 1)

				So(m.Service["acme"], ShouldNotBeNil)
				So(m.Service["acme"].Port, ShouldEqual, 1234)
			})

			Convey("And I expect the check to be assigned", func() {
				So(len(m.Check), ShouldEqual, 1)

				So(m.Check["website"], ShouldNotBeNil)
				So(m.Check["website"].Interval, ShouldEqual, "the interval")
				So(m.Check["website"].Script, ShouldEqual, "the script")
				So(m.Check["website"].Url, ShouldEqual, "the url")
			})

			Convey("And I expect the tags to be assigned", func() {
				So(m.Tags["hello"], ShouldEqual, "world")
				So(m.Tags["argle"], ShouldEqual, "bargle")
			})
		})
	})
}
