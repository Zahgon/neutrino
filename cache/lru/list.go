// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Copyright (c) 2017-2022 Lightning Labs

// Package list implements a doubly linked list.
//
// To iterate over a list (where l is a *List):
//
//	for e := l.Front(); e != nil; e = e.Next() {
//		// do something with e.Value
//	}
package lru

// Element is an element of a linked list.
type Element[V any] struct {
	// Next and previous pointers in the doubly-linked list of elements.
	// To simplify the implementation, internally a list l is implemented
	// as a ring, such that &l.root is both the next element of the last
	// list element (l.Back()) and the previous element of the first list
	// element (l.Front()).
	next, prev *Element[V]

	// The list to which this element belongs.
	list *List[V]

	// The value stored with this element.
	Value V
}

// Next returns the next list element or nil.
func (e *Element[V]) Next() *Element[V] { _ = "STUB: not implemented"; return nil }

// Prev returns the previous list element or nil.
func (e *Element[V]) Prev() *Element[V] { _ = "STUB: not implemented"; return nil }

// List represents a doubly linked list.
// The zero value for List is an empty list ready to use.
type List[V any] struct {
	root Element[V] // sentinel list element, only &root, root.prev, and root.next are used
	len  int        // current list length excluding (this) sentinel element
}

// Init initializes or clears list l.
func (l *List[V]) Init() *List[V] { _ = "STUB: not implemented"; return nil }

// NewList returns an initialized list.
func NewList[V any]() *List[V] { _ = "STUB: not implemented"; return nil }

// Len returns the number of elements of list l.
// The complexity is O(1).
func (l *List[V]) Len() int {
	_ = "STUB: not implemented"

	// Front returns the first element of list l or nil if the list is empty.
	return 0
}

func (l *List[V]) Front() *Element[V] { _ = "STUB: not implemented"; return nil }

// Back returns the last element of list l or nil if the list is empty.
func (l *List[V]) Back() *Element[V] { _ = "STUB: not implemented"; return nil }

// lazyInit lazily initializes a zero List value.
func (l *List[V]) lazyInit() { _ = "STUB: not implemented"; return }

// insert inserts e after at, increments l.len, and returns e.
func (l *List[V]) insert(e, at *Element[V]) *Element[V] { _ = "STUB: not implemented"; return nil }

// insertValue is a convenience wrapper for insert(&Element{Value: v}, at).
func (l *List[V]) insertValue(v V, at *Element[V]) *Element[V] {
	_ = "STUB: not implemented"
	return nil
}

// remove removes e from its list, decrements l.len
func (l *List[V]) remove(e *Element[V]) { _ = "STUB: not implemented"; return }

// avoid memory leaks
// avoid memory leaks

// move moves e to next to at.
func (l *List[V]) move(e, at *Element[V]) { _ = "STUB: not implemented"; return }

// Remove removes e from l if e is an element of list l.
// It returns the element value e.Value.
// The element must not be nil.
func (l *List[V]) Remove(e *Element[V]) any {
	_ = "STUB: not implemented"

	// if e.list == l, l must have been initialized when e was inserted
	// in l or l == nil (e is a zero Element) and l.remove will crash
	return *new(any)
}

// PushFront inserts a new element e with value v at the front of list l and returns e.
func (l *List[V]) PushFront(v V) *Element[V] { _ = "STUB: not implemented"; return nil }

// PushBack inserts a new element e with value v at the back of list l and returns e.
func (l *List[V]) PushBack(v V) *Element[V] { _ = "STUB: not implemented"; return nil }

// InsertBefore inserts a new element e with value v immediately before mark and returns e.
// If mark is not an element of l, the list is not modified.
// The mark must not be nil.
func (l *List[V]) InsertBefore(v V, mark *Element[V]) *Element[V] {
	_ = "STUB: not implemented"
	return nil
}

// see comment in List.Remove about initialization of l

// InsertAfter inserts a new element e with value v immediately after mark and returns e.
// If mark is not an element of l, the list is not modified.
// The mark must not be nil.
func (l *List[V]) InsertAfter(v V, mark *Element[V]) *Element[V] {
	_ = "STUB: not implemented"
	return nil
}

// see comment in List.Remove about initialization of l

// MoveToFront moves element e to the front of list l.
// If e is not an element of l, the list is not modified.
// The element must not be nil.
func (l *List[V]) MoveToFront(e *Element[V]) { _ = "STUB: not implemented"; return }

// see comment in List.Remove about initialization of l

// MoveToBack moves element e to the back of list l.
// If e is not an element of l, the list is not modified.
// The element must not be nil.
func (l *List[V]) MoveToBack(e *Element[V]) { _ = "STUB: not implemented"; return }

// see comment in List.Remove about initialization of l

// MoveBefore moves element e to its new position before mark.
// If e or mark is not an element of l, or e == mark, the list is not modified.
// The element and mark must not be nil.
func (l *List[V]) MoveBefore(e, mark *Element[V]) { _ = "STUB: not implemented"; return }

// MoveAfter moves element e to its new position after mark.
// If e or mark is not an element of l, or e == mark, the list is not modified.
// The element and mark must not be nil.
func (l *List[V]) MoveAfter(e, mark *Element[V]) { _ = "STUB: not implemented"; return }

// PushBackList inserts a copy of another list at the back of list l.
// The lists l and other may be the same. They must not be nil.
func (l *List[V]) PushBackList(other *List[V]) { _ = "STUB: not implemented"; return }

// PushFrontList inserts a copy of another list at the front of list l.
// The lists l and other may be the same. They must not be nil.
func (l *List[V]) PushFrontList(other *List[V]) { _ = "STUB: not implemented"; return }
