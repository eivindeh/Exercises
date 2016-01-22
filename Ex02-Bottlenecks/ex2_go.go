// Go 1.2
// go run helloworld_go.go

package main

import (
    . "fmt"
    "runtime"
    "time"
)
var i int = 0

func add(done chan bool) {
    for j :=0; j<1000000;j++ {
    i++
    done <- true
    }
}
func sub(done chan bool) {
    for j :=0; j<1000000;j++ {
    i--
    done <- true
    }
}

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())    // I guess this is a hint to what GOMAXPROCS does...
                                            // Try doing the exercise both with and without it! 
    done := make(chan bool, 1)
    go add(done)
    go sub(done)		                    
    <- done
    time.Sleep(100*time.Millisecond)
    Println("%d\n",i)
}
