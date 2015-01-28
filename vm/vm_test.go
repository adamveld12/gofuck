package vm

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestInputOutput(t *testing.T) {
	Convey("For the given programs that takes input", t, func() {
		Convey("When input 66 '+' output should be 'B'", func() {
			program := "++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++."
			_, output := Execute(program)
			result := <-output
			So(result, ShouldEqual, 'B')
		})

		Convey("2 should print 2", func() {
			program := ",."
			input, output := Execute(program)
			input <- 2
			result := <-output
			So(result, ShouldEqual, 2)
		})

		Convey("2 plus 3 should print 5", func() {
			program := ",>,<[->+<]>."
			input, output := Execute(program)
			input <- 2
			input <- 3
			result := <-output
			So(result, ShouldEqual, 5)
		})
	})
}

func TestForeignCharacters(t *testing.T) {
	Convey("For the given programs that have foreign characters", t, func() {
		Convey("should print 2 with ,abce.", func() {
			program := ",abcde."
			input, output := Execute(program)
			input <- 2
			result := <-output
			So(result, ShouldEqual, 2)
		})

		Convey("should print 5 with 3 and 2 using ',!>,?<[->+898<]asd>.'", func() {
			program := ",!>,?<[->+898<]asd>."
			input, output := Execute(program)
			input <- 2
			input <- 3
			result := <-output
			So(result, ShouldEqual, 5)
		})
	})
}

func TestLoop(t *testing.T) {
	Convey("When looping", t, func() {
		program := "++++++[>++++++++++<-]>+++++."
		Convey("Output should be 'A'", func() {
			_, output := Execute(program)
			result := <-output
			So(result, ShouldEqual, 'A')
		})
	})
}

func TestOutput(t *testing.T) {
}
