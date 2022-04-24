package concurrentqueue

import "sync"

type ConcurrentQueue struct {
	data []interface{}
	mux  sync.Mutex
}

func New() *ConcurrentQueue {
	return &ConcurrentQueue{
		data: make([]interface{}, 0),
	}
}

func (c *ConcurrentQueue) Enqueue(v interface{}) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.data = append(c.data, v)
}

func (c *ConcurrentQueue) Dequeue() (interface{}, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	if len(c.data) > 0 {
		v := c.data[0]
		c.data = c.data[1:len(c.data)]
		return v, true
	} else {
		return nil, false
	}
}
func (c *ConcurrentQueue) Clear() {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.data = make([]interface{}, 0)
}

func (c *ConcurrentQueue) Lenth() int {
	return len(c.data)
}
