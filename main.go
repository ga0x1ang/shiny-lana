package main

import (
	"fmt"

	"github.com/patriot7/shiny-lana/sort/bucket"
)

func main() {
	result := sort.BucketSort([]int{1, 2, 37, 3, 2, 5, 6, 7, 9})
	fmt.Println(result)
}
