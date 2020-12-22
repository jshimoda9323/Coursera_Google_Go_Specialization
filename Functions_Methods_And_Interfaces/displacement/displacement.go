package main

import "fmt"
import "strconv"
import "math"

func main() {

    fmt.Println("Enter space separated values for accelearation, initial velocity and initial displacement followed by a blank line:")
    var input [3]string
    var input_numbers [3]float64
    num, err := fmt.Scanln(&input[0], &input[1], &input[2])
    if err != nil || num != 3 {
        if err != nil { fmt.Println(err) }
        fmt.Println("Error: Must input three numbers")
        return
    }
    f1, e1 := strconv.ParseFloat(input[0], 64)
    f2, e2 := strconv.ParseFloat(input[1], 64)
    f3, e3 := strconv.ParseFloat(input[2], 64)
    if e1 != nil || e2 != nil || e3 != nil {
        fmt.Println("Error parsing floating point number.")
        return
    }
    input_numbers[0] = f1
    input_numbers[1] = f2
    input_numbers[2] = f3

    fn := GenDisplaceFn(input_numbers[0], input_numbers[1], input_numbers[2])

    num = 1
    for {
        fmt.Println("Enter duration:")
        num, err := fmt.Scanln(&input[0])
        if err != nil || num != 1 {
            if err != nil { fmt.Println(err) }
            fmt.Println("Error: Must enter a single floating point number.")
            return
        }
        t, e4 := strconv.ParseFloat(input[0], 64)
        if e4 != nil {
            fmt.Println("Error parsing floating point number.")
            return
        }
        fmt.Println("Computed displacement: " + strconv.FormatFloat(fn(t), 'E', -1, 64))
    }
    return
}

func GenDisplaceFn(a, vo, so float64) func(float64) float64 {
    return func(t float64) float64 {
        return((0.5*a*math.Pow(t,2))+(vo*t)+so)
    }
}


