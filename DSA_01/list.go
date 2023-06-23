package main 

import (
    "fmt"
    "container/list"
)

func main() {
    var intList list.List
    intList.PushBack(69)
    intList.PushBack(420)
    intList.PushBack(42069)

    for element := intList.Front(); element != nil; element = element.Next() {
        fmt.Println(element.Value.(int))
    }
}
