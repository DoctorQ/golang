package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//var a = []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	//var b = []int{5, 6, 4}
	//aNode := buildNode(a)
	//printNode(aNode)
	//bNode := buildNode(b)
	//printNode(bNode)
	//sumNode := addTwoNumbers(aNode, bNode)
	//printNode(sumNode)

	a := []int{2, 4, 3}
	b := []int{5, 6, 4}
	sumNode := addTwoNumbers(buildNode(a), buildNode(b))
	printNode(sumNode)
	a = []int{5}
	b = []int{5}
	sumNode = addTwoNumbers(buildNode(a), buildNode(b))
	printNode(sumNode)
	a = []int{1}
	b = []int{9, 9}
	sumNode = addTwoNumbers(buildNode(a), buildNode(b))
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
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	t1 := l1
	t2 := l2
	temp := new(ListNode)
	var a []int
	for {
		if t1 == nil || t2 == nil {
			if t1 == nil {
				temp = t2
				fmt.Println("t1 over")
			} else {
				temp = t1
				fmt.Println("t2 over")
			}
			for {
				if temp == nil {
					break
				}
				if temp.Val > 9 {
					if temp.Next == nil {
						temp.Next = new(ListNode)
					}
					temp.Next.Val = temp.Next.Val + 1
					temp.Val = temp.Val % 10
				}
				a = append(a, temp.Val)
				temp = temp.Next
			}
			break
		}
		if temp == nil {
			temp = new(ListNode)
		}
		total := t1.Val + t2.Val
		bt := 0
		if total > 9 {
			bt = total / 10
			total = total % 10
		}
		temp.Val = temp.Val + total
		a = append(a, temp.Val)
		fmt.Println(fmt.Sprintf("t1:%d,t2:%d,total:%d,bt:%d", t1.Val, t2.Val, temp.Val, bt))

		//超过10，给上一位加1
		if bt != 0 {
			if t1.Next != nil {
				t1.Next.Val = t1.Next.Val + bt
				fmt.Println(fmt.Sprintf("bt:%d", bt))
			} else if t2.Next != nil {
				t2.Next.Val = t2.Next.Val + bt
			} else {
				//两者都为空特殊情况
				temp.Next = new(ListNode)
				temp.Next.Val = bt
				a = append(a, bt)
				break
			}
		}
		t1 = t1.Next
		t2 = t2.Next
		temp = temp.Next
	}

	result := buildNode(a)
	return result
}

func convert(node *ListNode) *ListNode {
	//fmt.Println("convert=====")

	var prev *ListNode
	var swap *ListNode
	prev = nil
	swap = nil

	for {
		if node == nil {
			break
		}
		//fmt.Println(node.Val)
		swap = node.Next
		node.Next = prev
		prev = node
		node = swap
	}

	return prev
}
func printNode(node *ListNode) {

	fmt.Println("print=====")
	for {
		if node == nil {
			break
		}
		fmt.Println(fmt.Sprintln("result item: %d", node.Val))
		node = node.Next
	}
}
