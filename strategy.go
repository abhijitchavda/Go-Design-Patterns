package main

import "fmt"
import "math/rand"

var unSortAr [11]int = [11]int{1, 5, 8, 5, 3, 5, 9, 4, 7, 6, 2}

type Sortint interface {
	sortAlgo() []int
}

var sortReference Sortint

type bubbleSort struct {
}

//--------implements bubble sort logic-----------
func (bs bubbleSort) sortAlgo() []int {

	fmt.Println("\nBubbleSort Implementation")
	arry := unSortAr
	for i := 1; i < len(arry); i++ {
		for j := 0; j < len(arry)-i; j++ {
			if arry[j] > arry[j+1] {
				arry[j], arry[j+1] = arry[j+1], arry[j]
			}
		}
	}
	return arry[:]
}

//---------------------------------------------

type mergeSort struct {
}

//------------implements merge sort logic----------
func (ms mergeSort) sortAlgo() []int {
	fmt.Println("\nMergeSort Implementation")
	arry := MrgSortprominent(unSortAr[:])
	return arry[:]
}

func MrgSortprominent(slice []int) []int {

	if len(slice) < 2 {
		return slice
	}
	mid := (len(slice)) / 2
	return Mrgit(MrgSortprominent(slice[:mid]), MrgSortprominent(slice[mid:]))
}

func Mrgit(left, right []int) []int {

	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int, size, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	return slice
}

//-----------------------------------------------

type quickSort struct {
}

//--------implements quick sort logic------------
func (qs quickSort) sortAlgo() []int {
	fmt.Println("\nQuickSort Implementation")
	arry := quickSortprominent(unSortAr[:])
	return arry[:]
}

func quickSortprominent(slice []int) []int {
	length := len(slice)

	if length <= 1 {
		sliceCopy := make([]int, length)
		copy(sliceCopy, slice)
		return sliceCopy
	}

	m := slice[rand.Intn(length)]

	less := make([]int, 0, length)
	middle := make([]int, 0, length)
	more := make([]int, 0, length)

	for _, item := range slice {
		switch {
		case item < m:
			less = append(less, item)
		case item == m:
			middle = append(middle, item)
		case item > m:
			more = append(more, item)
		}
	}

	less, more = quickSortprominent(less), quickSortprominent(more)

	less = append(less, middle...)
	less = append(less, more...)

	return less
}

//--------------------------------------------------

//---------take type of reference stored in sortReference and call sorting logic accordingly
func Sortit() []int {

	return sortReference.sortAlgo()

}

//-----change the strategy for sorting---------
func changeStatergy(sortInterface Sortint) {
	sortReference = sortInterface
}

func cleaning() []int {
	var a = [11]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	fmt.Println("reseting the sorted array")
	return a[:0]
}

func main() {

	var SortedArray []int

	bubble := bubbleSort{}
	mrg := mergeSort{}
	quick := quickSort{}

	fmt.Println("Unsorted Array:", unSortAr)

	changeStatergy(bubble)
	SortedArray = Sortit()
	fmt.Println(SortedArray)

	SortedArray = cleaning()

	changeStatergy(mrg)
	SortedArray = Sortit()
	fmt.Println(SortedArray)

	SortedArray = cleaning()

	changeStatergy(quick)
	SortedArray = Sortit()
	fmt.Println(SortedArray)
}
