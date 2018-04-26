package main

import "fmt"

func main() {

    nums := []int{2,3,4}
    sum := 0
    for _,num := range nums {
        sum += num
    }
    fmt.Println("sum:",sum)

    for i,num := range nums {
        if num == 3 {
            fmt.Println("index:",i)
        }
    }
    //range也可以用在map的健值对上
    kvs := map[string] string {"a":"apple"}
    for k,v := range kvs {
        fmt.Printf("%s -> %s\n",k,v)
    }

    for i,c := range "go" {
        fmt.Println(i,c)
    }
}
