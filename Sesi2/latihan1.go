package main

import (
	"fmt"
	"sync"
)

func print(key string , data string) {
	fmt.Println(key , data)
}

func main() {
	wg := sync.WaitGroup{}
	arr := map[string]string{
		"Nama":    "NooBee",
		"Class":   "Backend Intermediate",
		"Address": "Jakarta",
	}

	for key, data := range arr {
		go func(key string , data string) {
			wg.Add(1)
			print(key , data)
			wg.Done()
		}(key , data)
	}

	wg.Wait()
}