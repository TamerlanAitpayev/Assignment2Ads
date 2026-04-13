type CircularQueue struct {
    data  []int
    head  int
    tail  int
    size  int
    cap   int
}

func NewCircularQueue(cap int) *CircularQueue {
    return &CircularQueue{
        data: make([]int, cap),
        cap:  cap,
    }
}

func (q *CircularQueue) Enqueue(val int) bool {
    if q.IsFull() {
        return false
    }
    q.data[q.tail] = val
    q.tail = (q.tail + 1) % q.cap
    q.size++
    return true
}

func (q *CircularQueue) Dequeue() (int, bool) {
    if q.IsEmpty() {
        return 0, false
    }
    val := q.data[q.head]
    q.head = (q.head + 1) % q.cap
    q.size--
    return val, true
}

func (q *CircularQueue) IsFull() bool  { return q.size == q.cap }
func (q *CircularQueue) IsEmpty() bool { return q.size == 0 }