package main
//import "time"
import "strings"
import "fmt"

func main() {
    id := strings.TrimPrefix("/account/register/", "/account/register/ru/")
    fmt.Println(id)

}