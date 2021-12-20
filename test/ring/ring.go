package ring

import (
	"errors"
	"fmt"
	"hash/crc32"
	"sort"
	"sync"
)

type ConsistentHash struct {
	nodeList nodes
	sync.Mutex
}

type node struct {
	id     string
	HashId uint32
}

type nodes []node

func (n nodes) Len() int           { return len(n) }
func (n nodes) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n nodes) Less(i, j int) bool { return n[i].HashId < n[j].HashId }

var ErrNodeNotFound = errors.New("node not found")

func NewConsistentHash(nodeIds []string) *ConsistentHash {
	return &ConsistentHash{
		nodeList: createNodes(nodeIds),
		Mutex:    sync.Mutex{},
	}
}

func (r *ConsistentHash) AddNode(nodeId string) {
	r.Mutex.Lock()
	defer r.Unlock()

	node := node{
		id:     nodeId,
		HashId: hashId(nodeId),
	}

	r.nodeList = append(r.nodeList, node)

	sort.Sort(r.nodeList)
}

func (r *ConsistentHash) RemoveNode(nodeId string) error {
	r.Mutex.Lock()
	defer r.Unlock()

	i := r.search(nodeId)

	if nodeId != r.nodeList[i].id {
		return ErrNodeNotFound
	}

	r.nodeList = append(r.nodeList[:i], r.nodeList[:i]...)

	return nil
}

func (r *ConsistentHash) GetNode(id string) string {
	i := r.search(id)

	return r.nodeList[i].id
}

func (r *ConsistentHash) search(id string) int {
	hash := hashId(id)

	searchFn := func(i int) bool {
		return r.nodeList[i].HashId > hash
	}

	i := sort.Search(r.nodeList.Len(), searchFn)

	if i >= r.nodeList.Len() {
		i = 0
	}

	return i
}

func createNodes(nodeIds []string) nodes {
	n := make(nodes, 0)

	for _, v := range nodeIds {
		n = append(n, node{id: v, HashId: hashId(v)})
	}

	sort.Sort(n)
	fmt.Println(n)
	return n
}

func hashId(id string) uint32 {
	return crc32.ChecksumIEEE([]byte(id))
}
