// Write two goroutines which have a race condition when executed concurrently.
// Explain what the race condition is and how it can occur.

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func increase_value_by_1(param *int) {
	*param += 1
	wg.Done()
}

func multiply_value_by_2(param *int) {
	*param *= 2
	wg.Done()
}

func main() {
	thisVar := 12
	wg.Add(4)
	go increase_value_by_1(&thisVar)
	go multiply_value_by_2(&thisVar)
	go increase_value_by_1(&thisVar)
	go multiply_value_by_2(&thisVar)
	wg.Wait()
	fmt.Printf("Value after both the operations: %d", thisVar)
}

/*
RACE CONDITION: Occurs when two or more goroutines have access to the same data and access it simulataneously.
	This kind of behaviour, if not controlled, can create undeterministic outcomes for each execution.
	The differece would depend on all possible combinations of the interleavings possible.

I am demonstrating this behaviour using two functions, i.e. `increase_value_by_1` and `multiply_value_by_2`.
	and calling each of these functions twice, just to amplify the undeterministic bahavior of the race conditions
	by increasing the number of interleaving combinations.

Both the functions share access to the same underlying variable (pointer). Each functions alters the value of the
	variable, as per the definition, in the order determined at runtime. This order can differ for each execution of the
	code.

When I ran this code for first 5 times, I got:
- "54" 3 times: (((12 + 1) * 2) + 1) * 2
- "51" one time: (((12 * 2) + 1) * 2) + 1
- "53" one time. (((12 + 1) * 2) * 2) + 1

There can be other combination possible, as interleaving are generated on machine code level, not source code level.

*/
