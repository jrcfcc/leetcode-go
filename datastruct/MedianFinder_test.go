package datastruct

import (
	"fmt"
	"testing"
)

func TestMedianFinder_AddNum(t *testing.T) {
	var m = Constructor()
	m.AddNum(1)
	fmt.Println(m.max)
	fmt.Println(m.min)
	m.AddNum(2)
	fmt.Println(m.max)
	fmt.Println(m.min)
	m.AddNum(3)
	fmt.Println(m.max)
	fmt.Println(m.min)
}