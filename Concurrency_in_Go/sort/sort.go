package main
import "fmt"
import "strconv"

func bubbleSortBy4(items []int) []int {
    if items == nil {
        return(nil)
    }
    chunk_sizes := make([]int, 4, 4)
    base_size := len(items) / 4
    rem_size := len(items) % 4
    for i := 0; i < 4; i++ {
        chunk_sizes[i] = base_size
        if rem_size > 0 {
            chunk_sizes[i]++
            rem_size--
        }
    }
    channels := make([]chan []int, 0, 4)
    start := 0
    for _, chunk_size := range(chunk_sizes) {
        if chunk_size > 0 {
            end := start + chunk_size
            ch := make(chan []int)
            go internalBubbleSort(items[start:end], ch)
            start = start + chunk_size
            channels = append(channels, ch)
        }
    }
    sorted := make([]int, 0, 0)
    for _, ch := range(channels) {
        sorted_part := <-ch
        sorted = merge(sorted, sorted_part)
    }
    return(sorted)
}

func internalBubbleSort(items []int, ch chan []int) {
    fmt.Println("Sorting the following integers")
    for _, v := range(items) {
        fmt.Println(v)
    }
    sorted := make([]int, len(items), len(items))
    copy(sorted, items)
    for i := 0; i < len(sorted); i++ {
        for j := 0; j < len(sorted); j++ {
            if sorted[i] < sorted[j] {
                sorted[i], sorted[j] = sorted[j], sorted[i]
            }
        }
    }
    ch <- sorted
    return
}

func merge(items1 []int, items2 []int) []int {
    output := make([]int, len(items1)+len(items2), len(items1)+len(items2))
    idx1 := 0
    idx2 := 0
    for idx1 < len(items1) || idx2 < len(items2) {
        if idx1 >= len(items1) {
            output[idx1+idx2] = items2[idx2]
            idx2++
        } else if idx2 >= len(items2) {
            output[idx1+idx2] = items1[idx1]
            idx1++
        } else {
            if items1[idx1] < items2[idx2] {
                output[idx1+idx2] = items1[idx1]
                idx1++
            } else {
                output[idx1+idx2] = items2[idx2]
                idx2++
            }
        }
    }
    return(output)
}

func main() {
    input_list := make([]int, 0, 20)
    for ; ; {
        fmt.Println("Input an integer or press enter to sort and print.")
        var input_string string
        _, scanf_err := fmt.Scanf("%s", &input_string)
        if scanf_err != nil {
            break
        }
        conv_int, int_err := strconv.Atoi(input_string)
        if int_err != nil {
            fmt.Println(int_err)
            return
        }
        input_list = append(input_list, conv_int)
    }
    sorted_list := bubbleSortBy4(input_list)
    fmt.Println("Final sorted list:")
    for _, v := range(sorted_list) {
        fmt.Println(strconv.Itoa(v))
    }
    return
}

