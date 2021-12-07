// 102. 二叉树的层序遍历
package main

import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    var result [][]int
    if root == nil {
        return result
    }
    var currentLever []*TreeNode
    currentLever = append(currentLever,root)
    for{
        if len(currentLever)==0 {
            break 
        }
        var leverVal []int
        var nextLever []*TreeNode
        for _,node := range currentLever{
            leverVal = append(leverVal,node.Val)
            if node.Left != nil {
                nextLever = append(nextLever,node.Left) 
            }
            if node.Right != nil {
                nextLever = append(nextLever,node.Right) 
            }
        }
        result = append(result,leverVal)
        currentLever = nextLever
    }
    return result
}


// 下面是数据验证过程

func main(){
    var root *TreeNode
    // left := &TreeNode{9,nil,nil}
    // right := &TreeNode{20,&TreeNode{15,nil,nil},&TreeNode{7,nil,nil}}
    // root = &TreeNode{3,left,right}

    root = nil

    result := levelOrder(root)
    fmt.Print(result);
}

