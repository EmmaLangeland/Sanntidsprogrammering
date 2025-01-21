Task 3 explenation

The result may end up being different from zero in the implementation.
The i++ and i-- operations are not seperate from each other.
Each operation involves multiple steps: 
- reading the value of i
- modifying it
- writing it back.
If two goroutines access i simultaneously, their operations may interfere in unexpected ways.
For example, i++ might read the value of i, but before it writes the incremented value, i-- might modify i, leading to lost updates.
Both incrementing and decrementing goroutines run concurrently without synchronization, which means they can interfere with each other.

GOMAXPROCS set the number of CPU threads. Using runtime.GOMAXPROCS to ensure a 1:1 mapping of goroutines to CPU cores (or threads). If set to 1, only one CPU will be used and the result will always end up as 0 due to no overlap. But it will work slow. In my implementation it is set to GOMAXPROCS(2) which increases the likelihood of race conditions because both goroutines might run on separate threads simultaneously.

	