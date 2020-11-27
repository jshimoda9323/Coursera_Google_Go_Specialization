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
            fname_size := 20
            if len(fname) < 20 { fname_size = len(fname) }
            lname_size := 20
            if len(lname) < 20 { lname_size = len(lname) }
            persons = append(persons, Person{fname: fname[:fname_size], lname: lname[:lname_size]})
        }
    }

    for _, perp := range(persons) {
        fmt.Println(perp.fname + " " + perp.lname)
    }

    return
}
