package main

import (
	"fmt"
)

/*
You are given two non-empty linked lists representing two non-negative integers.
The digits are stored in reverse order and each of their nodes contain a single digit.
Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:
	Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
	Output: 7 -> 0 -> 8
	Explanation: 342 + 465 = 807.
*/

/*
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	// 先处理相同的数据
	head := &ListNode{}
	tmp := head
	tl1, tl2 := l1, l2
	l := l1

	total := 0
	for tl1 != nil && tl2 != nil {
		total += tl1.Val + tl2.Val
		tl1 = tl1.Next
		tl2 = tl2.Next

		tmp.Next = &ListNode{
			Val:  total % 10,
			Next: nil,
		}
		tmp = tmp.Next
		total /= 10
	}

	if tl1 != nil {
		l = tl1
	} else if tl2 != nil {
		l = tl2
	} else {
		l = nil
	}

	// 处理最后的数据
	for l != nil {
		total += l.Val
		l = l.Next

		tmp.Next = &ListNode{
			Val:  total % 10,
			Next: nil,
		}
		tmp = tmp.Next
		total /= 10
	}

	if total > 0 {
		tmp.Next = &ListNode{
			Val:  total % 10,
			Next: nil,
		}
	}

	// 处理头结点
	if head.Next != nil {
		head = head.Next
	}

	return head
}

func (s *ListNode) print() {

	for s != nil {
		fmt.Printf("%d ", s.Val)
		s = s.Next
	}
	fmt.Println()
}

func instanceListNode(array []int) *ListNode {
	head := &ListNode{}
	tmp := head
	if len(array) > 0 {
		for _, num := range array[:] {
			tmp.Next = &ListNode{
				Val:  num,
				Next: nil,
			}
			tmp = tmp.Next
		}
	}
	return head.Next
}

func main() {

	/*
		[1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1]
		[5,6,4]
	*/

	num1 := []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	num2 := []int{5, 6, 4}

	l1 := instanceListNode(num1)
	l2 := instanceListNode(num2)

	l1 = instanceListNode([]int{5})
	l2 = instanceListNode([]int{5})

	ret := addTwoNumbers(l1, l2)

	ret.print()
}
