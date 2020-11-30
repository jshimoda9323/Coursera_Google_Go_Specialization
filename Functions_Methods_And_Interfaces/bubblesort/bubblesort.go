package main

import "io"
import "fmt"
import "strconv"

func main() {
    fmt.Println("Enter one to ten integers:")
    var num int
    var err error = nil
    var input string
    input_integers := make([]int, 0, 10)
    for ; err == nil && err != io.EOF && len(input_integers) < 10; {
        num, err = fmt.Scanf("%s",&input)
        if num > 0 {
            i, e := strconv.Atoi(input)
            if e == nil {
                input_integers = append(input_integers, i)
            } else {
                fmt.Println("Error reading integer. Enter again.")
            }
        }
    }
    BubbleSort(input_integers)
    fmt.Println(input_integers)
}

func BubbleSort(inSli []int) {
    for idx1, val1 := range(inSli) {
        for idx2, val2 := range(inSli) {
            if val1 < val2 {
                Swap(&inSli[idx1], &inSli[idx2])
            }
        }
    }
}

func Swap(a, b *int) {
    *a, *b = *b, *a
}
