package main

import (
	"fmt"
)

func sub() {
    for {
        fmt.Println("sub loop")
    }
}

func main() {
    go sub()
    for {
        fmt.Println("main loop")
    }
}
