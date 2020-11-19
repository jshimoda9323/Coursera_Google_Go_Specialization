package main
import "fmt"
func exp() {
    fmt.Printf("Hello World :) \n");
    var x int = 1
    var y int
    var ip *int
    ip = &x
    y = *ip
    ptr := new(int)
    fmt.Scan(&x)
}
func foo() *int {
    var x int = 1
    return(&x)
}
