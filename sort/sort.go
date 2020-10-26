package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//test := []int{3, 2, 4, 5, 6, 2, 1}
	arr := make([]int, 10000)
	brr := make([]int, 10000)
	crr := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data := rand.Intn(1000)
		arr[i] = data
		brr[i] = data
		crr[i] = data
	}
	aST := time.Now().UnixNano()
	bubbleSort(arr)
	aET := time.Now().UnixNano()
	InsertionSort(brr)
	bET := time.Now().UnixNano()
	selectionSort(crr)
	cET := time.Now().UnixNano()
	fmt.Println(aET - aST)
	fmt.Println(bET - aET)
	fmt.Println(cET - bET)
	//fmt.Println(arr)
	//fmt.Println(brr)
}

//冒泡排序 (时间复杂度最好 O(n) 最坏O(n^2) 平均O(n^2) 空间复杂度O(1))
func bubbleSort(arr []int) {
	lenArr := len(arr)
	if lenArr <= 1 {
		return
	}
	for i := 0; i < lenArr; i++ {
		//如果没有交换操作，则说明已经排好序，提前结束循环
		tmpFlag := true
		for j := 0; j < lenArr-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				tmpFlag = false
			}
		}
		if tmpFlag {
			break
		}
	}
}

//插入排序 (时间复杂度最好 O(n) 最坏O(n^2) 平均O(n^2) 空间复杂度O(1))
func InsertionSort(arr []int) {
	lenArr := len(arr)
	if lenArr <= 1 {
		return
	}
	for i := 1; i < lenArr; i++ {
		tmpValue := arr[i]
		j := i - 1
		for ; j >= 0; j-- {
			if arr[j] > tmpValue {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		arr[j+1] = tmpValue
	}
}

//选择排序 (时间复杂度最好 O(n^2) 最坏O(n^2) 平均O(n^2) 空间复杂度O(1))
func selectionSort(arr []int) {
	lenArr := len(arr)
	for i := 0; i < lenArr; i++ {
		index := i
		for j := i + 1; j < lenArr-1; j++ {
			if arr[index] > arr[j]{
				index = j
			}
		}
		if index != i {
			arr[i], arr[index] = arr[index], arr[i]
		}
	}
}
