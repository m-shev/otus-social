package main

import (
	"fmt"
	"github.com/m-shev/otus-social/test/ring"
)

func main() {
	m := []string{"message_shard_1", "message_shard_2", "message_shard_3", "message_shard_4", "message_shard_5"}
	_ = ring.NewRing(m)
	fmt.Println()
}
