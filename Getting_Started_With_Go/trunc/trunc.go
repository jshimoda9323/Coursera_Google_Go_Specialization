package main
import "fmt"
import "strconv"

func main() {
  var x float32
  for _, err := fmt.Scan(&x); err == nil; _, err = fmt.Scan(&x) {
    i := int(x)
    print(strconv.Itoa(i)+"\n")
  }
}
