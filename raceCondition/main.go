package main

import (
	"fmt"
	"time"
)

var counter = 0

func increment() {
	for i := 0; i < 1000; i++ {
		value := counter
		time.Sleep(time.Millisecond) // Simulate some work
		counter = value + 1
	}
}

func main() {
	// Goroutine 1
	go increment()

	// Goroutine 2
	go increment()

	// Allowing some time for goroutines to execute
	time.Sleep(2 * time.Second)

	fmt.Println("Result with race condition:", counter)
}

/*
In the provided code, a race condition occurs due to the concurrent, unsynchronized access and modification of the shared variable counter by two goroutines. Let's break down how the race condition happens:

Goroutine Execution:

The main function launches two goroutines concurrently, both executing the increment function.
The increment function reads the current value of counter, performs some simulated work (time.Sleep), and then increments the value.


Unsynchronized Access:

Since there is no synchronization mechanism, both goroutines can access and modify the counter variable simultaneously.
The value := counter line reads the current value of counter in both goroutines.


Interleaved Execution:

The execution of the two goroutines is interleaved due to concurrent execution.
For instance, let's say Goroutine 1 reads counter and Goroutine 2 also reads counter before either of them updates it.


Overlapping Modifications:

If Goroutine 1 and Goroutine 2 both read the same value of counter before updating, they might both increment it and then store the updated value independently.


Lost Update:

The final value of counter becomes unpredictable as the increments from both goroutines might not be properly accounted for. One goroutine's increment might be overwritten by the other, leading to lost updates.


Race Condition:

The race condition occurs because the final value of counter depends on the interleaved and overlapping execution of the two goroutines, and it is not guaranteed to be the sum of their increments.


In summary, the absence of synchronization mechanisms allows the two goroutines to concurrently read and modify the shared variable, resulting in a race condition where the final value of counter is unpredictable and varies between different runs of the program.

*/
