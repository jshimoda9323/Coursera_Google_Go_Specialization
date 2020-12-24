package main
import "fmt"
//import "io"
//import "strconv"

type Animal struct {
    food string
    locomotion string
    noise string
}

func (animal *Animal) Eat() {
    fmt.Println(animal.food)
}

func (animal *Animal) Move() {
    fmt.Println(animal.locomotion)
}

func (animal *Animal) Speak() {
    fmt.Println(animal.noise)
}

type MyInt int
func main() {
    //v := MyInt(3)
    //fmt.Println(v.Double())
    //fmt.Println(v.Show())
    //v.DoubleMe()
    //fmt.Println(v.Show())

    cow := Animal{
        food: "The cow eats grass.",
        locomotion: "The cow struts along.",
        noise: "The cow moos!",
    }

    bird := Animal{
        food: "The bird eats worms.",
        locomotion: "The bird soars high.",
        noise: "The bird squawks!",
    }

    snake := Animal{
        food: "The snake eats rodents.",
        locomotion: "The snake slithers along.",
        noise: "The snake hisssssses.",
    }

    //var input1 string
    //var input2 string
    input_strings := make([]string, 2, 2)

    input_loop: for ; ; {
        var animal Animal
        fmt.Print("Enter an animal (cow, bird, or snake), followed by a space, followed by an action (eat, move, or speak)> ")
        num, scanf_err := fmt.Scanf("%s %s", &input_strings[0], &input_strings[1])
        if scanf_err != nil || num != 2 {
            fmt.Println("Error reading input.")
            break
        }
        //input_strings = append(input_strings, input1)
        //input_strings = append(input_strings, input2)
        switch input_strings[0] {
        case "cow":
            animal = cow
        case "bird":
            animal = bird
        case "snake":
            animal = snake
        default:
            fmt.Println("Unknown animal.")
            break input_loop
        }
        switch input_strings[1] {
        case "eat":
            animal.Eat()
        case "move":
            animal.Move()
        case "speak":
            animal.Speak()
        default:
            fmt.Println("Unknown action.")
            break input_loop
        }
        input_strings = make([]string, 2, 2)
    }
    return
}

