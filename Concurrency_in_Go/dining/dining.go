package main
import "fmt"
import "strconv"
import "sync"
import "container/list"

type Chopstick struct { sync.Mutex }

type Philosopher struct {
    id int
    leftChopstick, rightChopstick *Chopstick
}

// Pass iteration_count < 1 for an infinite loop
func (phil *Philosopher) Eat(iteration_count int, can_eat chan bool, ready chan int, done chan int) {
    for i := 0;; {
        if iteration_count > 0 {
            if i >= iteration_count {
                break
            } else {
                i++
            }
        }
        // Indicate to host we are ready to eat
        ready <- phil.id
        // Wait for permission from host
        <-can_eat
        // Obtain locks
        phil.leftChopstick.Lock()
        phil.rightChopstick.Lock()
        // Eat
        fmt.Println("starting to eat "+strconv.Itoa(phil.id+1))
        fmt.Println("finishing eating "+strconv.Itoa(phil.id+1))
        // Free locks
        phil.leftChopstick.Unlock()
        phil.rightChopstick.Unlock()
    }
    done <- phil.id
}

func main() {

    // Program settings
    chopstick_count := 5       // Also determines number of philosophers.
    simultaneous_eaters := 2   // Number of philosophers allowed to eat at the same time.
    times_to_eat := 3          // Number of times a single philosopher will eat.  Set to 0 for infinite.

    // Declare chopsticks and philosophers
    chopsticks := make([]*Chopstick, chopstick_count)
    philosophers := make([]*Philosopher, len(chopsticks))
    // Keep track of who is eating
    is_eating := make([]bool, len(philosophers))
    // Assign chopsticks to philosophers, and things.
    for i := 0; i < len(chopsticks); i++ {
        chopsticks[i] = new(Chopstick)
    }
    for i := 0; i < len(philosophers); i++ {
        philosophers[i] = new(Philosopher)
        is_eating[i] = false
        philosophers[i].id = i
        philosophers[i].leftChopstick = chopsticks[i]
        philosophers[i].rightChopstick = chopsticks[(i+1)%len(chopsticks)]
    }

    // One common channel for receiving ready signal from philosophers.
    ready := make(chan int)
    // A channel per philosopher for sending eat signal to philosophers.
    can_eat := make([]chan bool, len(philosophers))
    // One common channel for receiving done signal from philosophers.
    done := make (chan int)
    // A queue of ids who are ready to eat.
    ready_queue := list.New()

    // The philosophers attempt to start eating
    for i, phil := range(philosophers) {
        can_eat[i] = make(chan bool)
        go phil.Eat(times_to_eat, can_eat[i], ready, done)
    }

    // Iterate until all philosophers are done (may be forever)
    available_threads := simultaneous_eaters
    num_done := 0
    dispatch_loop: for {
        select {
        case ready_id := <-ready:
            // Push this id onto the ready queue.
            ready_queue.PushBack(ready_id)
            if is_eating[ready_id] {
                is_eating[ready_id] = false
                available_threads++
            }
            // Dispatch the next available id (may not necessarily be the first in the queue)
            dispatch(&available_threads, is_eating, can_eat, ready_queue)
        case v := <-done:
            if is_eating[v] {
                is_eating[v] = false
                available_threads++
            }
            // Dispatch the next available id (may not necessarily be the first in the queue)
            dispatch(&available_threads, is_eating, can_eat, ready_queue)
            num_done++
            if num_done >= len(philosophers) { break dispatch_loop }
        }
    }
    return
}

func dispatch(available_threads *int, is_eating []bool, can_eat []chan bool, ready_queue *list.List) {
    if *available_threads > 0 {
        // Find the next appropriate philosopher that can eat.  Cannot be adjacent to current eaters.
        for e := ready_queue.Front(); e != nil; e = e.Next() {
            e_idx := e.Value.(int)
            l_idx := e_idx - 1
            if l_idx < 0 { l_idx = len(is_eating) - 1 }
            r_idx := (e_idx+1) % len(is_eating)
            if !is_eating[l_idx] &&
                    !is_eating[r_idx] {
                next_id := e_idx
                ready_queue.Remove(e)
                is_eating[e_idx] = true
                *available_threads--
                can_eat[next_id] <- true
                return
            }
        }
    }
}

