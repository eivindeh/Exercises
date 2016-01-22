// Go 1.2
// go run helloworld_go.go

package main

import (
    . "fmt"
    "runtime"
    "time"
)
var i int = 0

func add() {
    for j :=0; j<1000000;j++ {
    i++
    }
}
func sub() {
    for j :=0; j<1000000;j++ {
    i--
    }
}

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())    // I guess this is a hint to what GOMAXPROCS does...
                                            // Try doing the exercise both with and without it!
    go add()
    go sub()		                    

    time.Sleep(100*time.Millisecond)
    Println("%d\n",i)
}
