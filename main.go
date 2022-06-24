package main

import (
    "fmt"
    "reflect"
)

func main() {

    const mess string = "Hello World" // constant - не возможно изменять value

    var mess2 int = 1 // возможно изменять value 

    mess3 := "variable" // := означает что иницализатции value. A  один = означает присваивание
    mess3 = "new value"

    fmt.Println(mess, mess2, mess3)
    fmt.Println(reflect.TypeOf(mess2))
}