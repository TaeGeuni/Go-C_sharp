# 2095-delete-the-middle-node-of-a-linked-list

## 💡 문제 링크
* [LeetCode - Delete the Middle Node of a Linked List](https://leetcode.com/problems/delete-the-middle-node-of-a-linked-list/)

## 🧠 문제 접근 및 풀이
* **알고리즘 분류:** Linked List, Two Pointers (Fast & Slow Pointers)
* **시간 복잡도:** $O(N)$
* **공간 복잡도:** $O(1)$

### 핵심 아이디어
* 연결 리스트의 전체 길이를 모르는 상태에서 중앙을 찾기 위해 투 포인터(런너 기법)를 사용합니다.
* 한 칸씩 이동하는 `slow` 포인터와 두 칸씩 이동하는 `fast` 포인터를 동시에 출발시킵니다.
* `fast` 포인터가 리스트의 끝에 도달했을 때, `slow` 포인터는 자연스럽게 리스트의 중앙(또는 중앙 직전)에 위치하게 되는 원리를 이용해 노드를 삭제합니다.

## 📝 배운 점

### 1. 알고리즘: 투 포인터(Fast & Slow Pointers)를 활용한 중앙값 찾기
배열과 달리 인덱스로 접근할 수 없는 연결 리스트에서, 한 칸씩 움직이는 포인터와 두 칸씩 움직이는 포인터를 활용하면 리스트를 두 번 순회하지 않고도 단번에 중앙 위치를 찾을 수 있다는 아주 유용한 테크닉을 배웠습니다. 

### 2. 언어 스터디: C#의 참조 타입(Reference Type)과 연결 리스트 구현
Go 언어에서는 명시적으로 포인터(`*ListNode`)를 사용하여 메모리 주소를 다루었기 때문에 연결 리스트의 다음 노드를 가리키는 것이 직관적이었습니다. 

하지만 C# 같은 언어는 명시적인 포인터 기호가 없어 "어떻게 다음 노드의 주소를 기억하고 연결할까?" 하는 의문이 있었습니다. 이번 구현을 통해 "C#에서 클래스(Class) 타입의 변수는 그 자체가 힙(Heap) 메모리의 주소값을 저장하는 포인터(참조 변수)로 작동한다"는 핵심 원리를 깨달았습니다. 따라서 `ListNode next` 자체가 이미 다음 노드의 주소를 가리키는 역할을 완벽하게 수행합니다.

## 🏃‍♂️ 코드 실행
* **Go:** `go run main.go`
* **C#:** `dotnet run` (현재 디렉토리에서)

---

## 💻 풀이 코드

### Go
```go
package main

type ListNode struct {
    Value int
    Next  *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return nil
    }
    slow := head
    fast := head.Next.Next

    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    slow.Next = slow.Next.Next

    return head
}
```
---
### C#
```c#
public class ListNode
{
    public int val;
    public ListNode next;
    public ListNode(int val = 0, ListNode next = null)
    {
        this.val = val;
        this.next = next;
    }
}

public class Solution
{
    public ListNode DeleteMiddle(ListNode head)
    {
        if (head == null || head.next == null)
        {
            return null;
        }
        var slow = head;
        var fast = head.next.next;

        while (fast != null && fast.next != null)
        {
            fast = fast.next.next;
            slow = slow.next;
        }

        slow.next = slow.next.next;

        return head;
    }
}
```