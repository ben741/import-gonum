package main

import (
	"fmt"

	"github.com/james-bowman/sparse"
	"gonum.org/v1/gonum"
)

func main() {
	v, c := gonum.Version()
	fmt.Println(v, c)

	vec := sparse.NewVector(0, []int{}, []float64{})
	fmt.Println(vec.Len())
}
