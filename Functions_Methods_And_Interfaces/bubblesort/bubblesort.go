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
    for i := 0; i < len(inSli); i++ {
        for j := 0; j < len(inSli)-1-i; j++ {
            if inSli[j] > inSli[j+1] {
                Swap(inSli, j)
            }
        }
    }
}

func Swap(sl []int, i int) {
    sl[i], sl[i+1] = sl[i+1], sl[i]
}
