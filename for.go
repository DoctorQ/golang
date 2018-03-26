
package main
import "fmt"
func main() {
	var C,c int //变量声明
	C = 1
	A: for C < 100 {
		C++
		for c = 2; c < C;c++ {
			if C%c == 0 {
				goto A
			}
		}
		fmt.Println(C,"是素数")
	}

}
