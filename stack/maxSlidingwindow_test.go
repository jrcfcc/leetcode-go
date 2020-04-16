package stack

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_maxSlidingWindow(t *testing.T) {
	Convey("test for max sliding window",t, func() {
		Convey("test 1", func() {
			res := maxSlidingWindow([]int{1,3,-1,-3,5,3,6,7},3)
			fmt.Println(res)
		})
		Convey("test 2", func() {
			res := maxSlidingWindow([]int{1,-1},1)
			fmt.Println(res)
		})
		Convey("test 3", func() {
			res := maxSlidingWindow([]int{1,3,-1,-3,5,3,6,7},3)
			fmt.Println(res)
		})
	})
}