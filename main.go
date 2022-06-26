package main

import (
    "fmt"
)

func main() {
    // ================= ARRAYS ===================
    var arr = [3]int{1,2}
   // arr = append(arr, 2) // Error => first argument to append must be slice; have [3]int

    fmt.Println(arr) // Output: [ 1, 2, 0]

    // =============== SLICES ===============

    var seaFishes = []string{"shark", "delphin", "squid", "akula", "sakkizoyoq", "toshbaqa"}
    // fmt.Printf("%q\n", seaFishes)

    // var names = make([]string, 3, 5)
    // names = append(names, "salom", "salom", "salom")
    // fmt.Printf("%q\n", names    )
    var other = []string{}
    // other = append(seaFishes[:2], seaFishes[3:]...) // seaFishes[:2] - 2gacha bulgan yana 0,1 indexdagi elmentlar. SeaFisher[3:] 3 - index dan bowlab uyogini oladi
    other = append(seaFishes[:3], seaFishes[4:]...)
    fmt.Println(other)



}
 