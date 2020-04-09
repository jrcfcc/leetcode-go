package dynamic

import (
	"fmt"
	"math/big"
	"testing"
)

func Test_pow(t *testing.T) {
	fmt.Println(pow(big.NewInt(3),4))
}