package main

import (
    "fmt"
    "errors"
)


func main() {
    mess, allowed := enterTheClub(6)
    fmt.Println(mess, allowed)

}

func enterTheClub(age int) (string, error) {
    if age >= 18 && age < 45 {
        return "Coming", nil
    } else if age >= 45 && age < 65 {
        return "Are you sure?", nil
    } else if age >= 65 {
        return "You are very old", errors.New("You are very old!!!")
    }
    return "Closed", errors.New("You are very young")
}