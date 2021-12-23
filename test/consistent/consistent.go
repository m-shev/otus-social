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

type Ring struct {
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

type NodeConfig struct {
	NodeId   string
	TargetId string
}

type nodes []*node

func (n nodes) Len() int           { return len(n) }
func (n nodes) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n nodes) Less(i, j int) bool { return n[i].hashId < n[j].hashId }

var ErrNodeNotFound = errors.New("node not found")
var ErrTargetShouldBeNotEmpty = errors.New("to add a node the target must be not empty")

func NewRing(configs []NodeConfig) (*Ring, error) {
	count := uint32(countInitialNodes(configs))
	step := math.MaxUint32 / count

	r := &Ring{
		nodeList:         createNodes(configs[:count], step),
		prevStep:         step,
		initialNodeCount: count,
		Mutex:            sync.Mutex{},
	}
	err := r.AddNodes(configs[count:])

	if err != nil {
		return nil, err
	}

	for _, v := range r.nodeList {

		_, _ = printer.Printf("%s: %d\n", v.id, v.hashId)
	}

	return r, nil
}

func (r *Ring) AddNodes(configs []NodeConfig) error {
	for _, v := range configs {
		err := r.AddNode(v)

		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Ring) AddNode(config NodeConfig) error {
	r.Lock()
	defer r.Unlock()
	target, err := r.searchTarget(config.TargetId)

	if err != nil {
		return err
	}

	target.step = target.step / uint32(2)

	n := &node{
		id:     config.NodeId,
		step:   target.step,
		hashId: target.hashId + target.step,
	}

	r.nodeList = append(r.nodeList, n)

	sort.Sort(r.nodeList)

	return nil
}

func (r *Ring) searchTarget(targetId string) (*node, error) {
	var target *node

	for _, v := range r.nodeList {
		if targetId == v.id {
			target = v
			break
		}
	}

	if target == nil {
		return nil, ErrTargetShouldBeNotEmpty
	}

	return target, nil
}

func (r *Ring) RemoveNode(nodeId string) error {
	r.Mutex.Lock()
	defer r.Unlock()

	i := r.search(nodeId)

	if nodeId != r.nodeList[i].id {
		return ErrNodeNotFound
	}

	r.nodeList = append(r.nodeList[:i], r.nodeList[:i]...)

	return nil
}

func (r *Ring) GetNode(id string) string {
	i := r.search(id)

	return r.nodeList[i].id
}

func (r *Ring) search(id string) int {
	hash := hashId(id)

	searchFn := func(i int) bool {
		return r.nodeList[i].hashId > hash
	}

	i := sort.Search(r.nodeList.Len(), searchFn)

	if i >= r.nodeList.Len() {
		i = 0
	}

	return i
}

func createNodes(configs []NodeConfig, step uint32) nodes {
	n := make(nodes, 0)
	printer.Println(math.MaxUint32)
	printer.Println("step", step)
	idCount := uint32(0)

	idsLen := uint32(len(configs))

	for i := uint32(0); i < idsLen; i++ {
		n = append(n, &node{
			id:     configs[i].NodeId,
			hashId: idCount,
			step:   step,
		})

		idCount += step
	}

	sort.Sort(n)

	return n
}

func hashId(id string) uint32 {
	return crc32.ChecksumIEEE([]byte(id))
}

func countInitialNodes(nodes []NodeConfig) int {
	count := 0

	for _, v := range nodes {
		if v.TargetId != "" {
			break
		}

		count++
	}

	return count
}
