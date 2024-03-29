package domain

import (
	"fmt"
	"sync"
)

type HeadElem struct {
	word     string
	count    int
	treeNode *FPNode
	pattern  map[string]int
}

type HeadElems []HeadElem

func (elems HeadElems) Len() int {
	return len(elems)
}
func (elems HeadElems) Less(i, j int) bool {
	if elems[i].count == elems[j].count {
		return elems[i].word < elems[j].word
	} else {
		return elems[i].count < elems[j].count
	}
}

func (elems HeadElems) Swap(i, j int) {
	elems[i], elems[j] = elems[j], elems[i]
}

type FPNode struct {
	word     string
	count    int
	children []*FPNode
	parent   *FPNode
	next     *FPNode
}

func (n *FPNode) Insert(node *FPNode) {
	node.parent = n
	n.children = append(n.children, node)
}

func (n *FPNode) String() string {
	return fmt.Sprintf("<%s,%d>", n.word, n.count)
}

func (n *FPNode) insertHeadElemTreeNode(headAddr map[string]*HeadElem) {
	tmp := headAddr[n.word].treeNode
	if tmp == nil {
		headAddr[n.word].treeNode = n
	} else {
		for tmp.next != nil {
			tmp = tmp.next
		}
		tmp.next = n
	}
}

//type FPTree struct {
//	Root *FPNode
//}
type FPRoot FPNode

func (r *FPRoot) String() string {
	var queue []*FPNode
	var tree string
	queue = append(queue, (*FPNode)(r))
	for len(queue) > 0 {
		node := queue[0]
		tree += fmt.Sprint(node) + " "
		queue = append(queue, node.children...)
		queue = queue[1:]
	}
	tree = tree[:len(tree)-1]
	return tree
}

func (r *FPRoot) BuildFPTree(wordBase [][]string, headAddr map[string]*HeadElem) {
	for _, words := range wordBase {
		var nodes []*FPNode

		for _, word := range words {
			nodes = append(nodes, &FPNode{word: word, count: 1, parent: nil, next: nil})
		}

		r.insertNodeToTree(nodes, headAddr)
	}
}

func (r *FPRoot) insertNodeToTree(nodes []*FPNode, headAddr map[string]*HeadElem) {
	p := (*FPNode)(r)
	for _, node := range nodes {
		notFound := true
		for _, child := range p.children {
			if child.word == node.word {
				child.count = child.count + node.count
				p = child
				notFound = false
				break
			}
		}
		if notFound {
			p.Insert(node)
			node.insertHeadElemTreeNode(headAddr)
			p = node
		}
	}
}

func (r *FPRoot) ConditionalPattern(headElems HeadElems, supportCount int, headAddr map[string]*HeadElem, paraNum int) {
	headChan := make(chan *HeadElem)
	go func() {
		defer close(headChan)
		for i := len(headElems) - 1; i >= 0; i-- {
			headChan <- headAddr[headElems[i].word]
		}
	}()
	var wg sync.WaitGroup

	for i := 0; i < paraNum; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for headElem := range headChan {
				node := headElem.treeNode
				pattern := make(map[string]int)
				for node != nil {
					p := node
					count := p.count
					for p.parent != (*FPNode)(r) {
						pattern[p.parent.word] += count
						p = p.parent
					}
					node = node.next
				}
				for word, count := range pattern {
					if count < supportCount {
						delete(pattern, word)
					}
				}
				headElem.pattern = pattern
			}
		}()
	}
	wg.Wait()
}

type Pair struct {
	key   string
	value int
}

type Pairs []Pair

func (pairs Pairs) Len() int {
	return len(pairs)
}

func (pairs Pairs) Less(i, j int) bool {
	if pairs[i].value == pairs[j].value {
		return pairs[i].key < pairs[j].key
	} else {
		return pairs[i].value < pairs[j].value
	}
}

func (pairs Pairs) Swap(i, j int) {
	pairs[i], pairs[j] = pairs[j], pairs[i]
}
