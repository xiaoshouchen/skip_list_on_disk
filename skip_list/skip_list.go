package skip_list

import (
	"errors"
	"math/bits"
	"math/rand"
	"time"

	"github.com/xiaoshouchen/skip-list/internal"
)

var NotFoundErr = errors.New("the key is not found")

type SkipListI interface {
	Get(key string) (string, error)
	Insert(key, value string)
	Remove(key string) error
	Size() uint
}

type SkipList struct {
	level uint
	size  uint
	head  *skipListNode
}

type skipListNode struct {
	key   string
	value string
	next  []*skipListNode
}

func NewSkipList(defaultLevel uint) SkipListI {
	return &SkipList{head: newSkipListNode("", "", defaultLevel), level: defaultLevel}
}

func (sl *SkipList) Get(key string) (string, error) {
	node := sl.findNode(key, nil)
	if node == nil || node.key != key {
		return "", NotFoundErr
	}
	return node.value, nil
}

// Insert insert new node
func (sl *SkipList) Insert(key, value string) {
	// 查找可以插入的位置
	prevArr := make([]*skipListNode, sl.level)
	node := sl.findNode(key, prevArr)

	// 如果key已经存在，则更新对应的值
	if node != nil && node.key == key {
		node.value = value
		return
	}

	//开始插入node
	level := sl.randomLevel()
	node = newSkipListNode(key, value, level)

	// re-balance
	//节点level的自平衡
	if level > sl.level {
		// Increase the level
		for i := sl.level; i < level; i++ {
			sl.head.next[i] = node
		}
		sl.level = level
	}

	//插入
	for i := uint(0); i < internal.Min(sl.level, level); i++ {
		node.next[i] = prevArr[i].next[i]
		prevArr[i].next[i] = node
	}

	sl.size++
}

// Remove delete a node
func (sl *SkipList) Remove(key string) error {
	prevArr := make([]*skipListNode, sl.level)
	node := sl.findRemoveNode(key, prevArr)
	if node == nil {
		return NotFoundErr
	}
	for i, v := range node.next {
		prevArr[i].next[i] = v
	}

	for sl.level > 1 && sl.head.next[sl.level-1] == nil {
		sl.level--
	}
	sl.size--
	return nil
}

// Size return the number of skip list element
func (sl *SkipList) Size() uint {
	return sl.size

}

// findNode 查找SL节点
func (sl *SkipList) findNode(key string, prevArr []*skipListNode) *skipListNode {
	prev := sl.head
	level := sl.level - 1
	for i := level; i >= 0; i-- {
		for next := prev.next[i]; next != nil; next = next.next[i] {
			if next.key == key {
				return next
			}
			if next.key > key {
				break
			}
			prev = next
		}

		// 如果是Get方法，则prevArr为nil，节省内存
		if prevArr != nil {
			prevArr[i] = prev
		}

		//防止uint类型溢出
		if i == 0 {
			break
		}

	}
	return nil
}

// findRemoveNode 查找SL节点
func (sl *SkipList) findRemoveNode(key string, prevArr []*skipListNode) *skipListNode {
	var resNode *skipListNode
	prev := sl.head
	level := sl.level - 1
	for i := level; i >= 0; i-- {
		for next := prev.next[i]; next != nil; next = next.next[i] {
			if next.key == key {
				resNode = new(skipListNode)
				resNode = next
				break
			}
			if next.key > key {
				break
			}
			prev = next
		}

		prevArr[i] = prev

		//防止uint类型溢出
		if i == 0 {
			break
		}

	}
	return resNode
}

// randomLevel 随机生成sl的高度
func (sl *SkipList) randomLevel() uint {
	total := uint64(1)<<uint64(sl.level) - 1 // 2^n-1
	rand.Seed(time.Now().UnixNano())
	k := rand.Uint64() & total
	level := sl.level - uint(bits.Len64(k)) + 1

	for level > 3 && 1<<(level-3) > sl.size {
		level--
	}
	return level
}

func newSkipListNode(key, value string, maxHeight uint) *skipListNode {
	nextNodes := make([]*skipListNode, maxHeight)
	return &skipListNode{key: key, value: value, next: nextNodes}
}
