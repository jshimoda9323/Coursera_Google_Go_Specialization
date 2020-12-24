package main
import "fmt"
import "sync"

var x int
var wg sync.WaitGroup
func set_x() {
    x = 1
    wg.Done()
}
func inc_x() {
    x = x + 2
    wg.Done()
}
/*
  The race condition in this case is the modification to "x" in set_x and inc_x functions.
  The race condition occurs because the order of what modifies x is undefined.  In other words,
  the order of execution between concurrent tasks set_x and inc_x is unknown.
*/
func main() {
    x = 0
    wg.Add(2)
    go set_x()
    go inc_x()
    wg.Wait()
    fmt.Println(x)
}
