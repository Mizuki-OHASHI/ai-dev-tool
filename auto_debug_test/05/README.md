## 概要

- - autodebug テスト 2-2
- 問題は UTTC 研修プログラム Go/課題 2 より引用 (一部改)

## 結果

- 正常に生成された

```go
package main

import (
        "fmt"
        "strconv"
        "strings"
)

// 文字列を整数の配列に変換する関数
func convertStringToIntArray(m string) []int {
        // 文字列をカンマで分割
        strArr := strings.Split(m, ", ")

        // 整数の配列に変換
        intArr := make([]int, len(strArr))
        for i, str := range strArr {
                num, _ := strconv.Atoi(str)
                intArr[i] = num
        }

        return intArr
}

// 各整数の出現回数を数える関数
func countNumberFrequency(a []int) map[int]int {
        frequency := make(map[int]int)

        for _, num := range a {
                frequency[num]++
        }

        return frequency
}

// 整数の合計をsにする方法が何通りあるかを数える関数
func countCardCombinations(a []int, s int) int {
        // メモ化テーブルを作成
        dp := make([]int, s+1)
        dp[0] = 1

        // 各整数を使って合計を作る方法の数を計算
        for _, num := range a {
                for i := s; i >= num; i-- {
                        dp[i] += dp[i-num]
                }
        }

        return dp[s]
}

// map[int]intのkeyとvalueをkeyの昇順に出力する関数
func printMapKeyAndValue(m map[int]int) {
        for key := range m {
                fmt.Println(key, m[key])
        }
}

func main() {
        var m string = "1, 3, 3, 4, 5, 5, 5, 7, 10"
        var s int = 15

        a := convertStringToIntArray(m)
        frequencyCount := countNumberFrequency(a)
        combinationCount := countCardCombinations(a, s)

        printMapKeyAndValue(frequencyCount)
        fmt.Println(combinationCount)

        if combinationCount != 19 {
                panic(fmt.Errorf("failed test case"))
        }
}
```
