package main

import (
	"Task_3/internal/cache"
	"fmt"
	"time"
)

func main() {
	c, err := cache.CreateCache(3)
	if err != nil {
		fmt.Println("Error creating cache:", err)
		return
	}

	go c.Scan()

	c.Set("a", 100)
	c.Set("b", "hello")
	c.Set("c", []int{1, 2, 3})

	fmt.Println("Initial cache:")
	fmt.Println("a:", c.Get("a"))
	fmt.Println("b:", c.Get("b"))
	fmt.Println("c:", c.Get("c"))

	removed := c.Remove("b")
	fmt.Println("Remove 'b':", removed)
	fmt.Println("b:", c.Get("b"))

	time.Sleep(time.Second * 11)
	if c.Get("a") != nil {
		fmt.Println("Element 'a' is still in the cache")
	} else {
		fmt.Println("Element 'a' has been removed by scan")
	}
}