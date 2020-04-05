package funcmod_test

import (
	"fmt"
	"testing"

	"github.com/yimoucao/funcmod"
)

func TestChain(t *testing.T) {
	var res funcmod.Iterable
	res = funcmod.Chain()
	fmt.Println(funcmod.List(res))

	res = funcmod.Chain(funcmod.IterList([]interface{}{1, 2, 3, 4}))
	fmt.Println(funcmod.List(res))

	res = funcmod.Chain(funcmod.IterList([]interface{}{1, 2, 3, 4}), funcmod.IterList([]interface{}{1, 2, 3, 4}))
	fmt.Println(funcmod.List(res))

	res = funcmod.Chain(funcmod.IterList([]interface{}{1, 2, 3, 4}), funcmod.IterList([]interface{}{1.1, 2.1, 3.1, 4.1}))
	fmt.Println(funcmod.List(res))
}
