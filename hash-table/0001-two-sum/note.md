# 0001-two-sum

## 💡 문제 링크
* [LeetCode - Two Sum](https://leetcode.com/problems/two-sum/)

## 🧠 문제 접근 및 풀이
* **알고리즘 분류:** Hash Table
* **시간 복잡도:** $O(N)$
* **공간 복잡도:** $O(N)$

### 핵심 아이디어
* 배열을 두 번 순회할 필요 없이, `단 한 번만 순회(One-pass)`하며 문제를 해결합니다.
* 현재 순회 중인 배열의 값(`nums[i]`)을 기준으로, 목표값을 만들기 위해 필요한 나머지 값(`target - nums[i]`)이 이미 해시맵에 존재하는지 확인합니다.
* 짝이 되는 값이 존재한다면 바로 인덱스를 반환하고, 없다면 현재 값을 해시맵에 추가한 뒤 다음 요소로 넘어갑니다.

## 📝 배운 점

### 1. 알고리즘: 효율적인 해시맵 활용
처음에는 배열의 모든 내용을 해시맵에 전부 집어넣은 뒤, 다시 배열을 돌며 값을 비교하는 방식을 생각했습니다. 하지만 코드를 읽으면서 `target`에서 현재 값을 뺀 값이 맵에 있는지 실시간으로 조회하고 저장하면, **배열을 한 번만 읽어도 해결 가능하다**는 것을 배웠습니다.

### 2. 언어 스터디: Go vs C# 해시맵(Dictionary) 문법 비교
Go 언어로 푼 로직을 C#으로 옮겨보며, C#의 컬렉션 문법을 새롭게 익혔습니다.

| 기능 | Go | C# |
|---|---|---|
| **자료구조 이름** | Map | Dictionary |
| **생성 방법** | `make(map[int]int)` | `new Dictionary<int, int>()` |
| **키 존재 여부 (bool)** | `if j, ok := map[key]; ok` | `if (map.ContainsKey(key))` |
| **데이터 할당** | `map[key] = value` | `map[key] = value` |

## 🏃‍♂️ 코드 실행
* **Go:** `go run main.go`
* **C#:** `dotnet run` (현재 디렉토리에서)

---

## 💻 풀이 코드

### Go
```go
package main

func twoSum(nums []int, target int) []int {
    res := make([]int, 2)
    numMap := make(map[int]int)

    for i := 0; i < len(nums); i++ {
        complement := target - nums[i]

        if j, ok := numMap[complement]; ok {
            res[0], res[1] = i, j
            break
        }
        numMap[nums[i]] = i
    }

    return res
}
```
---
### C#
```c#
using System;
using System.Collections.Generic;

public class Solution
{
    public int[] TwoSum(int[] nums, int target)
    {
        int[] res = new int[2];
        Dictionary<int, int> numMap = new Dictionary<int, int>();

        for (int i = 0; i < nums.Length; i++)
        {
            if (numMap.ContainsKey(target - nums[i]))
            {
                res[0] = i;
                res[1] = numMap[target - nums[i]];
                break;
            }
            numMap[nums[i]] = i;
        }
        return res;
    }
}
```