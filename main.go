package main

import (
	"fmt"

	"github.com/james-bowman/sparse"
)

func main() {
	vec := sparse.NewVector(0, []int{}, []float64{})
	fmt.Println(vec.Len())
}
