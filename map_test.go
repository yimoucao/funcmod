package funcmod_test

import (
	"fmt"
	"testing"

	"github.com/yimoucao/funcmod"
)

func TestMap(t *testing.T) {
	mfunc := func(x interface{}) interface{} {
		v := x.(int)
		v++
		return v
	}
	res := funcmod.Map(mfunc, funcmod.IterList([]interface{}{1, 2, 3, 4}))
	fmt.Println(funcmod.List(res))
}
