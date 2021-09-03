package main

import (
	"math/rand"
	"reflect"
)

const size int = 11

type queue struct {
	pod []int
}

func (q *queue) push(val int) {
	q.pod = append(q.pod, val)
}

func (q *queue) pop() int {
	top := q.pod[0]
	q.pod = q.pod[1:]
	return top
}

func (q *queue) size() int {
	return len(q.pod)
}

type Arr [size]int

func weak_partition(arr []int, i int, j int, piv int) int {

	swapF := reflect.Swapper(arr)

	for i < j {
		if arr[i] < piv {
			i++
			continue
		}

		if arr[i] == piv {
			i = weak_partition(arr, i+1, j, piv)
			continue
		}

		for i < j {

			if arr[j] <= piv {
				swapF(i, j)
				j--
				break
			}
			j--
		}
	}

	return i
}

func pivots2right(arr []int, piv int) int {

	swapF := reflect.Swapper(arr)

	var smt queue = queue{make([]int, 0)} // smaller than piv queue

	i := len(arr) - 1
	lm := 0 // left most pivot

	/*
	* Scan valus from right to left.
	*  - Remeber the position of right most smaller than pivot.
	*  - For each pivot, try swap with right most smaller than pivot.
	*
	 */
	for i >= 0 {

		if arr[i] == piv {

			if smt.size() == 0 {
				lm = i
				i--
				continue
			}

			lm = smt.pop()
			swapF(i, lm)
			continue
		}

		smt.push(i)
		i--
	}

	return lm
}

/*
* @Desc: Partition implementation.
* @PROC: 1. Do soft partitioning with weak_partition():
*	    - All smaller or equales to pivot are on the left.
*	    - All greater than pivot are on the right.
*	 2. Sort dangling pivots to the right with sort_pivots().
 */
func partition(arr []int, i int, j int, piv int) (int, int) {

	b_piv := weak_partition(arr, i, j, piv)
	a_piv := pivots2right(arr[0:b_piv], piv)

	return a_piv, b_piv
}

/*
* @Desc: QuickSort implementation.
* @PROC: 1. Pick random pivot.
*	 2. Orginize elements around pivot.
*	 3. Run again over smaller & largerer elements.
 */
func qsort_impl(arr []int) {

	if len(arr) <= 1 {
		return
	}

	m := rand.Intn(len(arr))
	piv := arr[m]
	i, j := 0, len(arr)-1

	i, j = partition(arr, i, j, piv)
	qsort_impl(arr[0:i])
	qsort_impl(arr[j:])
}

/*
* @Desc: Warrper method.
 */
func (arr *Arr) qsort() {

	qsort_impl(arr[0:])
}

func main() {

	var arr Arr
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(size * 30)
	}

	arr.qsort()
}
