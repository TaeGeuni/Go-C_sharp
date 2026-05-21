# 2657-find-the-prefix-common-array-of-two-arrays

## 💡 문제 링크
* [LeetCode - Find the Prefix Common Array of Two Arrays](https://leetcode.com/problems/find-the-prefix-common-array-of-two-arrays/)

## 🧠 문제 접근 및 풀이
* **알고리즘 분류:** Array, Counting (Hash Table 최적화)
* **시간 복잡도:** $O(N)$
* **공간 복잡도:** $O(N)$

### 핵심 아이디어
* 두 배열 `A`와 `B`가 `1`부터 `n`까지의 정수로 이루어진 순열(Permutation)이라는 점을 이용합니다.
* 배열을 순회하면서 특정 숫자가 등장했는지를 기록합니다. 
* 이미 등장했던 숫자가 다시 등장했다면, 이는 `A`와 `B` 양쪽 모두에서 최소 한 번씩 출현했다는 의미이므로 공통 원소의 개수(`num`)를 1 증가시킵니다.

## 📝 배운 점

### 1. 알고리즘: 해시맵(Hash Map) 대신 배열(Array) 사용하기
처음에는 숫자의 등장 여부를 기록하기 위해 해시맵을 사용해야 한다고 생각했습니다. 하지만 문제의 조건에서 **숫자의 범위가 배열의 길이(`n`)와 같은 순열**이라는 것을 파악했습니다.
* 해시맵을 생성하고 해시를 계산하는 오버헤드를 줄이기 위해, `n+1` 크기의 `bool` 배열(`freq`)을 사용하는 것이 훨씬 효율적이라는 것을 배웠습니다. 
* 값 자체가 배열의 인덱스가 되는 방식(Direct Addressing)으로, 조회와 갱신을 $O(1)$의 비용으로 아주 가볍게 처리할 수 있습니다.

### 2. 언어 스터디: C# 배열 선언 방법
Go 언어의 슬라이스/배열 선언 방식과 C#의 배열 선언 방식을 비교하며 C# 문법을 익혔습니다.

| 언어 | 배열(또는 슬라이스) 선언 및 초기화 문법 |
|---|---|
| **Go** | `res := make([]int, n)`<br>`freq := make([]bool, n+1)` |
| **C#** | `int[] res = new int[n];`<br>`bool[] freq = new bool[n + 1];` |

* C#에서는 타입 뒤에 `[]`를 붙여 배열임을 명시하고, `new` 키워드와 함께 크기를 지정하여 힙(Heap) 메모리에 할당한다는 것을 알게 되었습니다.

## 🏃‍♂️ 코드 실행
* **Go:** `go run main.go`
* **C#:** `dotnet run` (현재 디렉토리에서)

---

## 💻 풀이 코드

### Go
```go
package main

func findThePrefixCommonArray(A []int, B []int) []int {
    n := len(A)
    res := make([]int, n)

    freq := make([]bool, n+1)
    num := 0

    for i := 0; i < n; i++ {
        if !freq[A[i]] {
            freq[A[i]] = true
        } else {
            num++
        }
        
        if !freq[B[i]] {
            freq[B[i]] = true
        } else {
            num++
        }
        res[i] = num
    }

    return res
}
```
---
### C#
```c#
using System;

public class Solution
{
    public int[] FindThePrefixCommonArray(int[] A, int[] B)
    {
        int n = A.Length;
        int[] res = new int[n];

        bool[] freq = new bool[n + 1];
        int num = 0;

        for (int i = 0; i < n; i++)
        {
            if (!freq[A[i]])
            {
                freq[A[i]] = true;
            }
            else
            {
                num++;
            }

            if (!freq[B[i]])
            {
                freq[B[i]] = true;
            }
            else
            {
                num++;
            }

            res[i] = num;
        }
        return res;
    }
}
```