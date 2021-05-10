/**
 * 328. Odd Even Linked List
 * https://leetcode.com/problems/odd-even-linked-list/
 * 
 * submission: faster than 92%
 */

struct ListNode
{
    int val;
    ListNode *next;
    ListNode() : val(0), next(nullptr) {}
    ListNode(int x) : val(x), next(nullptr) {}
    ListNode(int x, ListNode *next) : val(x), next(next) {}
};

ListNode *oddEvenList(ListNode *head)
{
    if (head == nullptr)
    {
        return head;
    }
    
    ListNode *evenHead = head->next;
    ListNode *curEven = evenHead, *curOdd = head;
    ListNode *cur = head;
    bool odd = true;
    while (cur->next) {
        ListNode *next = cur->next;
        if (odd)
        {
            if (cur->next->next) {
                curOdd->next = cur->next->next;
                curOdd = curOdd->next;
            }
        }
        else
        {
            curEven->next = cur->next->next;
            curEven = curEven->next;
        }
        odd = !odd;
        cur = next;
    }
    curOdd->next = evenHead;
    return head;
}