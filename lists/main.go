// Package nlists the technique i learned from the golang implmentation
// note: this is not the copy of the whole but the study of golang lists implmentation
package nlists

import (
	"cmp"
	"log"
)

type Element struct {
	next  *Element
	prev  *Element
	slist *SingleList
	val   int
}
type SingleList struct {
	root Element
	len  int
}

func (li *SingleList) Insert(to *Element, val int) *Element {
	e := &Element{val: val}
	e.next = to.next
	to.next = e
	e.slist = li
	li.len++
	return e
}

// Len returns the len
func (li *SingleList) Len() int {
	return li.len
}

// InsertValue unorthodx way of writing linked list but i am still learning it from golang implmentation
func (li *SingleList) InsertValue(val int) *Element {
	return li.Insert(&li.root, val)
}

// Print prints the data
func (li *SingleList) Print() {
	current := li.root.next
	for current != nil {
		log.Println("single list values: ", current.val)
		current = current.next
	}
}

// DeleteNode delets the node
func (li *SingleList) DeleteNode() bool {
	if li == nil {
		return true
	}
	delete := li.root.next
	li.root.next = delete.next
	delete.next = nil
	return true
}

// OnPosErase removes the value on give pos true if removed
func (li *SingleList) OnPosErase(pos int) bool {
	if pos > li.Len() || pos < 0 {
		return true
	}

	current := &li.root
	for i := 0; i < pos-1; i++ {
		if current.next == nil {
			return false
		}
		current = current.next
	}

	target := current.next     // get the value
	current.next = target.next // erase the value

	if li.len > 0 {
		li.len--
	}

	// garbage collector
	target.next = nil
	current.next = nil
	return true
}

// Erase erases the prev pushed node
func (li *SingleList) Erase() {
	li.root.next = nil
}

// DoMerge combing the list
// what this bascially is adds the node at the last-nth pos
func (li *SingleList) DoMerge(with *Element) *SingleList {
	if li == nil {
		return nil
	} else if with == nil {
		return li
	}
	current := &li.root
	for current.next != nil {
		current = current.next
	}
	current.next = with // place it after last node

	return li
}

// Merge merges the list
func (li *SingleList) Merge(with *SingleList) *SingleList {
	return li.DoMerge(&with.root)
}

// Sort sorts the list in ascending order
func (li *SingleList) Sort() {
	li.root = *SingleListMergeSort(&li.root)
}

// SingleListMergeSort the implmentation requires the basic knowdledge of merge sort
// understanding:
// find min,max
// free up min
// start going ahead till the nil pos
func SingleListMergeSort(root *Element) *Element {
	if root == nil || root.next == nil {
		return root
	}

	min := Mid(root)
	max := min.next
	min.next = nil
	left := SingleListMergeSort(root)
	right := SingleListMergeSort(max)
	return _Merge(right, left)
}

// Mid finds the mid of the list
func Mid(root *Element) *Element {
	fast, slow := root, root
	for fast != nil && fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}

// _Merge merge algorithm
// implments the normal binary search which is played around the value
func _Merge(a, b *Element) *Element {
	dummy := &Element{}
	max := dummy
	for a != nil && b != nil {
		if a.val < b.val {
			max.next = a
			a = a.next
		} else {
			max.next = b
			b = b.next
		}
		max = max.next
	}
	// note this is imp cause the the for loop is break from a or b to nil
	if a != nil {
		max.next = a
	} else {
		max.next = b
	}
	return dummy.next
}

// MergeSort performs sorting
func MergeSort[T cmp.Ordered](data []T) []T {
	if len(data) < 2 {
		return data
	}
	mid := len(data) / 2
	min := MergeSort(data[:mid])
	max := MergeSort(data[mid:])
	return Merge(min, max)
}

// Merge implments the binary search
func Merge[T cmp.Ordered](a, b []T) []T {

	res := make([]T, len(a)+len(b))
	i, j, k := 0, 0, 0

	for i < len(a) && j < len(b) {
		if cmp.Less(a[i], b[j]) {
			res[k] = a[i]
			i++
		} else {
			res[k] = b[j]
			j++
		}
		k++
	}
	if i < len(a) {
		copy(res[k:], a[i:])
	} else {
		copy(res[k:], b[j:])
	}
	return res
}

// IsSort returns true if sorted
func IsSort[T cmp.Ordered](data []T) bool {

	for i := 1; i < len(data); i++ {
		if data[i] < data[i-1] {
			return false
		}
	}
	return true
}
