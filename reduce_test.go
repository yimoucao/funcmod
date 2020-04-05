package funcmod_test

import (
	"fmt"
	"testing"

	"github.com/yimoucao/funcmod"
)

func TestReduce(t *testing.T) {
	rf := func(x, y interface{}) interface{} {
		a := x.(int)
		b := y.(int)
		return a + b
	}
	res := funcmod.Reduce(funcmod.IterList([]interface{}{1, 2, 3, 4}), rf, 0)
	fmt.Println(res)
}
