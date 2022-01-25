pakcage tree

func isValidBST(root *TreeNode) bool {
    if root == nil {
        return true
    }
    return isValidHelper(root,nil,nil)
}

func isValidHelper(root *TreeNode, min, max *TreeNode) bool {
    if root == nil {
        return true
    }
    if min != nil && min.Val >= root.Val || max != nil && max.Val <= root.Val {
        return false
    }  
    return isValidHelper(root.Left, min, root) && isValidHelper(root.Right, root, max)
}
