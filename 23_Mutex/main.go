package main

import (
	"fmt"
	"sync"
)

type post struct {
	mu    sync.Mutex
	views int
}

func (p *post) incrementViews(wg *sync.WaitGroup) {
	defer func() {
		p.mu.Unlock()
		wg.Done()
	}()

	p.mu.Lock()
	p.views += 1
}

func main() {

	var wg sync.WaitGroup
	myPost := post{views: 0}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go myPost.incrementViews(&wg)
	}
	wg.Wait()
	fmt.Println(myPost.views)
}
