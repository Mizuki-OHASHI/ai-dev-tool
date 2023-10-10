## 概要

- autodebug テスト 2-1
- 問題は UTTC 研修プログラム Go/課題 1 より引用 (一部改)

## 結果

- 正常に生成された

```go
package main

import "fmt"

func countCardCombinations(s int, a []int) int {
    return helper(s, a, len(a)-1)
}

func helper(s int, a []int, i int) int {
    if s == 0 {
        return 1
    }
    if s < 0 || i < 0 {
        return 0
    }
    return helper(s-a[i], a, i-1) + helper(s, a, i-1)
}

func main() {
    if countCardCombinations(15, []int{1, 3, 3, 4, 5, 5, 5, 7, 10}) != 19 {
        panic(fmt.Errorf("Test case failed"))
    }
}
```
