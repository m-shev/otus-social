package main

import (
	"fmt"
	"github.com/m-shev/otus-social/test/consistent"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"sort"
	"strconv"
)

func main() {
	m := []string{"1", "2", "3", "4", "5", "6", "7", "8"}
	c := consistent.NewConsistentHash(m, 2)
	fmt.Println()
	dic := make(map[string]int)

	for _, v := range m {
		dic[v] = 0
	}
	fmt.Println(m)
	fmt.Println(dic)

	for i := 0; i <= 10000000; i++ {
		node := c.GetNode(strconv.Itoa(i))
		if _, ok := dic[node]; !ok {
			panic(fmt.Sprintf("node %s not found in dictionary"))
		}

		dic[node]++
	}

	for k, v := range dic {
		p := message.NewPrinter(language.Romanian)
		_, _ = p.Printf("%s: %d\n", k, v)
	}
}

func search(elm int) int {
	arr := []int{2, 3, 5, 6, 7, 12, 34, 54, 64, 76, 98, 101, 202, 303, 404}

	searchFn := func(i int) bool {
		return arr[i] >= elm
	}

	return sort.Search(len(arr), searchFn)
}
