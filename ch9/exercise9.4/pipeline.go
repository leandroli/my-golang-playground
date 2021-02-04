package exercise9_4

func Pipeline(stage int) (in, out chan int) {
	first := make(chan int)
	in = first
	for i := 0; i < stage; i++ {
		out = make(chan int)
		go func(out, in chan int) {
			for v := range in {
				out <- v
			}
			close(out)
		}(out, in)
		in = out
	}
	return first, out
}
