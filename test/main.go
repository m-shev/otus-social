package main

import (
	"fmt"
	"github.com/m-shev/otus-social/test/consistent"
	"github.com/satori/go.uuid"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"log"
	"sort"
)

var shard1 = consistent.NodeConfig{NodeId: "message_shard_1", TargetId: ""}
var shard2 = consistent.NodeConfig{NodeId: "message_shard_2", TargetId: ""}
var shard3 = consistent.NodeConfig{NodeId: "message_shard_3", TargetId: ""}
var shard4 = consistent.NodeConfig{NodeId: "message_shard_4", TargetId: ""}
var shard5 = consistent.NodeConfig{NodeId: "message_shard_5", TargetId: "message_shard_2"}
var shard6 = consistent.NodeConfig{NodeId: "message_shard_6", TargetId: "message_shard_5"}

func main() {
	config := []consistent.NodeConfig{shard1, shard2, shard3, shard4}
	c, err := consistent.NewRing(config)
	if err != nil {
		log.Fatalln(err)
	}

	dic := make(map[string]int)

	for _, v := range config {
		dic[v.NodeId] = 0
	}
	fmt.Println(config)
	fmt.Println(dic)

	for i := 0; i <= 100_000_000; i++ {
		node := c.GetNode(uuid.NewV4().String())
		if _, ok := dic[node]; !ok {
			panic(fmt.Sprintf("node %s not found in dictionary"))
		}

		dic[node]++
	}

	for k, v := range dic {
		p := message.NewPrinter(language.Russian)
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
