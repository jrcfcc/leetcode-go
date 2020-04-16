package tree

import (
	"fmt"
	"testing"
)

func Test_depth(t *testing.T) {
	fmt.Println(minDepth(&TreeNode{Val:1,Left:&TreeNode{Val:2}}))
}