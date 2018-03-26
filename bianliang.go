
package main

import "fmt"
var g int
func main() {
    /*声明局部变量*/    
    var a,b,c int
    /*初始化参数*/
    a = 10
    b = 20
    c = a + b
    g = a + b
    fmt.Printf("结果:a=%d,b=%d and c=%d\n",a,b,c)
    fmt.Printf("结果:a=%d,b=%d and g=%d\n",a,b,g)
    g = sum(a,b)
    fmt.Printf("main() 函数中g = %d\n",g)
}


func sum(a,b int) int {
    fmt.Printf("sum() 函数中a = %d\n",a)
    fmt.Printf("sum() 函数中b = %d\n",b)
    return a + b
}
