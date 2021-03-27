package exercise9_4_test

import (
	exercise9_4 "github.com/leandroli/my-golang-playground/ch9/exercise9.4"
	"testing"
	"time"
)

func BenchmarkPipeline(b *testing.B) {
	i := 1000000
	start := time.Now()
	in, out := exercise9_4.Pipeline(i)
	in <- i
	<-out
	b.Log(time.Since(start))
}
