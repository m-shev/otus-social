package consistent

import (
	"errors"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"hash/crc32"
	"math"
	"sort"
	"sync"
)

var printer = message.NewPrinter(language.Russian)

type ConsistentHash struct {
	nodeList         nodes
	initialNodeCount uint32
	prevStep         uint32
	sync.Mutex
}

type node struct {
	id     string
	hashId uint32
	step   uint32
}

type nodes []*node

func (n nodes) Len() int           { return len(n) }
func (n nodes) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n nodes) Less(i, j int) bool { return n[i].hashId < n[j].hashId }

var ErrNodeNotFound = errors.New("node not found")

func NewConsistentHash(nodeIds []string, initialNodeCount uint32) *ConsistentHash {
	step := math.MaxUint32 / initialNodeCount
	c := &ConsistentHash{
		nodeList:         createNodes(nodeIds, initialNodeCount, step),
		prevStep:         step,
		initialNodeCount: initialNodeCount,
		Mutex:            sync.Mutex{},
	}
	restNodes := nodeIds[initialNodeCount:]
	for _, v := range restNodes {
		c.AddNode(v)
	}

	return c
}

func (c *ConsistentHash) AddNode(nodeId string) {
	c.Lock()
	defer c.Unlock()
	nextStep := c.prevStep / uint32(2)

	target := c.searchNodeByStep()

	if target == nil {
		target = c.nodeList[0]
		c.prevStep = nextStep
	}

	node := &node{
		id:     nodeId,
		step:   nextStep,
		hashId: target.hashId + nextStep,
	}

	target.step = nextStep

	c.nodeList = append(c.nodeList, node)

	sort.Sort(c.nodeList)
}

func (c *ConsistentHash) searchNodeByStep() *node {
	for _, v := range c.nodeList {
		if v.step == c.prevStep {
			return v
		}
	}

	return nil
}

func (c *ConsistentHash) RemoveNode(nodeId string) error {
	c.Mutex.Lock()
	defer c.Unlock()

	i := c.search(nodeId)

	if nodeId != c.nodeList[i].id {
		return ErrNodeNotFound
	}

	c.nodeList = append(c.nodeList[:i], c.nodeList[:i]...)

	return nil
}

func (c *ConsistentHash) GetNode(id string) string {
	i := c.search(id)

	return c.nodeList[i].id
}

func (c *ConsistentHash) search(id string) int {
	hash := hashId(id)

	searchFn := func(i int) bool {
		return c.nodeList[i].hashId > hash
	}

	i := sort.Search(c.nodeList.Len(), searchFn)

	if i >= c.nodeList.Len() {
		i = 0
	}

	return i
}

func createNodes(nodeIds []string, initialNodeCount uint32, step uint32) nodes {
	n := make(nodes, 0)
	printer.Println(math.MaxUint32)
	printer.Println("step", step)
	idCount := uint32(0)

	idsLen := uint32(len(nodeIds))

	for i := uint32(0); i < initialNodeCount && i < idsLen; i++ {
		n = append(n, &node{
			id:     nodeIds[i],
			hashId: idCount,
			step:   step,
		})

		idCount += step
	}

	//for _, v := range nodeIds {
	//	n = append(n, node{id: v, hashId: start})
	//	start += step
	//}

	sort.Sort(n)

	for _, v := range n {

		_, _ = printer.Printf("%s: %d\n", v.id, v.hashId)
	}
	return n
}

func hashId(id string) uint32 {
	return crc32.ChecksumIEEE([]byte(id))
}
