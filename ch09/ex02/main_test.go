package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestInitPCOnce(t *testing.T) {
	done := make(chan struct{})

	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			<-done
			count := PopCount(0x12345678901234)
			if count != 20 {
				panic(fmt.Sprintf("count is %d, want 20\n", count))
			}
			wg.Done()
		}()
	}
	close(done)
	wg.Wait()
}