package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	const N = 5

	var m sync.Map
	var wg sync.WaitGroup

	wg.Add(2 * N)

	// Писатели
	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			value := rand.Intn(100)
			m.Store(i, value)
			fmt.Printf("Writer %d wrote: m[%d] = %d\n", i, i, value)
		}()
	}

	// Читатели
	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			val, ok := m.Load(i)
			if ok {
				fmt.Printf("Reader %d read: m[%d] = %d\n", i, i, val.(int))
			} else {
				fmt.Printf("Reader %d: m[%d] not found\n", i, i)
			}
		}()
	}

	wg.Wait()

	fmt.Println("Finished")

	m.Range(func(key, value any) bool {
		fmt.Printf("m[%d] = %d\n", key.(int), value.(int))
		return true
	})
}
