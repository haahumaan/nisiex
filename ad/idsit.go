package main

import (
	"fmt"
	"sync"
)

type MyStruct struct {
	// Define the fields of MyStruct here
}

func main() {
	pool := &sync.Pool{
		New: func() interface{} {
			return &MyStruct{} // Create a new instance of MyStruct
		},
	}

	// Get an instance from the pool
	instance := pool.Get().(*MyStruct)

	// Use the instance
	fmt.Println(instance)

	// Put the instance back into the pool
	pool.Put(instance)

	// Get the instance again (reused from the pool)
	reusedInstance := pool.Get().(*MyStruct)

	// Verify that the instance is the same
	fmt.Println(reusedInstance == instance) // Output: true
}
