package cache_test

import (
	"Task_3/internal/cache"
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	c, err := cache.CreateCache(2)
	if err != nil {
		t.Fatalf("failed to create cache: %v", err)
	}

	c.Set("x", 100)
	c.Set("y", "hello")

	if got := c.Get("x"); got != 100 {
		t.Errorf("Get(x) = %v; want 100", got)
	}
	if got := c.Get("y"); got != "hello" {
		t.Errorf("Get(y) = %v; want 'hello'", got)
	}

	if removed := c.Remove("y"); !removed {
		t.Errorf("Remove(y) = false; want true")
	}
	if got := c.Get("y"); got != nil {
		t.Errorf("Get(y) after remove = %v; want nil", got)
	}

	go c.Scan()
	c.Set("a", 42)
	time.Sleep(11 * time.Second)
	if got := c.Get("a"); got != nil {
		t.Errorf("Get(a) after Scan = %v; want nil", got)
	}
}
