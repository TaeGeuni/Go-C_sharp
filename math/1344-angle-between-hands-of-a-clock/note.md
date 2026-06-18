# 1344-angle-between-hands-of-a-clock

## 💡 문제 링크
* [LeetCode - Angle Between Hands of a Clock](https://leetcode.com/problems/angle-between-hands-of-a-clock/)

## 🧠 문제 접근 및 풀이
* **알고리즘 분류:** Math
* **시간 복잡도:** $O(1)$
* **공간 복잡도:** $O(1)$

### 핵심 아이디어
* 시계의 **분침**은 1분에 6도씩 움직입니다. (`minutes * 6`)
* 시계의 **시침**은 1시간에 30도씩 움직이며, 분이 지남에 따라 1분에 0.5도씩 추가로 움직입니다. (`hour * 30 + minutes * 0.5`)
* 두 바늘이 가리키는 각도를 각각 12시(0도)를 기준으로 계산한 뒤, 두 각도의 차이(절댓값)를 구합니다.
* 만약 그 차이가 180도를 넘는다면, 더 작은 쪽의 각도를 구해야 하므로 360에서 빼줍니다.

## 📝 배운 점

### 1. 알고리즘: 단순 수학 시뮬레이션
이번 문제에서는 특별한 자료구조나 복잡한 알고리즘 기법보다는, 문제의 요구사항을 수식으로 정확히 모델링하는 수학적 시뮬레이션 능력이 중요했습니다. 

### 2. 언어 스터디: 수학 유틸리티 함수 (Go vs C#)
두 값의 대소를 비교하는 기능에 대해 Go 언어와 C#의 차이점을 알게 되었습니다.

| 언어 | `min`, `max` 활용 방식 |
|---|---|
| **Go** | 과거에는 정수형 변환의 번거로움 때문에 직접 함수를 구현(`func min(a, b)`)해서 사용하는 패턴이 흔했습니다. *(참고: Go 1.21 버전부터는 제네릭 기반의 내장 함수 `min()`, `max()`가 공식 추가되었습니다.)* |
| **C#** | `System` 네임스페이스의 `Math.Max()`, `Math.Min()` 메서드를 사용하여 별도의 구현 없이 즉시 깔끔하게 대소 비교와 절댓값 계산이 가능합니다. |

* C#의 강력한 표준 라이브러리(BTL)가 제공하는 편의성을 체감할 수 있었습니다.

## 🏃‍♂️ 코드 실행
* **Go:** `go run main.go`
* **C#:** `dotnet run` (현재 디렉토리에서)

---

## 💻 풀이 코드

### Go
```go
package main

func min(a, b float64) float64 {
    if a > b {
        return b
    }
    return a
}

func max(a, b float64) float64 {
    if a > b {
        return a
    }
    return b
}

func angleClock(hour int, minutes int) float64 {
    var h, m, res float64

    m = float64(minutes) * 6
    h = float64(hour)*30 + float64(minutes)*0.5

    res = max(h, m) - min(h, m)

    if res > 180 {
        return 360 - res
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
    public double AngleClock(int hour, int minutes)
    {
        double h, m, res = 0;

        m = minutes * 6;
        h = hour * 30 + minutes * 0.5;

        // C# 내장 Math 클래스 활용
        res = Math.Max(h, m) - Math.Min(h, m);

        if (res > 180)
        {
            return 360 - res;
        }
        return res;
    }
}
```