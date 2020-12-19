package main
import "fmt"
import "strings"
import "io"

func main() {
    var err error = nil
    var complete_string string = ""
    var partial string
    fmt.Println("Enter text.  End your text with a blank line to find 'ian'.")
    for ; err == nil; {
        _, err = fmt.Scanf("%s",&partial)
        if err == io.EOF { break; }
        complete_string = complete_string + partial
    }
    complete_string = strings.ToLower(complete_string)
    if !strings.HasPrefix(complete_string, "i") { print("Not Found!\n"); return }
    if !strings.ContainsRune(complete_string, 'a') { print("Not Found!\n"); return }
    if !strings.HasSuffix(complete_string, "n") { print("Not Found!\n"); return }
    print("Found!\n")
    return
}
