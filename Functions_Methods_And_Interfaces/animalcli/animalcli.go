package main
import "fmt"

type Animal interface {
    Eat()
    Move()
    Speak()
}

type Cow struct { name string }
func (cow Cow) Eat() {
    fmt.Println("grass")
}
func (cow Cow) Move() {
    fmt.Println("walk")
}
func (cow Cow) Speak() {
    fmt.Println("moo")
}

type Bird struct { name string }
func (bird Bird) Eat() {
    fmt.Println("worms")
}
func (bird Bird) Move() {
    fmt.Println("fly")
}
func (bird Bird) Speak() {
    fmt.Println("peep")
}

type Snake struct { name string }
func (snake Snake) Eat() {
    fmt.Println("mice")
}
func (snake Snake) Move() {
    fmt.Println("slither")
}
func (snake Snake) Speak() {
    fmt.Println("hsss")
}

func main() {
    input_strings := make([]string, 3, 3)
    animals := make(map[string]Animal)
    input_loop: for ; ; {
        fmt.Print(">")
        num, scanf_err := fmt.Scanf("%s %s %s", &input_strings[0], &input_strings[1], &input_strings[2])
        if scanf_err != nil || num != 3 {
            fmt.Println("Invalid input.")
            continue input_loop
        }
        switch input_strings[0] {
        case "newanimal":
            if _, animal_exists := animals[input_strings[1]]; animal_exists {
                fmt.Println("Animal '"+input_strings[1]+"' already exists!")
                continue input_loop
            }
            switch input_strings[2] {
            case "cow":
                animals[input_strings[1]] = Cow { name: input_strings[1] }
            case "bird":
                animals[input_strings[1]] = Bird { name: input_strings[1] }
            case "snake":
                animals[input_strings[1]] = Snake { name: input_strings[1] }
            default:
                fmt.Println("Invalid animal type '"+input_strings[2]+"'")
            }
            fmt.Println("Created It!")
        case "query":
            if animal, ok := animals[input_strings[1]]; ok {
                switch input_strings[2] {
                case "eat":
                    animal.Eat()
                case "move":
                    animal.Move()
                case "speak":
                    animal.Speak()
                default:
                    fmt.Println("Invalid animal property.")
                }
            } else {
                fmt.Println("Animal '"+input_strings[1]+"' does not exist.")
            }
        default:
            fmt.Println("Unknown command '"+input_strings[0]+"'")
            continue input_loop
        }
    }
    return
}
