// 验证二叉搜索树
package main

import "fmt"

type TreeNode struct {
    Val string
    Left *TreeNode
    Right *TreeNode
}

/**
* 二叉树前序遍历   根-> 左-> 右
* @param node 二叉树节点
*/
func preOrderTraveral(node *TreeNode){
    if(node == nil){
        return
    }
    fmt.Print(node.Val)
    preOrderTraveral(node.Left)
    preOrderTraveral(node.Right)
}

/**
* 二叉树前序遍历   根-> 左-> 右 非递归实现方式
* @param node 二叉树节点
*/
func preOrderTraveral2(node *TreeNode){
    var stack []*TreeNode
    treeNode := node
    for { 
        if treeNode == nil && len(stack)==0{
            break 
        }
        for{
            if treeNode == nil {
                break; 
            }            
            stack = append(stack,treeNode)
            treeNode = treeNode.Left     
        }
        if len(stack)>0 {
            treeNode = stack[len(stack)-1]
            fmt.Print(treeNode.Val)
            treeNode = treeNode.Right
            stack = stack[:len(stack)-1]
        }   
    }
}

/**
* 二叉树前序遍历   根-> 左-> 右 非递归实现方式
* @param node 二叉树节点
*/
func postOrderTraveral(node *TreeNode){
    var stack []*TreeNode
    treeNode := node
    var last *TreeNode 
    for { 
        if treeNode == nil && len(stack)==0{
            break 
        }
        // if treeNode != nil {
        //     fmt.Println(treeNode.Val) 
        // }else{
        //     fmt.Println(stack[0].Val) 
        // }

        for{
            if treeNode == nil {
                break; 
            }            
            stack = append(stack,treeNode)
            treeNode = treeNode.Left     
        }
        if len(stack)>0 { 
            treeNode = stack[len(stack)-1]          
            if treeNode.Right == nil || treeNode.Right == last {
                fmt.Print(treeNode.Val)
                stack = stack[:len(stack)-1]           
                last = treeNode
                treeNode = nil
            }else{                
                treeNode = treeNode.Right    
            } 
        }   
    }
}

func levelOrder(node *TreeNode){
    var queue []*TreeNode
    queue = append(queue,node)
    for{
        if len(queue)==0 {
            break 
        }
        root := queue[0]
        queue = queue[1:] 
        fmt.Print(root.Val)
        if root.Left != nil {
            queue = append(queue,root.Left) 
        }
        if root.Right != nil {
            queue = append(queue,root.Right) 
        }
    }
}

// 下面是数据验证过程

func main(){
    var root *TreeNode
    left := &TreeNode{"B",&TreeNode{"D",nil,nil},&TreeNode{"F",&TreeNode{"E",nil,nil},nil}}
    right := &TreeNode{"C",&TreeNode{"G",nil,&TreeNode{"H",nil,nil}},&TreeNode{"I",nil,nil}}
    root = &TreeNode{"A",left,right}

    levelOrder(root)
}

