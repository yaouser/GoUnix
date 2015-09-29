//golang闭包实例
import (
    "fmt"
)

func Func() (func(), func()) {
    i := 10
    return func() {
            i++
        }, func() {
            fmt.Println(i)
        }
}

func main() {
    Add, Print := Func()

    for i := 0; i < 10; i++ {
        Add()
    }

    Print()
}