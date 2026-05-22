# 0033-search-in-rotated-sorted-array

## 💡 문제 링크
* [LeetCode - Search in Rotated Sorted Array](https://leetcode.com/problems/search-in-rotated-sorted-array/)

## 🧠 문제 접근 및 풀이
* **알고리즘 분류:** Binary Search (이진 탐색)
* **시간 복잡도:** $O(\log N)$
* **공간 복잡도:** $O(1)$

### 핵심 아이디어
* 오름차순 정렬된 배열이 특정 지점을 기준으로 회전(Rotation)되어 있어도, **어느 지점을 `mid`로 잡고 반으로 나누든 한쪽 서브 배열은 무조건 정렬되어 있다**는 성질을 이용합니다.
* 정렬된 쪽(`nums[left] <= nums[mid]`)이 어디인지 먼저 찾고, `target`이 그 정렬된 구간 내에 존재하는지 확인하며 탐색 범위를 좁혀나갑니다.

## 📝 배운 점

### 1. 알고리즘: 직관적인 관찰을 통한 복잡한 조건문 단순화
처음에는 회전된 여러 상황을 하나하나 대입해가며 예외 처리를 하려고 했습니다. 하지만 그렇게 조건문을 촘촘하게 짜다 보니 코드가 너무 복잡해지고 스스로도 헷갈리는 상황이 생겼습니다.

무작정 예외 조건만 늘리는 대신, 배열의 구조적 특징을 다시 차분히 살펴보았습니다. 어떤 상태이든 **"반으로 쪼개면 최소한 한쪽은 예쁘게 오름차순 정렬되어 있다"**는 힌트를 중심으로 접근 방식을 바꾸니, 복잡하던 분기 처리가 훨씬 단순해지고 명확한 코드를 작성할 수 있었습니다. 복잡한 문제를 만날수록 무작정 구현하기보다 구조적 특징을 먼저 파악하는 게 중요하다는 것을 깊이 깨달았습니다.

### 2. 언어 스터디: 문법적 특이사항 없음
이번 문제에서는 C#의 새로운 문법이나 독특한 자료구조를 사용하지 않고, 제어문과 기본 배열 구조만으로 해결했습니다. Go와 C#의 `while` / `for` 반복문 구성 및 기본 연산 구조가 매우 유사하여 구현 상의 큰 차이점 없이 자연스럽게 이식할 수 있었습니다.

## 🏃‍♂️ 코드 실행
* **Go:** `go run main.go`
* **C#:** `dotnet run` (현재 디렉토리에서)

---

## 💻 풀이 코드

### Go
```go
package main

func search(nums []int, target int) int {
    left, right := 0, len(nums)-1

    for left <= right {
        mid := (left + right) / 2

        if nums[mid] == target {
            return mid
        }

        // 왼쪽 반쪽이 정렬되어 있는 경우
        if nums[left] <= nums[mid] {
            if nums[left] <= target && target < nums[mid] {
                right = mid - 1
            } else {
                left = mid + 1
            }
        // 오른쪽 반쪽이 정렬되어 있는 경우
        } else {
            if nums[mid] < target && target <= nums[right] {
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
    }

    return -1
}
```
---
### C#
```c#
using System;

public class Solution
{
    public int Search(int[] nums, int target)
    {
        int left = 0;
        int right = nums.Length - 1;

        while (left <= right)
        {
            int mid = (left + right) / 2;

            if (target == nums[mid])
            {
                return mid;
            }

            // 왼쪽 반쪽이 정렬되어 있는 경우
            if (nums[left] <= nums[mid])
            {
                if (nums[left] <= target && target < nums[mid])
                {
                    right = mid - 1;
                }
                else
                {
                    left = mid + 1;
                }
            }
            // 오른쪽 반쪽이 정렬되어 있는 경우
            else
            {
                if (nums[mid] < target && target <= nums[right])
                {
                    left = mid + 1;
                }
                else
                {
                    right = mid - 1;
                }
            }
        }
        return -1;
    }
}
```