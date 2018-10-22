/*
Package smallest finds the smallest number in this list:

x :=[] int{
    48,96,86,68,
    57,82,63,70,
    37,34,83,27,
    19,97, 9,17,
}
*/
package main

import "fmt"

/*from function QuickSort https://varunpant.com/posts/basic-sorting-algorithms-implemented-in-golang */
func QuickSort(items []int) {

	if len(items) > 1 {
		pivot_index := len(items) / 2
		var smaller_items = []int{}
		var larger_items = []int{}

		for i := range items {
			val := items[i]
			if i != pivot_index {
				if val < items[pivot_index] {
					smaller_items = append(smaller_items, val)
				} else {
					larger_items = append(larger_items, val)
				}
			}
		}

		QuickSort(smaller_items)
		QuickSort(larger_items)

		var merged []int = append(append(append([]int{}, smaller_items...), []int{items[pivot_index]}...), larger_items...)
		//merged := MergeLists(smaller_items, items[pivot_index], larger_items)

		for j := 0; j < len(items); j++ {
			items[j] = merged[j]
		}

	}

}

/* end: from function QuickSort https://varunpant.com/posts/basic-sorting-algorithms-implemented-in-golang   */

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	fmt.Println(x[0])

	QuickSort(x)
	fmt.Println(x[0])

}
