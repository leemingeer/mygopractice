package main


type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}


func levelOrder(root *TreeNode) [][]int {

	if root == nil{
		return [][]int{}
	}
	res := [][]int{}
	q := []*TreeNode{root}
	for len(q) != 0{
		cur_res := []int{} // 当前层，每层都要清空
		length := len(q)
		for i:=0;i<length;i++{ // 遍历当前层的所有节点
			cur_res = append(cur_res, q[i].Val)
			if q[i].Left != nil{
				q = append(q, q[i].Left) // 当前节点的左节点
			}
			if q[i].Right != nil{
				q = append(q, q[i].Right) //// 当前节点的左节点
			}
		}
		res = append(res, cur_res) // 每一层数据都要写入res中
		q = q[length:] //已经被返回过的结点需要出队
	}
	return res
}