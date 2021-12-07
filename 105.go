// 105. 从前序与中序遍历序列构造二叉树
package main

import "fmt"
import "strconv"

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
func buildTree(preorder []int, inorder []int) *TreeNode {
    var root *TreeNode
    if len(preorder) == 0 {
        return nil
    }
    // 将中序遍历每个value和key建立队里，左节点的中序遍历key一定在前面
    var inorderTrans = make(map[string]int)
    for key, value := range inorder {
        inorderTrans[strconv.Itoa(value)] = key
    }
    root = &TreeNode{preorder[0],nil,nil}
    for i := 1; i < len(preorder); i++ {
        value := preorder[i]
        newNode := &TreeNode{value,nil,nil}
        node := root
        for{
            if(inorderTrans[strconv.Itoa(value)] > inorderTrans[strconv.Itoa(node.Val)]){
                if node.Right == nil {
                    node.Right = newNode
                    break
                }else{
                    node = node.Right
                }
                
            }            
            if(inorderTrans[strconv.Itoa(value)] < inorderTrans[strconv.Itoa(node.Val)]){
                if node.Left == nil {
                    node.Left = newNode
                    break
                }else{
                    node = node.Left
                }
                
            }
        }
    }
    return root
}

func inOrderTraveral(node *TreeNode){
    if(node == nil){
        return
    }
    fmt.Print(node.Val)
    fmt.Print(" ")
    inOrderTraveral(node.Left)
    inOrderTraveral(node.Right)
}

// 下面是数据验证过程

func main(){
    // var root *TreeNode
    // left := &TreeNode{9,nil,nil}
    // right := &TreeNode{20,&TreeNode{15,nil,nil},&TreeNode{7,nil,nil}}
    // root = &TreeNode{3,left,right}

    // root = nil
    // preorder := []int{3,9,20,15,7}
    // inorder  := []int{9,3,15,20,7}
    // preorder := []int{-1}
    // inorder  := []int{-1}
    // preorder := []int{3,1,2,4}
    // inorder  := []int{1,2,3,4}
    preorder := []int{7,-10,-4,3,-1,2,-8,11}
    inorder  := []int{-4,-10,3,-1,7,11,-8,2}

    result := buildTree(preorder,inorder)
    fmt.Println(result)
    inOrderTraveral(result)
}

