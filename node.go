package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var a = []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2}
	var b = []int{5, 6, 4}
	aNode := buildNode(a)
	printNode(aNode)
	bNode := buildNode(b)
	printNode(bNode)
	sumNode := addTwoNumbers(buildNode(a), buildNode(b))
	printNode(sumNode)

}
func buildNode(a []int) *ListNode {
	result := new(ListNode)

	if len(a) == 0 {
		return result
	}
	temp := result
	for i := range a {
		result.Val = a[i]
		l := new(ListNode)
		if i != len(a)-1 {
			result.Next = l
			result = result.Next
		}
	}

	return temp
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	cl1 := convert(l1)
	cl2 := convert(l2)

	fmt.Println("======")
	sum1 := 0
	for {
		if cl1 == nil {
			break
		}

		fmt.Println(fmt.Sprintln("val: %d", cl1.Val))
		sum1 = sum1*10 + cl1.Val
		fmt.Println(fmt.Sprintln("sum1: %d", sum1))
		cl1 = cl1.Next
	}

	fmt.Println(fmt.Sprintln("sum1 total: %d", sum1))
	sum2 := 0
	for {
		if cl2 == nil {
			break
		}

		fmt.Println(fmt.Sprintln("val: %d", cl2.Val))
		sum2 = sum2*10 + cl2.Val
		fmt.Println(fmt.Sprintln("sum2: %d", sum2))
		cl2 = cl2.Next
	}

	fmt.Println(fmt.Sprintln("sum2 total: %d", sum2))
	sum := sum1 + sum2
	fmt.Println(fmt.Sprintln("sum: %d", sum))

	result := new(ListNode)
	temp := result

	for {
		if sum < 10 {
			temp.Val = sum
			break
		}
		temp.Val = sum % 10
		sum = sum / 10
		temp.Next = new(ListNode)
		temp = temp.Next
	}
	return result
}

func convert(node *ListNode) *ListNode {
	fmt.Println("start=====")

	var prev *ListNode
	var swap *ListNode
	prev = nil
	swap = nil

	for {
		if node == nil {
			break
		}
		fmt.Println(node.Val)
		swap = node.Next
		node.Next = prev
		prev = node
		node = swap
	}

	return prev
}
func printNode(node *ListNode) {
	for {
		if node == nil {
			break
		}
		fmt.Println(fmt.Sprintln("result item: %d", node.Val))
		node = node.Next
	}
}
