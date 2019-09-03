package factory

import (
	"fmt"
	"sync"
)

const (
	BUFFERS = 1000
)

func genPoolChanBuffer(nums []float64) <-chan interface{} {
	out := make(chan interface{}, BUFFERS)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}


type Handler func(w *Work)

type Work struct{
	input interface{}
}

type Pool struct{
	channum 	int
	workernum	int
	wg			*sync.WaitGroup
	ch			chan Work
	Func 		Handler
}

func (p Pool) RunWorker() {
	for i := 0; i < p.workernum; i++ {
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()
			for work := range p.ch {
				p.Func(&work)
			}
		}()
	}
}

func (p Pool) FeedWorker(in <-chan interface{}) {
	go func() {
		for n := range in {
			work := Work{
				input: n,
			}
			p.ch <- work
		}
	}()
}

func (p Pool) Wait() {
	p.wg.Wait()
}

func InitPool(channum, workernum int, f Handler) *Pool {
	return &Pool{
		channum:	channum,
		workernum:	workernum,
		wg:			&sync.WaitGroup{},
		ch:			make(chan Work, channum),
		Func:		f,
	}
}

// 测试用例
func test() {
	out := genPoolChanBuffer([]float64{1, 2, 4})

	p := InitPool(3, 1, func(w *Work) {
		fmt.Println(w.input)
	})

	p.RunWorker()
	p.FeedWorker(out)
	p.Wait()
}
