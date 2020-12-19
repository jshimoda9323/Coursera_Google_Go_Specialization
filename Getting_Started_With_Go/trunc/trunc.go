package main
import "fmt"
import "strconv"

func main() {
    var x float32
    fmt.Println("Input a floating point number:")
    for _, err := fmt.Scanf("%f", &x); err == nil; _, err = fmt.Scanf("%f", &x) {
        if err == nil {
            fmt.Println("The integer portion is:")
            fmt.Println(strconv.Itoa(int(x)))
        } else {
            return
        }
        fmt.Println("Input another floating point number:")
    }
}
