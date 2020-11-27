package main

import "io"
import "fmt"
import "strconv"
import "sort"

func main() {
    var intarray []int = make([]int, 0, 3)
    var err error = nil
    var input_string string
    for ; err == nil; {
        _, err = fmt.Scanf("%s",&input_string)
        if err == io.EOF || input_string == "X" { break; }
        var input_integer int
        input_integer, err = strconv.Atoi(input_string)
        if err != nil { break; }
        intarray = append(intarray, input_integer)
        sort.Ints(intarray)
        fmt.Println(intarray)
    }
    return
}
