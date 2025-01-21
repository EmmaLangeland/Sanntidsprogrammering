// Compile with `gcc foo.c -Wall -std=gnu99 -lpthread`, or use the makefile
// The executable will be named `foo` if you use the makefile, or `a.out` if you use gcc directly

#include <pthread.h>
#include <stdio.h>

int i = 0;

pthread_mutex_t mutex; //mutex ensures that only one thread can modify the shared variable i at a time, a binary lock, preventing race conditions and guaranteeing a consistent final result. A semaphore is more general. multiple threads can access i simultaneously up to the limit defined. 
//mutex seems better for a simple lock

/*
// Note the return type: void*
void* incrementingThreadFunction(){
    // TODO: increment i 1_000_000 times
    for (int j= 0; j < 1000000; j++){
        i++;
    }
    return NULL;
}
*/

void* incrementingThreadFunction(){
    // TODO: increment i 1_000_000 times
    for (int j= 0; j < 1000000; j++){
        //- Acquire the lock, do your work in the critical section, and release the lock.
        pthread_mutex_lock(&mutex);
        i++;
        pthread_mutex_unlock(&mutex);
        
    }
    return NULL;
}
/*
void* decrementingThreadFunction(){
    // TODO: decrement i 1_000_000 times
    for (int j = 0; j <1000000; j++){
        i--;
    }
    return NULL;
}
*/

void* decrementingThreadFunction(){
    // TODO: decrement i 1_000_000 times
    for (int j = 0; j <1000000; j++){
        pthread_mutex_lock(&mutex);
        i--;
        pthread_mutex_unlock(&mutex);
    }
    return NULL;
}


int main(){
    
    // TODO: 
    
    // start the two functions as their own threads using `pthread_create`
    // Hint: search the web! Maybe try "pthread_create example"?
    
    // create the thread objs
    pthread_t thread1, thread2;
    // Initialize the mutex before use
    pthread_mutex_init(&mutex, NULL);
    // start the threads
    pthread_create(&thread1, NULL, &incrementingThreadFunction, NULL);
    pthread_create(&thread2, NULL, &decrementingThreadFunction, NULL);
    
    // TODO:
    // wait for the two threads to be done before printing the final result
    // Hint: Use `pthread_join`   
    pthread_join(thread1,NULL);
    pthread_join(thread2,NULL);

    pthread_mutex_destroy(&mutex); 
    
    printf("C The magic number is: %d\n", i);
    return 0;
}
