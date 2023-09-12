package main

import (
	"fmt"
	"go-leetcode/leetcode/hot100"
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

	//words := leetcode.ReverseWords("i like you ")
	//fmt.Println(words)

	//arr := []int{0, 0, 1, 2, 4, 2, 2, 3, 1, 4}
	//numbers := getOffer.GetLeastNumbers(arr, 4)
	//fmt.Println(numbers)

	//arr := [][]int{[]int{9, 3, 6}, []int{8, 1, 7}, []int{6, 6, 8}, []int{8, 4, 9}, []int{4, 2, 9}}
	//arr := [][]int{{5, 4, 7}, {7, 4, 8}, {4, 1, 8}}
	//arr := [][]int{{7, 5, 6}, {6, 7, 8}, {10, 1, 6}}
	//arr := [][]int{{3, 6, 9}, {10, 2, 3}, {1, 6, 8}, {2, 1, 6}, {9, 3, 9}}
	//arr := [][]int{{10, 5, 7}, {10, 3, 4}, {7, 1, 8}, {6, 3, 4}}
	//arr := [][]int{{2, 2, 7}, {9, 5, 6}, {10, 1, 7}, {3, 3, 6}, {2, 1, 4}}
	//pooling := leetcode.CarPooling(arr, 24)
	//fmt.Println(pooling)

	//nums := getOffer.SumNums(4)
	//fmt.Println(nums)

	//arr := []int{3, 4, 2, 1, 1, 0}
	//number := getOffer.FindRepeatNumber(arr)
	//fmt.Println(number)

	//arr := [][]int{
	//	{1, 3, 5, 7, 9},
	//	{2, 4, 6, 8, 10},
	//	{11, 13, 15, 17, 19},
	//	{12, 14, 16, 18, 20},
	//	{21, 22, 23, 24, 25},
	//}

	//arr := [][]int{
	//	{1}, {3}, {5},
	//}
	//number := getOffer.FindNumberIn2DArray(arr, 13)
	//fmt.Println(number)

	//words := getOffer.ReverseLeftWords("lrloseumgh", 6)
	//fmt.Println(words)

	//fib := getOffer.Fib(5)
	//fmt.Println(fib)

	//exchange := getOffer.Exchange([]int{1, 2, 3, 4})
	//fmt.Println(exchange)

	//search := getOffer.Search([]int{2, 2}, 2)
	//fmt.Println(search)

	//substring := hot100.LengthOfLongestSubstring("tmmzuxt")
	//fmt.Println(substring)

	s := "([}}])"
	valid := hot100.IsValid(s)
	fmt.Println(valid)
}
