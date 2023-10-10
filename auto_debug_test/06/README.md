## 概要

- - autodebug テスト 2-3
- 問題は UTTC 研修プログラム Go/課題 2 より引用 (一部改)
- 説明を減らした

## 結果

- ロジックが間違っている

```go
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func convertStringToIntArray(s string) []int {
	arr := strings.Split(s, ", ")
	result := make([]int, len(arr))

	for i, v := range arr {
		num, _ := strconv.Atoi(v)
		result[i] = num
	}

	return result
}

func countNumberFrequency(arr []int) map[int]int {
	frequencyCount := make(map[int]int)

	for _, num := range arr {
		frequencyCount[num]++
	}

	return frequencyCount
}

func countCardCombinations(arr []int, targetSum int) int {
	combinations := make(map[string]bool)

	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			sum := 0
			key := ""

			for k := i; k <= j; k++ {
				sum += arr[k]
				key += strconv.Itoa(arr[k]) + ","
			}

			if sum == targetSum {
				combinations[key[:len(key)-1]] = true
			}
		}
	}

	return len(combinations)
}

func printMapKeyAndValue(m map[int]int) {
	for k, v := range m {
		fmt.Printf("%d: %d\n", k, v)
	}
}

func main() {
	var m string = "1, 3, 3, 4, 5, 5, 5, 7, 10"
	var s int = 15

	a := convertStringToIntArray(m)
	frequencyCount := countNumberFrequency(a)
	combinationCount := countCardCombinations(a, s)

	printMapKeyAndValue(frequencyCount)

	/*if combinationCount != 19 {
		panic(fmt.Errorf("failed test case"))
	}*/

	fmt.Println(combinationCount)
}
```
