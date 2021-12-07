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



/**
 * 搜索二叉树的中序遍历是递增的，然后查询到不符合的2个数组下标，进行替换即可
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree(root *TreeNode)  {
    list := []int{}
    var inOrder func(node *TreeNode)
    inOrder = func(node *TreeNode){
        if(node == nil){
            return
        }
        inOrder(node.Left)
        list = append(list,node.Val)
        inOrder(node.Right)
    }
    inOrder(root)
    fmt.Println(list)
    x,y := findTwoSwap(list)
    fmt.Println(x,y)
    swapTwoVal(root,2,x,y)

    list = []int{}
    inOrder(root)
    fmt.Println(list)
}

func findTwoSwap(list []int)(int,int){
    x,y := -1,-1
    for i:=0;i<len(list)-1;i++{
        if(list[i+1]<list[i]){
            y = i+1
            if x != -1 {
                break 
            }else{
               x = i
            }
        }
    }
    x,y = list[x],list[y]
    return x,y
}

func swapTwoVal(node *TreeNode,count,x,y int){
    if node == nil {
        return 
    }
    if node.Val == x {
        node.Val = y
        count --
    }else if(node.Val==y){
        node.Val = x
        count --
    }
    if count == 0 {
        return 
    }
    swapTwoVal(node.Left,count,x,y)
    swapTwoVal(node.Right,count,x,y)
}


// 下面是数据验证过程

func main(){
    var root *TreeNode
    // left := &TreeNode{1,nil,nil}
    // right := &TreeNode{4,&TreeNode{2,nil,nil},nil}
    // root = &TreeNode{3,left,right}
    // left := &TreeNode{3,nil,&TreeNode{2,nil,nil}}
    // root = &TreeNode{1,left,nil}

    root = &TreeNode{-1,nil,&TreeNode{-3,nil,&TreeNode{-2,nil,&TreeNode{-4,nil,nil}}}}

    recoverTree(root)
}

