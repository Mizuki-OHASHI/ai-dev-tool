## 概要

- autodebug テスト 2-3
- 問題は UTTC 研修プログラム Go/課題 2 より引用 (一部改)
- 説明を減らした
- クエリの修正 (+ `want ... but got ...`)

## 結果

- 初め間違ったコードを生成した, want/got 付きのエラーを送信したところ, 正しいコードを生成した

| No. | want | got | remarks                      |
| :-- | :--: | :-: | :--------------------------- |
| 0   |  19  | 152 | oops...                      |
| 1   |  19  | 19  | noticed logical error! `(*)` |

`(*)` evidence ?

```shell
The error in the code is that the `countCardCombinations` function is not correctly calculating the number of combinations. It is using the wrong logic for counting combinations. ...
```

- 生成された正しいコード

```go
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var m string = "1, 3, 3, 4, 5, 5, 5, 7, 10"
	var s int = 15

	a := convertStringToIntArray(m)
	frequencyCount := countNumberFrequency(a)
	combinationCount := countCardCombinations(a, s)

	printMapKeyAndValue(frequencyCount)
	fmt.Println(combinationCount)

	if combinationCount != 19 {
		panic(fmt.Errorf("failed test case: want 19 but got %v", combinationCount))
	}
}

func convertStringToIntArray(s string) []int {
	strArr := strings.Split(s, ", ")
	intArr := make([]int, len(strArr))

	for i, str := range strArr {
		num, err := strconv.Atoi(str)
		if err != nil {
			panic(fmt.Errorf("failed to convert string to int: %v", err))
		}
		intArr[i] = num
	}

	return intArr
}

func countNumberFrequency(arr []int) map[int]int {
	frequencyCount := make(map[int]int)

	for _, num := range arr {
		frequencyCount[num]++
	}

	return frequencyCount
}

func countCardCombinations(arr []int, targetSum int) int {
	combinations := make([]int, targetSum+1)
	combinations[0] = 1

	for _, num := range arr {
		for i := targetSum; i >= num; i-- {
			combinations[i] += combinations[i-num]
		}
	}

	return combinations[targetSum]
}

func printMapKeyAndValue(m map[int]int) {
	for key, value := range m {
		fmt.Printf("%d: %d\n", key, value)
	}
}
```
