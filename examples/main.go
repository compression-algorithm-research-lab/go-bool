package main

import (
	"fmt"
	"github.com/compression-algorithm-research-lab/go-bool/bool_array"
)

func main() {

	array := bool_array.New(10)
	array.Set(1, true)
	fmt.Println(array.Get(1))

}
