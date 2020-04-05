package funcmod_test

import (
	"fmt"
	"testing"

	"github.com/yimoucao/funcmod"
)

func TestFilter(t *testing.T) {
	f := func(x interface{}) bool { v := x.(int); return v%2 == 0 }
	res := funcmod.Filter(f, funcmod.IterList([]interface{}{1, 2, 3, 4}))
	fmt.Println(funcmod.List(res))
}
