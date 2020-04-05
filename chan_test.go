package funcmod_test

import (
	"fmt"
	"testing"

	"github.com/yimoucao/funcmod"
)

func TestChan(t *testing.T) {
	ch := make(chan interface{}, 4)
	for i := 0; i < 4; i++ {
		ch <- i
	}
	close(ch)
	iterable := funcmod.IterChan(ch)
	fmt.Println(funcmod.List(iterable))
}
