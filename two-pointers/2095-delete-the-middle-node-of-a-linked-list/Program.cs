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