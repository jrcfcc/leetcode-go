package nums

import (
	"fmt"
	"strings"
	"testing"
)

func Test_chinesetonums(t *testing.T) {
	a := chinesetonums("一百零三万五千零二")
	if a != 1035002 {
		t.Errorf("expected 1035002 got %d",a)
	}
	a = chinesetonums("十三万零二百")
	if a != 130200 {
		t.Errorf("expected 130200 got %d",a)
	}
	fmt.Println(len(strings.Split("一亿亿","亿")))

	var max = ^uint(0) >> 1
	fmt.Println(max)
}