package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			// たいてい"i : 5"が出力される
			fmt.Printf("i : %d\n", i)
			wg.Done()
		}()
	}
	wg.Wait()
	for j := 0; j < 5; j++ {
		wg.Add(1)
		go func(j int) {
			// 0から4
			fmt.Printf("j: %d\n", j)
			wg.Done()
		}(j)
	}
	wg.Wait()
	for k := 0; k < 5; k++ {
		k := k
		wg.Add(1)
		go func() {
			fmt.Printf("k : %d\n", k)
			wg.Done()
		}()
	}
	wg.Wait()
}
