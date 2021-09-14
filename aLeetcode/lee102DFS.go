package main
var res1 [][]int
func levelOrder1(root *TreeNode) [][]int {

	if root == nil{
		return [][]int{}
	}
	res = [][]int{}
	dfs(root, 0)
	return res1
}

func dfs(root *TreeNode, level int){
	if root == nil {
		return
	}
	if len(res1) == level{
		res1 = append(res1, []int{})
	}
	res1[level] = append(res1[level], root.Val)
	dfs(root.Left, level+1)
	dfs(root.Right, level+1)
}