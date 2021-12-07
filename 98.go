// 验证二叉搜索树
package main

import "fmt"
import "math"

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
func isValidBST(root *TreeNode) bool {
    return checkCurentNode(root,math.MinInt64,math.MaxInt64)
}

func checkCurentNode(curentNode *TreeNode,min int,max int) bool {
    if curentNode == nil {
        return true
    }
    if curentNode.Val <= min || curentNode.Val >= max {
        return false
    }

    return checkCurentNode(curentNode.Left,min,curentNode.Val) && checkCurentNode(curentNode.Right,curentNode.Val,max)
}

// 下面是数据验证过程

func main(){
    var root *TreeNode
    left := &TreeNode{4,nil,nil}
    right := &TreeNode{6,&TreeNode{3,nil,nil},&TreeNode{7,nil,nil}}
    root = &TreeNode{5,left,right}
    // left := &TreeNode{1,nil,nil}
    // right := &TreeNode{3,nil,nil}
    // root = &TreeNode{2,left,right}

    result := isValidBST(root)
    fmt.Println(result);
}

