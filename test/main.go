package main

import (
	"fmt"
	"github.com/m-shev/otus-social/test/ring"
	"sort"
	"strconv"
)

func main() {
	m := []string{"1", "2", "3", "4", "5"}
	c := ring.NewConsistentHash(m)
	fmt.Println()
	dic := make(map[string]int)

	for _, v := range m {
		dic[v] = 0
	}
	fmt.Println(m)
	fmt.Println(dic)

	for i := 0; i <= 1000; i++ {
		node := c.GetNode(strconv.Itoa(i))
		if _, ok := dic[node]; !ok {
			panic(fmt.Sprintf("node %s not found in dictionary"))
		}

		dic[node]++
	}

	for k, v := range dic {
		fmt.Println(k, ":", v)
	}
}

func search(elm int) int {
	arr := []int{2, 3, 5, 6, 7, 12, 34, 54, 64, 76, 98, 101, 202, 303, 404}

	searchFn := func(i int) bool {
		return arr[i] >= elm
	}

	return sort.Search(len(arr), searchFn)
}
