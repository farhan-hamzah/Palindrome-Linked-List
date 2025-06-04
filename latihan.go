package main

import (
	"fmt"
)
const NMAX int = 100
type palindrom [NMAX]int
func main(){
    var A palindrom
    var n, i, p1, p2 int
    fmt.Scan(&n)
    for i = 0; i < n; i++{
        fmt.Scan(&A[i])
    }
    p1 = 0
    for i = 0; i < n; i++{
        p1 = (p1*10)+A[i]
    }
    p2 = 0
    for i = n-1; i >= 0; i--{
        p2 = (p2*10)+A[i]
    }
    if p1 == p2{
        fmt.Print(true)
    }else{
        fmt.Print(false)
    }
}