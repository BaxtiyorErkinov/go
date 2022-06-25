package main

import (
    "fmt"
)

var mess string = "World" // global scope

func main() {
    // var mess string = "Hello World" // local scope
    // mess := sayHi("Baxtiyor", 18)
    // printMess(mess)
    // printMess("Func 1")
    // printMess("Func 2")
    // printMess("Func 3")
    fmt.Println(findMax(1,2,3,4,5))
    fmt.Println(findMin(1,2,3,4,5))
}

// func printMess(str string) {
//     fmt.Println(str)
// }

// func sayHi(name string, age int) string {
//     return fmt.Sprintf("Hello %s, Age: %d years old", name, age )
// }

func findMin(numbers ...int) int {
    if len(numbers) == 0 {
        return 0
    }
    min := numbers[0]
    for _, i := range numbers {
        if i < min {
            min = i
        }
    }
    return min
}

func findMax(numbers ...int) int {
    if len(numbers) == 0 {
        return 0
    }
    max := numbers[0]

    for _, i := range numbers {
        if i > max {
            max = i
        }
    }
    return max
}