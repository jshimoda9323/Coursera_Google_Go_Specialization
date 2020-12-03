package main

import "io"
import "fmt"
import "strconv"
import "math"

func main() {

    fmt.Println("Enter space separated values for accelearation, initial velocity and initial displacement:")
    var num int
    var err error = nil
    var input string
    input_numbers := make([]float64, 0, 3)
    for ; err == nil && err != io.EOF && len(input_numbers) < 10; {
        num, err = fmt.Scanf("%s",&input)
        if num > 0 {
            f, e := strconv.ParseFloat(input, 64)
            if e == nil {
                input_numbers = append(input_numbers, f)
            } else {
                fmt.Println("Error reading number!")
                return
            }
        }
    }
    if len(input_numbers) != 3 {
        fmt.Println("Must input 3 space separated numbers.")
        return
    }

    fn := GenDisplaceFn(input_numbers[0], input_numbers[1], input_numbers[2])

    num = 1
    for ; num == 1; {
        fmt.Println("Enter duration:")
        num, err = fmt.Scanf("%s",&input)
        if num > 0 {
            t, e1 := strconv.ParseFloat(input, 64)
            if e1 == nil {
                fmt.Println("Computed displacement: " + strconv.FormatFloat(fn(t), 'E', -1, 64))
            } else {
                fmt.Println("Error reading number!")
                return
            }
        }
    }
    return
}

func GenDisplaceFn(a, vo, so float64) func(float64) float64 {
    return func(t float64) float64 {
        return((0.5*a*math.Pow(t,2))+(vo*t)+so)
    }
}


