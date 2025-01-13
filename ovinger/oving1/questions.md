Exercise 1 - Theory questions
-----------------------------

### Concepts

What is the difference between *concurrency* and *parallelism*?
> Concurrency focuses on doing different tasks by switching between them (almost at the same time) on one resource. Parallelism utilizes multiple resources to execute different tasks simultaneously. ex: Having one line and wait for turn = concurrency. Having two lines = parallelism.

What is the difference between a *race condition* and a *data race*? 
> A race condition or race hazard is the condition of systems where the system's substantive behavior is dependent on the sequence or timing of other uncontrollable events, leading to unexpected or inconsistent results. A data race occurs when 2 instructions from different threads access the same memory location, at least one of these accesses is a write and there is no synchronization that is mandating any particular order among these accesses.
A data race occurs when 2 instructions from different threads access the same memory location, at least one of these accesses is a write and there is no synchronization that is mandating any particular order among these accesses.
They are not a subset of one another, but are not the same. 
 
*Very* roughly - what does a *scheduler* do, and how does it do it?
> The scheduler ensures the execution of jobs at specific times or in response to specific triggering events.


### Engineering

Why would we use multiple threads? What kinds of problems do threads solve?
> Threads opens for faster programs because of abling concurrency and faster overall execution by using all the resources. Threads solve problems when it comes to multitasking. It can handle operations such as user input at the same time as autosaving documents and checking grammar. It also ables handling asynchronous events, that do not need to depend on each other. Such as the blinking lights in https://www.youtube.com/watch?v=k5xVDfu8BoQ&list=PLK2nnjsOlyODF2wrbnJUjDP3kNf-NG8Gb&index=5. 

Some languages support "fibers" (sometimes called "green threads") or "coroutines"? What are they, and why would we rather use them over threads? 
> Fibers, green threads, or coroutines are managed by the programming language/application itself, and not the OS like "usual" threads are. Coroutines are generalizations of subroutines (functions) that can pause their execution (yield) and resume later, maintaining their state. The programmer/code decides when to pause and switch fibers and not the OS. It is said to be lightweight because it doesnt involve the kernel. It is also typically limited to one thread resulting in less use of resources.

Does creating concurrent programs make the programmer's life easier? Harder? Maybe both?
> Maybe harder because of the complex task it is ti create synchronized systems. Understanding and debugging concurrent programs is harder because multiple threads or tasks execute independently and may interact in unexpected ways. Since the execution is not predictable, it can be harder to reproduce. 

What do you think is best - *shared variables* or *message passing*?
> I think it varies what is "best" depending on the problem that is to be solved. For small scaled problems shared variables seem to be the best. On the other hand message passing seems better because there is lower risk of race condtions because of no shared memory. It also works across different machines. Maybe combine them is best, if possible. 


