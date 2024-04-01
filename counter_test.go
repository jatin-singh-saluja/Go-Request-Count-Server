package main

import (
	"sync/atomic"
	"testing"
)

func TestIncrementVal(t *testing.T) {

	counter := &requestCounter{
		val: 0,
	}

	counter.incrementVal(1)
	value := int(atomic.LoadInt64(&counter.val))

	if value != 1 {
		t.Errorf("Expected value of request counter as 1, but got %v", value)
	}

}
