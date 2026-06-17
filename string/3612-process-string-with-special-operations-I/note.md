# 3612-process-string-with-special-operations-I

## 💡 문제 링크
* [LeetCode - Process String with Special Operations I](https://leetcode.com/problems/process-string-with-special-operations-i/) 

## 🧠 문제 접근 및 풀이
* **알고리즘 분류:** String, Simulation
* **시간 복잡도:** $O(N)$ (문자열 순회)
* **공간 복잡도:** $O(N)$

### 핵심 아이디어
* 주어진 문자열을 순회하면서, 특정 특수 문자(`*`, `#`, `%`)가 등장할 때마다 정의된 연산을 수행하는 시뮬레이션 문제입니다.
* 문자열을 동적으로 수정해야 하므로, 불변(Immutable)인 문자열 대신 동적 배열(Slice 또는 List)을 활용하여 문자를 조작한 후 마지막에 다시 문자열로 변환합니다.

## 📝 배운 점

이번 문제에서는 새로운 알고리즘 기법보다는 언어의 타입 시스템과 설계 철학에 대해 깊이 탐구할 수 있었습니다.

### 1. 언어 스터디: 문자열 처리와 `char` 타입의 부재 (Go vs C# / Java)
Go 언어로 문자열을 처리할 때는 각 요소를 `byte` 타입으로 다루었는데, C#에서는 `char` 타입으로 다루는 것을 확인했습니다. "왜 Go에는 `char` 타입이 없을까?" 하는 의문이 들어 조사해 본 결과, 언어 설계 역사의 흥미로운 차이점을 알게 되었습니다.

* **C# / Java의 `char` (2바이트, UTF-16):**
  초기 설계 당시에는 2바이트(65,536개)면 전 세계의 모든 문자를 표현할 수 있을 것이라 판단했습니다. 하지만 유니코드가 확장되고 이모지나 고문자가 추가되면서, 2바이트를 초과하는 문자는 `char` 한 개로 표현하지 못하고 2개를 묶어 써야 하는 써로게이트 페어(Surrogate pair) 문제가 발생했습니다.
* **Go의 `rune` (4바이트, UTF-32):**
  비교적 최근에 설계된 Go 언어는 이러한 문제를 인지하고, 애초에 `char` 타입을 없앴습니다. 대신 전 세계 모든 문자(이모지 포함)를 완벽하게 수용할 수 있는 4바이트 정수형인 `rune`(`int32`의 별칭)을 도입하여 유니코드 문제를 깔끔하게 해결했습니다. 영문이나 ASCII만 처리할 때는 가벼운 `byte`(`uint8`의 별칭)를 사용합니다.

### 2. 언어 스터디: 동적 배열 자료구조 (`List<T>`)
Go 언어의 슬라이스(`[]`)처럼 C#에서도 배열의 크기를 동적으로 조절할 수 있는 자료구조인 `List<T>`를 새롭게 배우고 적용해 보았습니다.

| 기능 | Go (Slice) | C# (`List<T>`) |
|---|---|---|
| **생성** | `make([]byte, 0)` | `new List<char>()` |
| **추가 (1개)** | `append(sli, s[i])` | `list.Add(s[i])` |
| **추가 (여러 개)** | `append(sli, sli...)` | `list.AddRange(list)` |
| **마지막 요소 삭제** | `sli[:len(sli)-1]` | `list.RemoveAt(list.Count - 1)` |
| **역순 정렬** | 별도 반복문 구현 | `list.Reverse()` |
| **문자열 변환** | `string(sli)` | `new string(list.ToArray())` |

* C#의 `List<T>`는 `Reverse()`, `AddRange()` 등 강력한 내장 메서드를 제공하여 문자열 조작 로직을 Go보다 훨씬 간결하게 작성할 수 있었습니다.

## 🏃‍♂️ 코드 실행
* **Go:** `go run main.go`
* **C#:** `dotnet run` (현재 디렉토리에서)

---

## 💻 풀이 코드

### Go
```go
package main

func processStr(s string) string {
    btSli := make([]byte, 0)

    for i := 0; i < len(s); i++ {
        if s[i] == '*' {
            if len(btSli) != 0 {
                btSli = btSli[:len(btSli)-1]
            }
        } else if s[i] == '#' {
            btSli = append(btSli, btSli...)
        } else if s[i] == '%' {
            rvSli := make([]byte, len(btSli))
            for j := len(btSli) - 1; j >= 0; j-- {
                rvSli[len(btSli)-1-j] = btSli[j]
            }
            btSli = rvSli
        } else {
            btSli = append(btSli, s[i])
        }
    }

    return string(btSli)
}
```
---
### C#
```c#
public class Solution
{
    public string ProcessStr(string s)
    {
        List<char> btArr = new List<char> { };

        for (int i = 0; i < s.Length; i++)
        {
            if (s[i] == '*')
            {
                if (btArr.Count != 0)
                {
                    btArr.RemoveAt(btArr.Count - 1);
                }
            }
            else if (s[i] == '#')
            {
                btArr.AddRange(btArr);
            }
            else if (s[i] == '%')
            {
                btArr.Reverse();
            }
            else
            {
                btArr.Add(s[i]);
            }
        }
        string res = new string(btArr.ToArray());
        return res;
    }
}
```