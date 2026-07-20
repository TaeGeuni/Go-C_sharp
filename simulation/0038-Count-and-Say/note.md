# 0038-count-and-say

## 💡 문제 링크
* [LeetCode - Count and Say](https://leetcode.com/problems/count-and-say/)

## 🧠 문제 접근 및 풀이
* **알고리즘 분류:** String, Simulation
* **시간 복잡도:** O(N * M) (N은 반복 횟수, M은 생성되는 문자열의 최대 길이)
* **공간 복잡도:** O(M) (다음 문자열을 생성하기 위한 버퍼 배열 공간)

### 핵심 아이디어
* 이전 단계의 문자열을 처음부터 끝까지 순회하며, 연속된 같은 숫자의 개수(`count`)와 해당 숫자(`now`)를 세어 다음 문자열을 만들어냅니다.
* 새로운 알고리즘 기법을 요구하기보다는, 문제의 규칙(읽고 말하기)을 코드로 정확하게 옮기는 시뮬레이션 구현 문제입니다.

## 📝 배운 점

### 1. 언어 스터디: 문자열 버퍼(Buffer) 활용 기법 (Go vs Python)
문자열을 순회하며 빈번하게 문자를 이어 붙여야(Concatenation) 할 때, 각 언어에서 제공하는 버퍼(Buffer) 활용법을 익혔습니다.

| 언어 | 문자열 합치기 최적화 방식 |
|---|---|
| **Go** | `strings.Builder`를 사용하여 메모리 할당을 최소화하면서 문자열을 이어 붙입니다. (`buf.WriteString()`) |
| **Python** | 파이썬의 문자열은 불변(Immutable) 객체이므로 `+=`로 합치면 매번 새로운 객체가 생성되어 비효율적입니다. 따라서 **리스트(`[]`)를 버퍼로 사용하여 요소들을 `append()` 한 뒤, 마지막에 `"".join()`으로 한 번에 문자열로 변환**하는 것이 정석입니다. |

## 🏃‍♂️ 코드 실행
* **Go:** `go run main.go`
* **Python:** `python main.py`

---

## 💻 풀이 코드

### Go
```go
package main

import (
    "strconv"
    "strings"
)

func countAndSay(n int) string {
    res := "1"
    for i := 1; i < n; i++ {
        var buf strings.Builder
        count := 0
        var now byte
        for j := 0; j < len(res); j++ {
            if j == 0 {
                now = res[j]
                count = 1
            } else {
                if now == res[j] {
                    count++
                } else {
                    buf.WriteString(strconv.Itoa(count))
                    buf.WriteString(string(now))
                    now = res[j]
                    count = 1
                }
            }
        }
        buf.WriteString(strconv.Itoa(count))
        buf.WriteString(string(now))
        res = buf.String()
    }
    return res
}
```
---
### Python

```python
class Solution:
    def countAndSay(self, n: int) -> str:
        res = "1"
        for i in range(1, n):
            buf = []
            count = 1
            now = res[0]
            for j in range(1, len(res)):
                if now == res[j]:
                    count += 1
                else:
                    buf.append(str(count))
                    buf.append(now)
                    now = res[j]
                    count = 1
            buf.append(str(count))
            buf.append(now)
            
            # 리스트 버퍼를 문자열로 합쳐서 다음 단계를 위해 res 갱신
            res = "".join(buf) 
            
        return res
```