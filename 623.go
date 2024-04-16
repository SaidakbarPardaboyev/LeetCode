/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
    dummy := &TreeNode{
        Val : 0,
        Left : root,
    }
    level := 1
    allNodesInPerLevel := []*TreeNode{dummy}
    allNodesInLastLevel := []*TreeNode{dummy}
    for len(allNodesInPerLevel) > 0 {
        if level == depth {
            AddNewNodesToNthLevel(allNodesInPerLevel, val)
        }
        l := len(allNodesInPerLevel)

        allNodesInLastLevel = make([]*TreeNode, len(allNodesInPerLevel))
        copy(allNodesInLastLevel, allNodesInPerLevel)
        
        for i := 0; i < l; i++ {
            node := allNodesInPerLevel[0]
            allNodesInPerLevel = allNodesInPerLevel[1:]
            if node.Left != nil {
                allNodesInPerLevel = append(allNodesInPerLevel, node.Left)
            }
            if node.Right != nil {
                allNodesInPerLevel = append(allNodesInPerLevel, node.Right)
            }
        }
        level++
    }
    // Add to the end of the last level
    if level-1 == depth {
        AddNewNodeToEnd(allNodesInLastLevel, val)
    }
    
    return dummy.Left
}
func AddNewNodesToNthLevel(allNodesInPerLevel []*TreeNode, val int) {
    for _, node := range allNodesInPerLevel {
        node.Left = &TreeNode{
            Val: val,
            Left: node.Left}
        node.Right = &TreeNode{
            Val: val,
            Right: node.Right}
    }
}

func AddNewNodeToEnd(parentNodes []*TreeNode, val int) {
    for _, node := range parentNodes {
        node.Left = &TreeNode{
            Val : val,
        }
        node.Right = &TreeNode{
            Val: val,
        }
    }
}