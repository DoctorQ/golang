package main
import "fmt"
func main() {
    var a,b int
    a,b = 100,200
    var ret int
 
    ret = max(a,b)
    fmt.Printf("最大值:%d\n",ret)
    c,d := swap("doctorq","doctorq1")
    fmt.Println(c,d)

}
func max(num1,num2 int) int {
    var result int
    if (num1 > num2) {
    	 result = num1
    }else {
        result = num2
    }
    return result
}


func swap(x,y string) (string,string) {
    return y,x
}


