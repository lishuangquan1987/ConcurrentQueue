package concurrentqueue

import (
	"sync"
	"testing"
	"time"
)

func TestConcurrentQueue(t *testing.T) {
	queue := New()

	var wg sync.WaitGroup
	wg.Add(150)
	for i := 0; i < 100; i++ {
		go func(d int) {
			queue.Enqueue(d)
			wg.Done()
		}(i)

	}
	time.Sleep(time.Second * 2)
	for i := 0; i < 50; i++ {
		go func() {
			queue.Dequeue()
			wg.Done()
		}()
	}

	wg.Wait()
	if queue.Lenth() != 50 {
		t.Fatalf("fail")
	}
	t.Logf("ok")
}

func TestConcurrentQueue1(t *testing.T) {
	queue := New()

	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(d int) {
			queue.Enqueue(d)
			wg.Done()
		}(i)

	}

	wg.Wait()
	if queue.Lenth() != 100 {
		t.Fatalf("fail")
	}
	t.Logf("ok")
}

func TestConcurrentQueue2(t *testing.T) {
	queue := New()

	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 10; i++ {
		queue.Enqueue(i)
	}

	for i := 0; i < 5; i++ {
		go func() {
			queue.Dequeue()
			wg.Done()
		}()
	}
	wg.Wait()

	if queue.Lenth() != 5 {
		t.Error("fail")
	}
	v, ok := queue.Dequeue()
	if v == 5 && ok {
		t.Logf("ok")
	} else {
		t.Errorf("fail,v:%v", v)
	}

}
