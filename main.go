package main

import (
    "fmt"
)

var mess string = "World" // global scope

func main() {
    // var mess string = "Hello World" // local scope
    mess := sayHi("Baxtiyor", 18)
    printMess(mess)
    // printMess("Func 1")
    // printMess("Func 2")
    // printMess("Func 3")
}

func printMess(str string) {
    fmt.Println(str)
}

func sayHi(name string, age int) string {
    return fmt.Sprintf("Hello %s, Age: %d years old", name, age )
}