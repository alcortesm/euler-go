package e0002

// Solution returns the solution to problem 2
func Solution() int {
	return solution(4000000)
}

func solution(ceil int) int {
	return sum(keepEven(fib(ceil)))
}

func fib(ceil int) <-chan int {
	ch := make(chan int)
	go func() {
		for a, b := 1, 1; b < ceil; a, b = b, a+b {
			ch <- b
		}
		close(ch)
	}()
	return ch
}

func keepEven(input <-chan int) <-chan int {
	output := make(chan int)
	go func() {
		for n := range input {
			if n%2 == 0 {
				output <- n
			}
		}
		close(output)
	}()
	return output
}

func sum(input <-chan int) int {
	sum := 0
	for n := range input {
		sum += n
	}
	return sum
}
