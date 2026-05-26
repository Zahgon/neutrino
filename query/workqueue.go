package query

// Task is an interface that has a method for returning their index in the
// work queue.
type Task interface {
	// Index returns this Task's index in the work queue.
	Index() uint64
}

// workQueue is struct implementing the heap interface, and is used to keep a
// list of remaining queryTasks in order.
type workQueue struct {
	tasks []Task
}

// Len returns the number of nodes in the priority queue.
//
// NOTE: This is part of the heap.Interface implementation.
func (w *workQueue) Len() int {
	_ = "STUB: not implemented"

	// Less returns whether the item in the priority queue with index i should sort
	// before the item with index j.
	//
	// NOTE: This is part of the heap.Interface implementation.
	return 0
}

func (w *workQueue) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// Swap swaps the nodes at the passed indices in the priority queue.
//
// NOTE: This is part of the heap.Interface implementation.
func (w *workQueue) Swap(i, j int) { _ = "STUB: not implemented"; return }

// Push add x as element Len().
//
// NOTE: This is part of the heap.Interface implementation.
func (w *workQueue) Push(x interface{}) { _ = "STUB: not implemented"; return }

// Pop removes and returns element Len()-1.
//
// NOTE: This is part of the heap.Interface implementation.
func (w *workQueue) Pop() interface{} { _ = "STUB: not implemented"; return nil }

// Peek returns the first item in the queue.
func (w *workQueue) Peek() interface{} { _ = "STUB: not implemented"; return nil }
