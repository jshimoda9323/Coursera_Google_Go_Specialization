package main
import "fmt"

type Animal interface {
    Eat()
    Move()
    Speak()
}

type Cow struct { }
func (cow Cow) Eat() {
    fmt.Println("grass")
}
func (cow Cow) Move() {
    fmt.Println("walk")
}
func (cow Cow) Speak() {
    fmt.Println("moo")
}

type Bird struct { }
func (bird Bird) Eat() {
    fmt.Println("worms")
}
func (bird Bird) Move() {
    fmt.Println("fly")
}
func (bird Bird) Speak() {
    fmt.Println("peep")
}

type Snake struct { }
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
    fmt.Println("The following commands are available:")
    fmt.Println("newanimal <name> [cow|bird|snake]")
    fmt.Println("query <name> [eat|move|speak]")
    input_loop: for ; ; {
        fmt.Print(">")
        num, scanf_err := fmt.Scanf("%s %s %s", &input_strings[0], &input_strings[1], &input_strings[2])
        if scanf_err != nil || num != 3 {
            fmt.Println("Invalid input.")
            fmt.Println("Usage: animalcli newanimal <name> <type>")
            fmt.Println("       query <name> <function>")
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
                animals[input_strings[1]] = Cow { }
            case "bird":
                animals[input_strings[1]] = Bird { }
            case "snake":
                animals[input_strings[1]] = Snake { }
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
