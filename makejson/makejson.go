package main

import "io"
import "fmt"
import "encoding/json"
import "strings"

func main() {
    var person map[string]string = make(map[string]string)
    var name_part string
    var name_parts []string = make([]string, 0, 1)
    var address_part string
    var address_parts []string = make([]string, 0, 1)
    var err error = nil
    var barr []byte
    var num int

    fmt.Println("Enter a name:")
    for ; err == nil && err != io.EOF; {
        num, err = fmt.Scanf("%s",&name_part)
        if num > 0 {
            name_parts = append(name_parts, name_part)
        }
    }
    person["name"] = strings.Join(name_parts, " ")
    fmt.Println("Enter an address:")
    err = nil
    for ; err == nil && err != io.EOF; {
        num, err = fmt.Scanf("%s",&address_part)
        if num > 0 {
            address_parts = append(address_parts, address_part)
        }
    }
    person["address"] = strings.Join(address_parts, " ")
    barr, err = json.Marshal(person)
    fmt.Println(string(barr))
    return
}
