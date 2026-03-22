package poolx_test

import (
	"fmt"
	"github.com/vikash-paf/goutils/poolx"
)

func ExampleTypedPool() {
	p := poolx.NewTypedPool(func() []byte {
		return make([]byte, 1024)
	})

	buf := p.Get()
	fmt.Println(len(buf))
	
	p.Put(buf)
	// Output: 1024
}
