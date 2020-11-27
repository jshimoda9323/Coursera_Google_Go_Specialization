package main

import "io/ioutil"
import "fmt"

func main() {

    dat, e := ioutil.ReadFile("input_file")
    if e != nil { fmt.Println(e); return }
    var lines []string = make([]string, 0, 10)
    var curline string = ""
    for _, val := range(dat) {
        char := rune(val)
        if char == '\n' {
            lines = append(lines, curline)
            curline = ""
        } else {
            curline = curline + string(char)
        }
    }

    type Person struct {
        fname string
        lname string
    }

    var persons []Person = make([]Person, 0, 10)
    for _, line := range(lines) {
        var fname string
        var lname string
        num, e := fmt.Sscanf(line, "%s %s", &fname, &lname)
        if e != nil { fmt.Println(e); return }
        if num > 0 {
            persons = append(persons, Person{fname: fmt.Sprintf("%-20.20s", fname), lname: fmt.Sprintf("%-20.20s", lname)})
        }
    }

    for _, peep := range(persons) {
        fmt.Println(peep.fname + " " + peep.lname)
    }

    return
}
