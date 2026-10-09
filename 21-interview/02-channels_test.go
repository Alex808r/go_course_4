package interview

import (
	"testing"
)

func TestSelect(t *testing.T) {
	Select()
}

func TestProduceConsume(t *testing.T) {
	ProduceConsume()
}

func TestFanIn(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		defer close(ch1)
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
	}()
	go func() {
		defer close(ch2)
		for i := 11; i < 20; i++ {
			ch2 <- i
		}
	}()

	ch := FanIn(ch1, ch2)
	count := 0
	for val := range ch {
		count++
		t.Log(val)
	}
	if count != 19 {
		t.Fatalf("expected 19 values from fan-in, got %d", count)
	}
}
