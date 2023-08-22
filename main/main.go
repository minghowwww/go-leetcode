package main

import (
	"fmt"
	"go-leetcode/leetcode"
)

func main() {
	//nums := []int{-1, 0, 3, 5, 9, 12}
	//target := 9
	//i := leetcode.Search(nums, target)
	//fmt.Println(i)

	//nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	//val := 2
	//element := leetcode.RemoveElement1(nums, val)
	//fmt.Println(element)

	//nums := []int{-4, -1, 0, 3, 10}
	//squares := leetcode.SortedSquares(nums)
	//fmt.Println(squares)

	// nums := []int{2, 3, 1, 2, 4, 3}
	// target := 7
	// arrayLen := leetcode.MinSubArrayLen(target, nums)
	// fmt.Println(arrayLen)

	// i := leetcode.GenerateMatrix(3)
	// fmt.Println(i)

	//myLinkedList := leetcode.Constructor()
	//myLinkedList.AddAtHead(1)
	//myLinkedList.AddAtTail(3)
	//myLinkedList.AddAtIndex(1, 2)
	//get := myLinkedList.Get(1)
	//myLinkedList.DeleteAtIndex(2)
	//fmt.Println(get)

	//num1 := []int{1, 2, 2, 1}
	//num2 := []int{2, 2}
	//intersection := leetcode.Intersection(num1, num2)
	//fmt.Println(intersection)

	//nums := []int{-2, 0, 0, 2, 2}
	//sum := leetcode.ThreeSum(nums)
	//fmt.Println(sum)

	//s := "abcdefg"
	//k := 8
	//str := leetcode.ReverseStr(s, k)
	//fmt.Println(str)

	words := leetcode.ReverseWords("i like you ")
	fmt.Println(words)
}
