package main

import (
	"testing"
)

func BenchmarkXxx(b *testing.B) {
	router := New()

	for i := 0; i < b.N; i++ {
		router.MyHandler("/user/:id", Req) //1 allocs/op
	}
}
