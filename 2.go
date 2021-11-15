package main

import "fmt"

/**
 *  第二题，两数相加
 */

type ListNode struct {
    Val int
    Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var plus int
    var head, cur *ListNode
    i := 0
    for l1 != nil || l2 != nil {
        i ++
        a, b := 0, 0
        if l1 != nil {
            a = l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            b = l2.Val
            l2 = l2.Next
        }
        sum := a + b + plus
        sum, plus = sum%10, sum/10
        if head == nil {
            head = &ListNode{Val: sum}
            cur = head
        } else {
            cur.Next = &ListNode{Val: sum}
            cur = cur.Next
        }
    }
    if plus > 0 {
        cur.Next = &ListNode{Val: plus}
    }
    return head
}


// 下面是数据验证过程

func main(){
    var l1 ListNode
    var l2 ListNode

    l1List := []int{9,9,9,9,9,9}
    l1.set(l1List)
    l1.add(9)

    l2List := []int{9,9,9,9}
    l2.set(l2List)

    fmt.Println(l1.toArray())
    fmt.Println(l2.toArray())
    result := addTwoNumbers(&l1,&l2)
    
    fmt.Println(result.toArray())
}


// 给list添加节点
func (listNode *ListNode)addNode(newListNode *ListNode){
    if listNode.Next==nil{
        listNode.Next=newListNode
    }else{
        listNode.Next.addNode(newListNode)
    }
}

// 初始化list给定的一组数据
func (listNode *ListNode)set(datas []int){
    for i:=0;i<len(datas);i++ {
        if i==0 {
            listNode.Val = datas[i]
        }else{
           var newListNode ListNode
           newListNode.Val = datas[i] 
           listNode.addNode(&newListNode)
        }
    }
}

// 初始化list给订单的一组数据
func (listNode *ListNode)add(data int){
    var newListNode ListNode
    newListNode.Val=data
    if listNode.Val==0 && listNode.Next==nil && data!=0{
        listNode.Val = data
    }else{
        listNode.addNode(&newListNode)
    }
}

// 返回所有节点的内容，以数组打印的方式比较方面查看
func (listNode *ListNode)toArray() (datas []interface{}){
    if listNode == nil{
        return nil
    }
    node := listNode
    for{
        datas=append(datas,node.Val)
        if node.Next!=nil{
            node = node.Next
        }else{
            break
        }
    }
    return datas
}

