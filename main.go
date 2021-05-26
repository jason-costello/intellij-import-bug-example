package main

import(
    "fmt"
)

func main(){
    ctx := context.TODO()
    if ctx == nil{
        fmt.Println("nil context")
    } else {
        fmt.Println("not nil context")
}
}
