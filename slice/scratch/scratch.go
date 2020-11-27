package main
import "fmt"
import "strings"
import "io"

func main() {
    var myarray [5]int
    var mylitarray [5]int = [5]{1,2,3,4,5}
    var x := [...]int{1,2,3,4,5}

    for i, v range x {
        fmt.Printf("%d, %d", i, v)
    }
    myslice = make([]int, 10)  // length=10, cap=10
    myslice2 = make([]int, 10, 15) // length=10, cap=15

    var idMap map[string]int
    idMap = make(map[string]int) { "joe" : 123 }
    fmt.Println(idMap["jod"])
    delete(idMap["joe"])
    id, p := idMap["joe"] // ,id is the value, p is a bool for existence
    for key, val := range idMap {
        fmt.Println(key, val)
    }


    type struct Person {
        name string
        addr string
        phone string
    }
    var p1 Person
    p1.name = "Joe"
    p1.addr = "666 Road to Oblivion"
    p2 := new(Person)
    p3 := Person(name: "joe", addr: "a st.", phone: "123")

}
