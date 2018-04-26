package main

import "fmt"


func main() {
    numbers := []int{0,1,2,3,4,5,6,7,8}
    printSlice(numbers)

    fmt.Println("numbers == ",numbers)
    fmt.Println("numbers [1:4] == ",numbers[1:4])

    fmt.Println("numbers[:3] ==",numbers[:3])
    fmt.Println("numbers[4:] ==",numbers[4:])
    numbers1 := make([]int,0,5)
    printSlice(numbers1)
    number2 := numbers[:2]
    printSlice(number2)
    number3 := numbers[2:5]
    printSlice(number3)



    var numbers0 []int
    printSlice(numbers0)
    numbers0 = append(numbers,0)
    printSlice(numbers0)
    
    numbers0 = append(numbers0,0)
    printSlice(numbers0)
    numbers0 = append(numbers0,1)
    printSlice(numbers0)
    numbers0 = append(numbers0,2,3,4) 
    printSlice(numbers0)
    numbers2 := make([]int,len(numbers0),(cap(numbers0))*2)
    copy(numbers2,numbers0)
    
    printSlice(numbers2)
}

func printSlice(x []int) {
    fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)

}
