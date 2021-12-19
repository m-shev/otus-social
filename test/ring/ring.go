package ring

import (
	"fmt"
	"hash/crc32"
	"sort"
)

type Ring struct {
}

type Node struct {
	id     string
	HashId uint32
}

type Nodes []Node

func (n Nodes) Len() int           { return len(n) }
func (n Nodes) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n Nodes) Less(i, j int) bool { return n[i].HashId < n[j].HashId }

func NewRing(nodeIds []string) *Ring {
	fmt.Println(createNodes(nodeIds))
	return &Ring{}
}

func createNodes(nodeIds []string) Nodes {
	nodes := make(Nodes, 0)

	for _, v := range nodeIds {
		nodes = append(nodes, Node{id: v, HashId: hashId(v)})
	}

	sort.Sort(nodes)
	return nodes
}

func hashId(id string) uint32 {
	return crc32.ChecksumIEEE([]byte(id))
}
